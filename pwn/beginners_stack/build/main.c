#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#define BUFFER_SIZE 0x20
#define READ_SIZE   0x200

void win(void);
void __show_stack(void*);
  
/*----- Read from here -----*/
void vuln(void) {
  char buf[BUFFER_SIZE];
  __show_stack(buf);
  printf("Input: ");
  read(0, buf, READ_SIZE); // vulnerable!
  __show_stack(buf);
}

int main(void) {
  setbuf(stdin, NULL);
  setbuf(stdout, NULL);
  setbuf(stderr, NULL);
  printf("Your goal is to call `win` function (located at %p)\n", &win);
  vuln();
  puts("Bye!");
  return 0;
}
/*-------- to here ---------*/

void win(void) {
  unsigned long rsp;
  asm volatile ("movq %%rsp, %0": "=r"(rsp));
  if (rsp % 0x10 != 0) {
    puts("Oops! RSP is misaligned!");
    puts("Some functions such as `system` use `movaps` instructions in libc-2.27 and later.");
    puts("This instruction fails when RSP is not a multiple of 0x10.");
    puts("Find a way to align RSP! You're almost there!");
    sleep(1);
  } else {
    puts("Congratulations!");
    system("/bin/sh");
  }
  exit(0);
}

void __show_stack(void *ptr) {
  int i;
  unsigned long *p = ptr;

  puts("\n   [ Address ]           [ Stack ]");
  for(i = 0; i < 21; i++) {
    if (i & 1) {
      printf("0x%016lx | 0x%016lx |", (unsigned long)&p[i/2], p[i/2]);
      if (&p[i/2] == ptr)
        printf(" <-- buf");
      else if (&p[i/2] == ptr + BUFFER_SIZE + 0x00)
        printf(" <-- saved rbp (vuln)");
      else if (&p[i/2] == ptr + BUFFER_SIZE + 0x08)
        printf(" <-- return address (vuln)");
      else if (&p[i/2] == ptr + BUFFER_SIZE + 0x10)
        printf(" <-- saved rbp (main)");
      else if (&p[i/2] == ptr + BUFFER_SIZE + 0x18)
        printf(" <-- return address (main)");
      putchar('\n');
    } else {
      puts("                   +--------------------+");
    }
  }
  putchar('\n');
}
