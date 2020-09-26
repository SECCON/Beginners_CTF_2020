FROM python:3.8

MAINTAINER ShiftCrops <simiyu@shift-crops.net>

RUN pip3 install pwntools && \
	mkdir /opt/sc_expwn

ADD https://raw.githubusercontent.com/shift-crops/sc_expwn/master/sc_expwn.py /opt/sc_expwn/
ADD files /root
ADD solver/exploit.py /root

ENV PYTHONPATH "${PYTHONPATH}:/opt/sc_expwn"
ENV TERM "linux"

WORKDIR /root
CMD ./exploit.py
