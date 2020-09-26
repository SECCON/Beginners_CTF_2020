NAME := noisy_equations_solver

all:

build:
	DOCKER_BUILDKIT=1 docker build -t ${NAME} .

run:
	docker run -it --net host ${NAME}

