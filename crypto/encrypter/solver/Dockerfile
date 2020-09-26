FROM python:3.8-alpine

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

WORKDIR /usr/src/app

COPY . .

CMD ["python", "./solver.py"]
