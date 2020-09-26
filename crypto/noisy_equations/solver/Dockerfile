FROM python:3.8

RUN mkdir -p /solver
WORKDIR /solver

COPY requirements.txt /solver
RUN pip install -r /solver/requirements.txt

COPY *.py /solver

CMD [ "python3", "solver.py" ]
