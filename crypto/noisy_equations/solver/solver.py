from telnetlib import Telnet
from time import time
from random import getrandbits, seed

import numpy as np


with Telnet("localhost", 3000) as tn:
    A1 = eval(tn.read_until(b"\n"))
    b1 = eval(tn.read_until(b"\n"))

with Telnet("localhost", 3000) as tn:
    A2 = eval(tn.read_until(b"\n"))
    b2 = eval(tn.read_until(b"\n"))

A = (np.array(A1) - np.array(A2)).astype(float)
b = (np.array(b1) - np.array(b2)).astype(float)

x = np.linalg.solve(A, b)
s = "".join([chr(int(c)) for c in x.round()])

if s.startswith("ctf4b{") and s.endswith("}"):
    print(s)

