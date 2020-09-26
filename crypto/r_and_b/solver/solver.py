from codecs import decode
from os import getenv
from base64 import b64decode


encoded = getenv("ENCODED")


def rot13(s):
    return decode(s, "rot13")


def base64(s):
    return b64decode(s.encode()).decode()


while True:
    if encoded.startswith("ctf4b"):
        break

    if encoded[0] == "R":
        encoded = rot13(encoded[1:])
    elif encoded[0] == "B":
        encoded = base64(encoded[1:])

print(encoded)
