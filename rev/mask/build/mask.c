#include <stdio.h>
#include <string.h>

#define MASK1 0b01110101
#define MASK2 0b11101011

int main(int argc, char **argv) {
	int i;
	char s[64], t[64], u[64];

	if (argc == 1) {
		printf("Usage: ./mask [FLAG]\n");
		return 0;
	}

	strcpy(s, argv[1]);
	int len = strlen(s);

	printf("Putting on masks...\n");
	for (i = 0; i < len; i++) {
		t[i] = s[i] & MASK1;
		u[i] = s[i] & MASK2;
	}

	t[len] = '\0';
	u[len] = '\0';

	printf("%s\n", t);
	printf("%s\n", u);

	if (!strcmp(t, "atd4`qdedtUpetepqeUdaaeUeaqau") && !strcmp(u, "c`b bk`kj`KbababcaKbacaKiacki")) {
		printf("Correct! Submit your FLAG.\n");
	} else {
		printf("Wrong FLAG. Try again.\n");
	}

	return 0;
}
