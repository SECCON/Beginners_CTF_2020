all:
	gcc -Wl,-z,lazy,-z,relro main.c -o chall -fstack-protector -no-pie
	cp ./main.c ../files
	mv ./chall  ../files
