flag = 'ctf4b{stegan0graphy_by_em000000ji}'
result = ''
for c in flag:
    result += chr(ord(c) + 0x1f300)

print(result)
