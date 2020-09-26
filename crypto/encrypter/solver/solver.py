import requests
import json
import os
import logging
from base64 import b64encode, b64decode
from padding_oracle import padding_oracle

host = os.getenv('CTF4B_HOST', 'localhost')
port = os.getenv('CTF4B_PORT', '10000')
url = 'http://{}:{}/encrypt.php'.format(host, port)

unpad = lambda x: x[:-x[-1]]

sess = requests.Session()

def oracle(data):
    req = {'mode': 'decrypt', 'content': b64encode(data).decode()}
    r = sess.post(url, data=json.dumps(req))
    res = r.json()
    if 'error' in res:
        return False

    return True

if __name__ == '__main__':
    r = sess.post(url, data=json.dumps({'mode': 'getflag'}))
    res = r.json()
    encrypted = b64decode(r.json()['result'])
    plaintext = padding_oracle(encrypted,
                           block_size=16,
                           oracle=oracle,
                           num_threads=8)
    print(unpad(plaintext).decode())
