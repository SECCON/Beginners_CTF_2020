FROM python:3.7-alpine

RUN apk --no-cache add socat
RUN adduser -D ctf

ADD build/server.py /home/ctf/server/server.py
ADD FLAG            /home/ctf/flag
RUN chmod 500 /home/ctf/server/server.py
RUN chmod 400 /home/ctf/flag
RUN chown ctf:root /home/ctf -R

USER ctf
WORKDIR /home/ctf/server/
CMD socat TCP-L:9712,fork,reuseaddr EXEC:"python3 ./server.py",stderr
