FROM python:3.8

ARG encoded
ENV ENCODED=${encoded}

RUN mkdir -p solver
WORKDIR solver

COPY solver.py /solver

CMD [ "python3", "/solver/solver.py" ]
