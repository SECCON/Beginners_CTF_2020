from codecs import encode
from base64 import b64encode
from os import getenv


FLAG = getenv("FLAG")
FORMAT = getenv("FORMAT")


def rot13(s):
    return encode(s, "rot13")


def base64(s):
    return b64encode(s.encode()).decode()


for t in FORMAT:
    if t == "R":
        FLAG = "R" + rot13(FLAG)
    if t == "B":
        FLAG = "B" + base64(FLAG)

print(FLAG)
