FROM python:3.7-alpine

EXPOSE 31337

WORKDIR /usr/src/app

RUN apk add --no-cache --update-cache socat gcc g++ make libffi-dev openssl-dev \
  && pip install --no-cache-dir pycrypto

COPY server.py params.py ./

CMD ["socat", "tcp-l:31337,fork,reuseaddr", "EXEC:\"python server.py\""]
