NAME := r_and_b
FLAG := $(shell cat ../FLAG)
FORMAT := RBRBRBBRBRBRBB

all:

build:
	DOCKER_BUILDKIT=1 docker build --build-arg flag="${FLAG}" --build-arg format="${FORMAT}" -t ${NAME} .

run:
	docker run -it ${NAME} > ../files/encoded_flag

