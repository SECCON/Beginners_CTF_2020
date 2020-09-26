# emoemoencode - Writeup

## 方針
UnicodeのU+1F300~U+1F3FFの領域を使ってASCIIコードがエンコードされている。

Unicodeを取得して下8bitを拾ってくっつけるとフラグ。

```python
enc = open('problem.txt', 'r').read()
flag = ''
for c in enc:
    flag += chr(ord(c) & 0xff)
print(flag)
```
