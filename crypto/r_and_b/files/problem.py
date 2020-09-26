from os import getenv


FLAG = getenv("FLAG")
FORMAT = getenv("FORMAT")


def rot13(s):
    # snipped


def base64(s):
    # snipped


for t in FORMAT:
    if t == "R":
        FLAG = "R" + rot13(FLAG)
    if t == "B":
        FLAG = "B" + base64(FLAG)

print(FLAG)
