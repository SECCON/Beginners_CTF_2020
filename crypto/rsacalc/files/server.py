from Crypto.Util.number import *
from params import p, q, flag
import binascii
import sys
import signal


N = p * q
e = 65537
d = inverse(e, (p-1)*(q-1))


def input(prompt=''):
    sys.stdout.write(prompt)
    sys.stdout.flush()
    return sys.stdin.buffer.readline().strip()

def menu():
    sys.stdout.write('''----------
1) Sign
2) Exec
3) Exit
''')
    try:
        sys.stdout.write('> ')
        sys.stdout.flush()
        return int(sys.stdin.readline().strip())
    except:
        return 3


def cmd_sign():
    data = input('data> ')
    if len(data) > 256:
        sys.stdout.write('Too long\n')
        return

    if b'F' in data or b'1337' in data:
        sys.stdout.write('Error\n')
        return

    signature = pow(bytes_to_long(data), d, N)
    sys.stdout.write('Signature: {}\n'.format(binascii.hexlify(long_to_bytes(signature)).decode()))

def cmd_exec():
    data = input('data> ')
    signature = int(input('signature> '), 16)

    check = long_to_bytes(pow(signature, e, N))
    if data != check:
        sys.stdout.write('Invalid signature\n')
        return

    chunks = data.split(b',')
    stack = []
    for c in chunks:
        if c == b'+':
            stack.append(stack.pop() + stack.pop())
        elif c == b'-':
            stack.append(stack.pop() - stack.pop())
        elif c == b'*':
            stack.append(stack.pop() * stack.pop())
        elif c == b'/':
            stack.append(stack.pop() / stack.pop())
        elif c == b'F':
            val = stack.pop()
            if val == 1337:
                sys.stdout.write(flag + '\n')
        else:
            stack.append(int(c))

    sys.stdout.write('Answer: {}\n'.format(int(stack.pop())))


def main():
    sys.stdout.write('N: {}\n'.format(N))
    while True:
        try:
            command = menu()
            if command == 1:
                cmd_sign()
            if command == 2:
                cmd_exec()
            elif command == 3:
                break
        except:
            sys.stdout.write('Error\n')
            break


if __name__ == '__main__':
    signal.alarm(60)
    main()
