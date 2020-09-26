C4B
===

## 問題文
Are you smart?

[URL]

## 難易度
Hard

## テスト用デプロイ
1. debug.env を .env にコピーする
2. docker-compose.ymlのganacheのコメントを外す
  * 必要があればappの待受ポート等も変更する
3. docker-composeで立ち上げる
4. 手元のMetaMaskの接続先を `[DockerのIP]:8545` にする
5. .envに書いてある `GANACHE_TESTACCOUNT` の秘密鍵をMetaMaskにインポートする

本番はRopstenテストネットワークを使う予定。
