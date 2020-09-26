NAME := r_and_b_answer
ENCODED := $(shell cat ../files/encoded_flag)

all:

build:
	DOCKER_BUILDKIT=1 docker build --build-arg encoded="${ENCODED}" -t ${NAME} .

run:
	docker run -it ${NAME}
