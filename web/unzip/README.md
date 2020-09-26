# Unzip

## 問題文

Unzip Your .zip Archive Like a Pro.
<環境への URL>

## 難易度

自明問枠(脆弱性が自明なので)。

## 動作

以下のコマンドで起動する。

```sh
docker-compose up -d --build
```

動作は次のようにして確認できる。

```sh
curl http://localhost:10000/
```

## アイデア

.zip アーカイブ中のファイル名には `../` のような文字列を入れられるため、これを使って directory traversal をしてもらう。
