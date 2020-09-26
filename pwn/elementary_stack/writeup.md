# Elementary Stack - Writeup

## 方針
1. OOBでbufをatol@got-0x08(=malloc@got)に向ける。
2. bufが更新されたので、readlongの中でatol@gotをprintf@pltに書き換えられる。
3. FSBでlibcのアドレスをリーク。（bufはmalloc@gotを向いているので8バイトまで使える。）
4. readlongの中でatol@gotをsystemに書き換え、`/bin/sh`を送る。

※OOBでリターンアドレスを書き換えても、retに到達するパスが存在しないので単純なROPはできない。
