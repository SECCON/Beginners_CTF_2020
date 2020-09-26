FROM python:3.8

ARG flag
ARG format
ENV FLAG=${flag}
ENV FORMAT=${format}

RUN mkdir -p problem
WORKDIR problem

COPY problem.py /problem

CMD [ "python3", "/problem/problem.py" ]
