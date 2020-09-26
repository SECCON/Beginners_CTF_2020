#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdlib.h>
#include <dlfcn.h>

const char *PRIV[] = {
  "---", "--x", "-w-", "-wx", "r--", "r-x", "rw-", "rwx"
};
char *A, *B;
unsigned long *main_arena, *tcache, *ptr_hook, *ptr_win;

int get_privilege(void *addr) {
  int priv = 0;
  unsigned long start, end, target;
  char buf[0x100], *ptr;
  int fd = open("/proc/self/maps", O_RDONLY);
  target = (unsigned long)addr;
  for(ptr = buf, start = end = 0; ; ptr++) {
    if (read(fd, ptr, 1) <= 0) {
      break;
    }
    if (*ptr == '-' && start == 0) {
      start = strtol(buf, NULL, 16);
      ptr = buf - 1;
    } else if (*ptr == ' ' && end == 0) {
      end = strtol(buf, NULL, 16);
      ptr = buf - 1;
    } else if (*ptr == '\n') {
      if (start <= target && target < end) {
        priv |= (buf[0] == 'r') << 2;
        priv |= (buf[1] == 'w') << 1;
        priv |= (buf[2] == 'x') << 0;
        break;
      }
      start = end = 0;
      ptr = buf - 1;
    }
  }
  
  close(fd);
  return priv;
}

void describe_tcache(void) {
  int i, priv;
  unsigned long ptr = (unsigned long)tcache[10];
  puts("-=-=-=-=-= TCACHE -=-=-=-=-=");
  puts("[    tcache (for 0x20)    ]");
  for(i = 0; i < 7; i++) {
    puts("        ||");
    puts("        \\/");
    if (ptr == 0) {
      puts("[      END OF TCACHE      ]");
      break;
    } else {
      priv = get_privilege((void*)ptr);
      printf("[ 0x%016lx(%s) ]\n", ptr, PRIV[priv]);
      if (priv & 0b100) {
        ptr = *(unsigned long*)ptr;
      } else {
        puts("        ||");
        puts("        \\/");
        puts("[       BROKEN LINK       ]");
        break;
      }
    }
  }
  puts("-=-=-=-=-=-=-=-=-=-=-=-=-=-=");
}

void describe_heap(void) {
  int i;
  unsigned long *ptr = (unsigned long*)A - 2;
  puts("-=-=-=-=-= HEAP LAYOUT =-=-=-=-=-");
  printf(" [+] A = %p\n [+] B = %p\n\n", A, B);
  for(i = 0; i < 10; i++, ptr++) {
    puts("                   +--------------------+");
    if (ptr == (unsigned long*)A) {
      printf("0x%016lx | 0x%016lx | <-- A\n", (unsigned long)ptr, *ptr);
    } else if (ptr == (unsigned long*)B) {
      printf("0x%016lx | 0x%016lx | <-- B\n", (unsigned long)ptr, *ptr);
    } else {
      printf("0x%016lx | 0x%016lx |\n", (unsigned long)ptr, *ptr);
    }
  }
  puts("                   +--------------------+");
  puts("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-");
}

void hint(void) {
  int i;
  unsigned long ptr = (unsigned long)tcache[10];
  int isChunkLinked, isBroken, isHookLinked, isSystemWritten;
  unsigned long size;
  isChunkLinked = isBroken = isHookLinked = isSystemWritten = 0;

  for(i = 0; i < 7; i++) {
    if (ptr == 0) {
      break;
    } else if (get_privilege((void*)(ptr - 8)) & 0b100) {
      if (ptr == (unsigned long)ptr_hook) {
        isHookLinked = 1;
        break;
      } else {
        size = *(unsigned long*)(ptr - 0x8);
      }
      ptr = *(unsigned long*)ptr;
    } else {
      isBroken = 1;
      break;
    }
  }
  if (*ptr_hook == (unsigned long)ptr_win)
    isSystemWritten = 1;

  if (isSystemWritten) {
    puts("It seems you did everything right!");
    puts("`free` is now equivalent to `win`");
  } else if (isHookLinked) {
    puts("It seems __free_hook is successfully linked to tcache!");
    if (i == 0) {
      puts("The first link of tcache is __free_hook!");
      if (B == NULL) {
        puts("Also B is empty! You know what to do, right?");
      } else {
        puts("But B is not empty...");
      }
    } else {
      if (size > 0x411) {
        puts("But the chunk size is broken or too big maybe...?");
      } else if (size < 0x20) {
        puts("But the chunk size is too small maybe...?");
      } else if (size == 0x20 || size == 0x21) {
        puts("But you can't get __free_hook since you can only malloc/free B.");
        printf("What if you change the chunk size to a value other than 0x%02lx...?\n", size);
      } else if ((size & 0b110) == 0) {
        puts("And the chunk size is properly forged!");
      } else {
        if (size & 0b100) {
          puts("But is_main_arena bit is 1...?");
        } else {
          puts("But is_mmaped bit is 1...?");
        }
      }
    }
  } else if (isBroken) {
    puts("Good. The tcache link is corrupted!");
    printf("Currently it's linked to 0x%016lx but what if it's __free_hook...?\n", ptr);
  } else {
    puts("Tcache manages freed chunks in linked lists by size.");
    puts("Every list can keep up to 7 chunks.");
    puts("A freed chunk linked to tcache has a pointer (fd) to the previously freed chunk.");
    puts("Let's check what happens when you overwrite fd by Heap Overflow.");
  }
  putchar('\n');
}

int menu(void) {
  char buf[16];
  puts("1. read(0, A, 0x80);");
  puts("2. B = malloc(0x18); read(0, B, 0x18);");
  puts("3. free(B); B = NULL;");
  puts("4. Describe heap");
  puts("5. Describe tcache (for size 0x20)");
  puts("6. Currently available hint");
  printf("> ");
  if (read(0, buf, 15) <= 0) exit(0);
  return atoi(buf);
}

void beginners_heap(void) {
  while(1) {
    switch(menu()) {
    case 1:
      read(0, A, 0x80);
      break;
    case 2:
      if (B == NULL) {
        B = malloc(0x18);
        read(0, B, 0x18);
      } else {
        puts("[-] B is already used...");
      }
      break;
    case 3:
      if (B == NULL) {
        puts("[-] B is already free...");
      } else {
        free(B);
        B = NULL;
      }
      break;
    case 4: describe_heap(); break;
    case 5: describe_tcache(); break;
    case 6: hint(); break;
    default:
      puts("Bye.");
      return;
    }
  }
}

void win(void) {
  char flag[0x40] = {0};
  int fd = open("flag", O_RDONLY);
  read(fd, flag, 0x40);
  puts("Congratulations!");
  puts(flag);
  close(fd);
  _exit(0);
}

int main(void) {
  setbuf(stdin, NULL);
  setbuf(stdout, NULL);
  setbuf(stderr, NULL);
  void *handle = dlopen("/lib/x86_64-linux-gnu/libc-2.27.so", RTLD_LAZY);
  ptr_hook = dlsym(handle, "__free_hook");
  main_arena = dlsym(handle, "__main_arena");
  ptr_win = (void*)&win;

  puts("Let's learn heap overflow today");
  puts("You have a chunk which is vulnerable to Heap Overflow (chunk A)\n");
  puts(" A = malloc(0x18);\n");
  puts("Also you can allocate and free a chunk which doesn't have overflow (chunk B)");
  puts("You have the following important information:\n");
  printf(" <__free_hook>: %p\n <win>: %p\n\n", ptr_hook, ptr_win);
  puts("Call <win> function and you'll get the flag.\n");

  A = malloc(0x18);
  B = NULL;
  tcache = (unsigned long*)((unsigned long)A & ~0xfff);

  beginners_heap();
  return 0;
}
