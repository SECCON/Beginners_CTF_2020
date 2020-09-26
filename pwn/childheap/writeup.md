# ChildHeap 2020 - Writeup

## 方針
1. off-by-one null byte を利用して 0x100 byte のチャンクを複数作る
2. 上記チャンクを free しまくって tcache を埋め尽くす
3. PREV\_INUSE がリセットされることを利用して，consolidate backward させて一部チャンクをオーバーラップ
4. next を書き換えて main\_arena のアドレス取得
5. free\_hook を system に改竄して shell 奪取

