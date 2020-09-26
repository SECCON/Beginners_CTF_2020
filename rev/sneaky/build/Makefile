all:
#	gcc main.c -o sneaky -O2 -lncurses
	gcc -static main.c -o sneaky -O2 -lncurses -ltinfo
	strip --strip-all sneaky
	mv sneaky ../files
