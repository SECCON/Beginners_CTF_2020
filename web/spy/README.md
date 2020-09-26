# Spy
## 問題文
As a spy, you are spying on the "ctf4b company".

You got the name-list of employees and the URL to the in-house web tool used by some of them.

Your task is to enumerate the employees who use this tool in order to make it available for social engineering.

{{ URL }}

## 難易度
自明。

## 概要
Account Enumerationの問題。

対象のWebアプリケーションに登録されているアカウントを全て正しく列挙することができればFLAGを表示する。

対象のWebアプリケーションではログイン処理の流れが下記のようになっているため、アカウントが存在する場合と存在しない場合で処理に時間差が発生する。この事実を用いることで、登録されているアカウントの列挙が可能となる。

```
--
input: ID, パスワード
--

アカウント = アカウント取得(ID)
if (アカウントが存在しない) {
    ログイン失敗時の処理()
}

ハッシュ値 = SHA-256(SHA-256( ... SHA-256(ソルト + パスワード) ... )) (*1 million times)

if (ハッシュ値 != アカウント.事前に計算されたパスワードのハッシュ値) {
    ログイン失敗時の処理()
}

ログイン成功時の処理()
```

なお、実際にソルト+ストレッチングの処理を実装すると負荷がかかる上、問題の要件を満たすためには「処理に時間差が発生する」という事実が提示できれば十分なので、実際には`time.sleep(1)`を実行するのみとする。

## 参考
[Once upon a time an account enumeration](https://sidechannel.tempestsi.com/once-upon-a-time-there-was-an-account-enumeration-4cf8ca7cd6c1)