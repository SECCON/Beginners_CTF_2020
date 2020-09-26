#!/usr/bin/env python3
import os

assert os.path.isfile('/home/ctf/flag') # readme

if __name__ == '__main__':
    path = input("File: ")
    if not os.path.exists(path):
        exit("[-] File not found")
    if not os.path.isfile(path):
        exit("[-] Not a file")
    if '/' != path[0]:
        exit("[-] Use absolute path")
    if 'ctf' in path:
        exit("[-] Path not allowed")
    try:
        print(open(path, 'r').read())
    except:
        exit("[-] Permission denied")
