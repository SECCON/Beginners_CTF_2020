from os import getenv
from time import time
from random import getrandbits, seed


FLAG = getenv("FLAG").encode()
SEED = getenv("SEED").encode()

L = 256
N = len(FLAG)


def dot(A, B):
    assert len(A) == len(B)
    return sum([a * b for a, b in zip(A, B)])

coeffs = [[getrandbits(L) for _ in range(N)] for _ in range(N)]

seed(SEED)

answers = [dot(coeff, FLAG) + getrandbits(L) for coeff in coeffs]

print(coeffs)
print(answers)
