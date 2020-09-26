# Somen

## 問題文

Somen is so tasty.
<環境への URL>

## 難易度

普通の CTF の Medium くらい。

## 動作

以下のコマンドで起動する。

```sh
cp _env .env
docker-compose up -d --build
```

動作は次のようにして確認できる。

```sh
curl http://localhost:20000/
```

## アイデア

DOM Clobbering による nonce + hash + strict-dynamic な CSP のバイパスと、base tag injection による script ロードの破壊。