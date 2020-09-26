# syntax=docker/dockerfile:experimental
FROM python:3.8

RUN mkdir -p /problem
WORKDIR /problem

RUN --mount=type=cache,target=/var/cache/apt,sharing=locked --mount=type=cache,target=/var/lib/apt,sharing=locked \
    apt update && \
    apt install -y socat

COPY requirements.txt /problem
RUN pip install -r /problem/requirements.txt

COPY *.py /problem

ARG flag
ARG seed
ENV FLAG=${flag}
ENV SEED=${seed}

CMD [ "socat", "tcp-l:3000,fork,reuseaddr", "EXEC:\"python3 problem.py\"" ]
