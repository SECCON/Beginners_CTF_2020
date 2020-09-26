NAME := noisy_equations
FLAG := $(shell cat ../FLAG)
SEED := $(shell python -c "from random import getrandbits; print(getrandbits(64))") 

all:

build:
	DOCKER_BUILDKIT=1 docker build --build-arg flag=${FLAG} --build-arg seed=${SEED} -t ${NAME} .

run:
	docker run -it --net host ${NAME}

