# ChildHeap 2020

## 問題文

nc xxxx.seccon.jp nnnn

## 難易度
中

## 概要
去年の BabyHeap を踏襲  
単一のエントリしか持たないメモ  
追加・削除の機能を持つ  

libc のアドレスのプレゼントは廃止  
利用する libc を 2.27 -> 2.29 とした  

## 脆弱性
- use after free (参照/解放)
- off-by-one null byte
- uninitialized buffer (heap)

