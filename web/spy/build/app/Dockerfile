FROM python:3.8.2-alpine

WORKDIR /usr/src/app

COPY . .
RUN apk add gcc build-base linux-headers && \
      pip install -r requirements.txt

CMD ["uwsgi", "--ini", "uwsgi.ini"]
