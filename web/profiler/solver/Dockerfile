FROM python:3.7-alpine

COPY . .
RUN pip install -r ./requirements.txt

# ENV CTF4B_HOST="profiler.quals.beginners.seccon.jp"
# ENV CTF4B_PORT=80

CMD ["python",  "/solver.py"]
