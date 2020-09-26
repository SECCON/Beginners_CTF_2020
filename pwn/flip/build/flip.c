// gcc flip.c -o flip
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define BUF_SIZE 32

static int getnline(char *buf, int len);
static long getlong(void);

__attribute__((constructor))
int init(){
	setbuf(stdout, NULL);
	setbuf(stderr, NULL);

	return 0;
}

int main(int argc, char *argv[]){
	char* target;
	int idx;

	printf("Input address >> ");
	target = (void*)getlong();

	puts("You can flip two times!");
	for(int i=0; i<2; i++){
		printf("Which bit (0 ~ 7) >> ");
		idx = getlong();

		if(idx > 7){
			puts("invalid bit...");
			return 0;
		}

		*target ^= 1 << idx;
	}
	puts("Done!");

	exit(0);
}

static int getnline(char *buf, int size){
	char *lf;
	long n;

	if(!buf || size < 1)
		return 0;

	fgets(buf, size, stdin);
	if((lf=strchr(buf, '\n')))
		*lf='\0';

	return 1;
}

static long getlong(void){
	char buf[BUF_SIZE] = {};

	getnline(buf, sizeof(buf));
	return atol(buf);
}
