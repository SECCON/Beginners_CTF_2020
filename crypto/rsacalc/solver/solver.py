from Crypto.Util.number import *
import binascii
import socket
import os

def recvuntil(s, delim=b'\n'):
    buf = b''
    while delim not in buf:
        buf += s.recv(1)
    return buf

host = os.getenv('CTF4B_HOST', 'localhost')
port = int(os.getenv('CTF4B_PORT', '4444'))

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((host, port))

e = 65537
cmd = b'1337,F,0'
cmdint = bytes_to_long(cmd)
CHR = b'\x02'

recvuntil(s, b'N: ')
N = int(recvuntil(s).strip())

recvuntil(s, b'> ')
s.send(b'1\n')

recvuntil(s, b'> ')
s.send(CHR + b'\n')

recvuntil(s, b': ')
sig1 = int(recvuntil(s).strip(), 16)


recvuntil(s, b'> ')
s.send(b'1\n')
recvuntil(s, b'> ')
s.send(long_to_bytes(cmdint // ord(CHR)) + b'\n')

recvuntil(s, b': ')
sig2 = int(recvuntil(s).strip(), 16)

sig = sig1 * sig2 % N

recvuntil(s, b'> ')
s.send(b'2\n')
recvuntil(s, b'> ')
s.send(cmd + b'\n')
recvuntil(s, b'> ')
s.send(binascii.hexlify(long_to_bytes(sig)) + b'\n')

print(recvuntil(s).decode())
