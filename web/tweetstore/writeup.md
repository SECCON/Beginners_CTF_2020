

# 解法

## 1. 配布されるソースコードを読む

- limit句以降にSQL Injectionがあることが確認できる。またDBはpostgresである．
-- ただし，複文実行対策のため";"でsplitし先頭だけがSQL文に連結される．
- DBのユーザ名がFLAGである
-- 方針として，SQL Injectionを用いてどのようにuser名を取ればいいか，となる

## 2. limit句以降のSQL Injectionについて考える

postgresでは，limit句の後ろでもサブクエリが使える．
そこで，以下のように SQL Injectionを行い，検索結果の数を見ることでuserのn番目の文字が何であるか，を特定することができる．

```
.. limit ascii(substr((select user), n, 1);
```

## 3. solverを書く
おわり



# 後日談

作問ミスをしており，searchパラメータ経由でもSQLインジェクションが可能であった．この場合，上記のpostgresに依存した解法ではなく，UNION句を用いた一般的なインジェクションが可能であった．
