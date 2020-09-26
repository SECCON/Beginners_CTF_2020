FROM python:3.7-alpine

WORKDIR /usr/src/app

RUN apk add --no-cache --update-cache socat gcc g++ make libffi-dev openssl-dev \
  && pip install --no-cache-dir pycrypto

COPY solver.py .

CMD ["python", "solver.py"]
