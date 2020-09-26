// gcc childheap.c -o childheap
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>

#define BUF_SIZE 0x30

static int menu(void);
static int getnline(char *buf, unsigned size);
static int getint(void);

__attribute__((constructor))
int init(){
	setbuf(stdout, NULL);
	setbuf(stderr, NULL);
	return 0;
}

int main(void){
	int n;
	unsigned size;
	char *content = NULL;

	printf( "Welcome to childheap 2020\n\n"
			"Last year, I was a baby...\n"
			"Now I'm not a baby but a child!!!\n");

	while(n = menu()){
		char buf[4] = {};
		switch(n){
			case 1:
				if(content){
					puts("No Space!!");
					break;
				}

				printf("Input Size: ");
				if((size = getint()) > 0x180){
					puts("Too Big!!");
					break;
				}

				content = malloc(size);
				printf("Input Content: ");
				getnline(content, size);
				break;
			case 2:
				printf("Content: '%s'\nRemove? [y/n] ", content);
				getnline(buf, sizeof(buf)-1);
				if(buf[0] == 'y')
					free(content);
				break;
			case 3:
				content = NULL;
				break;
		}
	}
	return 0;
}

static int menu(void){
	printf(	"\nMENU\n"
			"1. Alloc\n"
			"2. Delete\n"
			"3. Wipe\n"
			"0. Exit\n"
			"> ");

	return getint();
}

static int getnline(char *buf, unsigned size){
	int len;
	char *lf;

	if(!size)
		return 0;

	len = read(0, buf, size);
	buf[len] = '\0';

	if((lf=strchr(buf,'\n')))
		*lf='\0';

	return len;
}

static int getint(void){
	char buf[BUF_SIZE] = {};

	getnline(buf, sizeof(buf)-1);
	return atoi(buf);
}
