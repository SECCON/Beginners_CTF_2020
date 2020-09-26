# Beginner's Stack
Author: ptr-yudai

## 問題文
`nc {host} {port}`

## 難易度
Beginner

## 概要
初心者なら必ず解けるpwn問。
単純なスタックオーバーフローがあり、シェルを呼び出す関数`win`が与えられている。
リターンアドレスを書き換えて単に`win`を呼べば良いが、libc-2.27はRSPがalignされていないとmovupsで落ちるので、ret gadgetを1つ挟むか`win`のpushを飛ばす。

超初心者向けということで、スタックの様子を可視化した。また、`win`が呼ばれた際にRSPがalignされていないとエラーメッセージを表示するようにした。
