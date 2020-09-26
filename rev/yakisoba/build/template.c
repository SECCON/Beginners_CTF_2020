#include <stdio.h>

int check(const unsigned char *flag) {
  %%%%
}

int main(void) {
  char flag[0x20];
  printf("FLAG: ");
  if (scanf("%31s", flag) == 0)
    return 1;
  if (check(flag)) {
    puts("Wrong!");
  } else {
    puts("Correct!");
  }
  return 0;
}
