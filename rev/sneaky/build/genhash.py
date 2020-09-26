from ptrlib import rol

flag = b'ctf4b{still_ez_2_cheat?}\0\0'

def hash(s):
    h = 0x3141
    h = (h * 0xcafe + s[0]) & 0xffff;
    h = (h * 0xdead + s[1]) & 0xffff;
    return h

l = []
for i in range(0, len(flag), 2):
    l.append(rol(hash(flag[i:i+2]), 1, bits=16))
print(l)
print(len(l))
