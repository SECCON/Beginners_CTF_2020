Androidアプリのapkが渡されます。
アプリを開くとFLAGを入れる欄があり、「VALIDATE A」「VALIDATE B」のボタンが提示されています。

「VALIDATE A」のボタンを押すとiOSでアプリを動かすように言われます(が、ipaは配布していないので当然動かせません)。
「VALIDATE B」のボタンを押すとバリデーションに失敗した旨の表示が出てきます。「Learn once, write anywhere」はReact Nativeのキャッチコピーです。

とりあえずapkを展開・逆コンパイルする（作問者はBytecode Viewer経由でapktoolとFernflowerを使用しました）とes.o0i.challengeapp.nativemodule名前空間の下にValidateFlagModuleというクラスが見つかります。validate関数を読んでいくと、あらかじめAES-GCMで暗号化しておいたバイト列を、同一クラス内にある鍵で復号した結果と入力を比較し、一致していればコールバック関数にtrueを渡します。一致しなければfalseを渡します。これを使ってバイト列を復号すると、"1pt_3verywhere}"となり、FLAGの後半が得られます。23文字目以降に(これもValidateFlagModuleの処理からわかります)「1pt_3verywhere}」が現れるように前半を適当に埋めて「VALIDATE B」を押すと、バリデーションに成功した旨が表示されます。

次は「VALIDATE A」についてです。「Learn once, write anywhere」や、逆コンパイル結果で得られるReact系ライブラリのインポートからReact Nativeを使用していることと推測が立ちます。apkのassetsディレクトリにはindex.android.bundleというファイルがありますが、これはReact Nativeで使用されるJavaScriptソースを一つにまとめたものです。難読化をjs-beautifier等で解除してから、アプリ中の文言などを使って関係ある場所を探すと以下が見つかります。

```js
function v() {
var t;
(0, l.default)(this, v);
for (var o = arguments.length, n = new Array(o), c = 0; c < o; c++) n[c] = arguments[c];
return (t = y.call.apply(y, [this].concat(n))).state = {
  flagVal: "ctf4b{",
  xored: [34, 63, 3, 77, 36, 20, 24, 8, 25, 71, 110, 81, 64, 87, 30, 33, 81, 15, 39, 90, 17, 27]
}, t.handleFlagChange = function(o) {
  t.setState({
    flagVal: o
  })
}, t.onPressValidateFirstHalf = function() {
  if ("ios" === h.Platform.OS) {
    for (var o = "AKeyFor" + h.Platform.OS + h.Platform.Version, l = t.state.flagVal, n = 0; n < t.state.xored.length; n++)
      if (t.state.xored[n] !== parseInt(l.charCodeAt(n) ^ o.charCodeAt(n % o.length), 10)) return void h.Alert.alert("Validation A Failed", "Try again...");
    h.Alert.alert("Validation A Succeeded", "Great! Have you checked the other one?")
  } else h.Alert.alert("Sorry!", "Run this app on iOS to validate! Or you can try the other one :)")
}, t.onPressValidateLastHalf = function() {
  "android" === h.Platform.OS ? p.default.validate(t.state.flagVal, function(t) {
    t ? h.Alert.alert("Validation B Succeeded", "Great! Have you checked the other one?") : h.Alert.alert("Validation B Failed", "Learn once, write anywhere ... anywhere?")
  }) : h.Alert.alert("Sorry!", "Run this app on Android to validate! Or you can try the other one :)")
}, t
}
```

iOSで実行した場合は以下の部分が実行されることになります（読みやすくしました）。入力を「AKeyForios10.3」で一文字ずつループさせながらXORした結果がxoredの配列の中身に等しくなればよいので、xoredの中身を「AKeyForios10.3」でXORしてやればしかるべき入力、すなわちFLAGの前半が手に入ります("ctf4b{jav4_and_j4va5cr")。

```js
...
  flagVal: "ctf4b{",
  xored: [34, 63, 3, 77, 36, 20, 24, 8, 25, 71, 110, 81, 64, 87, 30, 33, 81, 15, 39, 90, 17, 27]
...
if ("ios" === h.Platform.OS) {
  var o = "AKeyFor" + h.Platform.OS + "10.3",
      l = t.state.flagVal;
  for (n = 0; n < t.state.xored.length; n++)
    if (t.state.xored[n] !== parseInt(l.charCodeAt(n) ^ o.charCodeAt(n % o.length), 10)) return void h.Alert.alert("Validation A Failed", "Try again...");
  h.Alert.alert("Validation A Succeeded", "Great! Have you checked the other one?")
}
```


