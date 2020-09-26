FROM python:3.8.2-alpine

WORKDIR /app
ADD . /app

RUN pip install -r requirements.txt

CMD ["python", "/app/solver.py"]