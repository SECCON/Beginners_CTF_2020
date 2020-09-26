コマンドライン引数にFLAGを与えられるようです。

Usage: ./mask [FLAG]

とりあえず適当な文字列を渡してみると、よくわからない文字列と共に、「FLAGが違うよ」とメッセージが表示されます。

./mask hogehogepiyopiyo123456
Putting on masks...
`eee`eeepaqepaqe101454
hkcahkca`iik`iik!"# !"
Wrong FLAG. Try again.

FLAGになりうる文字列を総当りしてマスクの結果を求めていくのは時間がかかるので、解析の結果とマスク演算の性質を使って元のFLAGを求めます。

逆コンパイラなどを使って解析します。コマンドライン引数でFLAGを受け取ると、各文字の文字コードと `0x75 (0b01110101)`・`0xeb (0b11101011)` のANDをとります(マスク)。`0x75` でマスクした結果が `atd4\`qdedtUpetepqeUdaaeUeaqau` 、`0xeb` でマスクすると `c\`b bk\`kj\`KbababcaKbacaKiacki` となるような文字列がFLAGです。

あるASCII文字コード c を `0x75 (0b01110101)` でマスクすると、ビットが1のところ(下から1, 3, 5-7ビット目)は c の同じ位置のビットがそのまま結果になります。また、ビットが0のところは(下から2, 4, 8ビット目)は c のビットにかかわらず 0 が出力され、出力からは c のその位置のビットが0であったか1であったかはわかりません。

ところが、`0x75 (0b01110101)` で0となっているビットは、もう一つのマスク `0xeb (0b11101011)` ではすべて1となっているので、2つの結果を合わせることで c の文字コードを復元できます。具体的には、マスクした結果の2つの文字のORを取れば元の文字が得られます。

式で表すと以下のようになります。flag_iはFLAGのi番目の文字の文字コードです。

masked_a = flag_i AND 0x75
masked_b = flag_i AND 0xeb
0x75 OR 0xeb = 0b01110101 OR 0b11101011 = 0b11111111

masked_a OR masked_b 
  = (flag_i AND 0x75) OR (flag_i AND 0xeb)
  = flag AND (0x75 OR 0xeb)
  = flag_i AND 0b11111111
  = flag_i

一番下の式の変形(分配法則)はベン図を描くとわかりやすいです。