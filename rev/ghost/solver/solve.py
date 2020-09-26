with open("../files/output.txt", "r") as f:
    ns = map(int, f.read().strip().split(' '))

mod = 163*397
d = 18151 # modinv(463, 162*396)

x = 1
flag = ""
for i, n in enumerate(ns):
    m = pow(n, d, mod)
    m //= x
    m ^= i + 1
    flag += chr(m)
    x = (n % 128) + 1

print(flag)
