#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username $POSTGRES_USER --dbname "$POSTGRES_DB" <<-EOSQL

	create table tweets(id SERIAL, url varchar(2083), text varchar(280), tweeted_at timestamptz);

	create user "$POSTGRES_CTF_USER" with password '$POSTGRES_CTF_PASSWORD';
	grant connect on database $POSTGRES_DB to "$POSTGRES_CTF_USER";
	grant select on tweets to "$POSTGRES_CTF_USER";
	
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1217080911310647296', 'CTF 初心者向けワークショップ「SECCON Beginners」の 2020 年度開催地の募集を開始しました。ぜひご検討ください。#seccon #ctf4b
https://t.co/WxacJblHjb', '2020-01-14 13:47:56');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1187987440717328384', 'すべての講義が終了し、CTF演習が始まりました！今回はWeb、Rev、Cryptoから出題されています。みなさん今回の講義の知識を活かして臨んでいます！ #seccon #ctf4b https://t.co/ac46QTEdJP', '2019-10-26 07:00:52');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1187968233275412481', '3つめの講義はReversingです。x86_64アーキテクチャを対象にアセンブリの基礎を学んだ後、IDAなどを使用しながら実際にバイナリ解析に挑戦していきます。#ctf4b #seccon https://t.co/QGWxhgfkEc', '2019-10-26 05:44:32');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1187962961836433408', '2つめの講義はCryptoです。RSAを題材に、数学的な基礎から学び、実際に手でRSAの暗号化までやってみました。
#seccon #ctf4b https://t.co/ojajYp7nBc', '2019-10-26 05:23:36');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1187932807936409600', 'SECCON Beginners 2019 福岡 最初の講義はWebでした！Webセキュリティの基礎から、XSSまで60分間しっかり講義をしました。 #ctf4b #seccon https://t.co/KKHVzrxGVt', '2019-10-26 03:23:46');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1187881056667496450', '本日 SECCON Beginners 2019 福岡が開催されます！会場は九州大学 伊都キャンパス 情報基盤研究開発センター2階 多目的講義室(最寄りのバス停は九大理学部 or 九大工学部) です。お気をつけてお越しください。 #seccon #ctf4b', '2019-10-25 23:58:08');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1184685870646824961', 'SECCON Beginners 2019 福岡の参加登録締切は 10月20日（日）23:5 JST となります。お忘れなく！ #seccon #ctf4b https://t.co/XKtKT96kuT', '2019-10-17 04:21:36');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1180346104774184960', '午後は Reversing の講義から始まりました。この講義ではアセンブリの読み方を学ぶところから始めて、最終的には簡単な Reversing 問題に挑戦します。講義スライドの末尾では angr の使い方等、発展的な話題も触れられ… https://t.co/DjqfCJiuYT', '2019-10-05 04:56:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1180344327228182529', '本日 SECCON Beginners 2019 東京が開催されています！午前中に開講された Web の講義では、90分間みっちり XSS について学びました。#ctf4b #seccon https://t.co/l8Fm3V1PG3', '2019-10-05 04:49:52');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1179996723462078464', '来る 10/26 (土) に福岡で開催される SECCON Beginners 2019 福岡の参加登録を開始しました。奮ってご応募ください！ https://t.co/XKtKT96kuT #seccon #ctf4b', '2019-10-04 05:48:36');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1178593985121505280', '今回も多数の応募を頂いていた SECCON Beginners 2019 東京についてですが、本日当選者にのみメールで事前配布資料等のアナウンスを配信しました。迷惑メールフォルダを含め、メールボックスをご確認ください。今週末(10… https://t.co/R2tFOBF7q3', '2019-09-30 08:54:37');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1174702546520301574', '2019年10月5日(土)に開催される SECCON Beginners 2019 東京の参加登録を開始しました！奮ってご参加ください。 #seccon #ctf4b https://t.co/QykqzO0D9L', '2019-09-19 15:11:26');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1170229732698275841', 'CTF演習が始まりました！今回、みなさんの解答スピードが早いです！ #ctf4b #seccon https://t.co/lc8e4gL0lr', '2019-09-07 06:58:04');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1170215901947260928', 'Reversingの講義が始まりました！ #ctf4b #seccon https://t.co/EjYOV07yXm', '2019-09-07 06:03:07');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1170211726467289088', 'Cryptoの講義が行われました！ #ctf4b #seccon https://t.co/LSzNnAAWgv', '2019-09-07 05:46:31');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1170211475081707520', 'Webの講義が行われました！ #ctf4b #seccon https://t.co/lIptsaXMAl', '2019-09-07 05:45:31');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1170149169278771200', 'SECCON Beginners 2019 石川、始まりました！本日は Crypto、Web、Reversingの講義が開講されます。 #ctf4b #seccon https://t.co/DkF6OqLi8A', '2019-09-07 01:37:56');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1165204857575198720', '本日はSECCON Beginners 2019 苫小牧にご参加頂きありがとうございました！次は金沢でお会いしましょう！ #SECCON #ctf4b', '2019-08-24 10:11:01');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1165157267064311808', 'すべての講義を終了し、CTF演習が始まりました！ #seccon #ctf4b https://t.co/6hINaBVNmt', '2019-08-24 07:01:54');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1165136263898206210', 'Webの講義が始まりました！ #ctf4b #seccon https://t.co/YTL7GK63ON', '2019-08-24 05:38:27');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1165120450529087488', '予定を変更し、Reversingの講義が始まりました！ #ctf4b #seccon https://t.co/8Wq3Br0PZJ', '2019-08-24 04:35:36');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1165088736033398784', 'Cryptoの講義ではユークリッドの互除法、拡張ユークリッドの互除法、逆元など基本的なところから解説を行っています。 #seccon #ctf4b', '2019-08-24 02:29:35');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1165084629725921282', 'Cryptoの講義が始まりました！ #seccon #ctf4b https://t.co/7WY3uqYOaB', '2019-08-24 02:13:16');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1165076904375771136', 'SECCON Beginners 2019 北海道、始まりました！本日は Crypto・Web・Reversing の 3 つの講義が開講されます。#ctf4b #seccon https://t.co/4FKLU8Llk9', '2019-08-24 01:42:34');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1161771179490697217', 'いよいよ開催が来週に迫った SECCON Beginners 2019 苫小牧ですが、現在席にゆとりがあるため、引き続き参加登録を受け付けています。ぜひお誘い合わせの上ご参加ください。#ctf4b #seccon https://t.co/GJWx4vjB6i', '2019-08-14 22:46:48');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1159108985938513921', 'SECCON Beginners 2019 苫小牧への登録はお済みですか？締切は 8/12 (月) 23:59 JST となっております。奮ってご参加ください！https://t.co/GJWx4vjB6i #ctf4b #seccon', '2019-08-07 14:28:12');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1151327695969476613', '苫小牧開催の登録開始と併せて、9/7 (土) に開催される SECCON Beginners 2019 石川の参加登録も開始しました。こちらは 8/25 (日) が登録締切です。ご登録をお待ちしております。#ctf4b… https://t.co/7gVJxwcYU2', '2019-07-17 03:08:08');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1151327205042974721', '来る 8/24 (土) に開催予定の SECCON Beginners 2019 苫小牧の参加登録を開始しました。登録の締切は 8/12 (月) となっております。奮ってご参加ください！#ctf4b #seccon
https://t.co/GJWx4vjB6i', '2019-07-17 03:06:10');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1137583938493329408', '全てのスコアサーバー、問題サーバーをシャットダウンしました。ご参加いただいた皆様ありがとうございました！ #seccon #ctf4b', '2019-06-09 04:55:20');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132546180150321152', '今回の問題サーバー・スコアサーバーは、6/9 (日) 12:00 JST 以降に順次停止作業を行います。それまでの期間は Writeup の準備や確認などにご利用いただけます。競技時間中とは異なり、サービスの動作保証や質問対応は致… https://t.co/EYloF4YzlM', '2019-05-26 07:17:05');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132528254365999105', 'SECCON Beginners CTF 2019 は 1 位が zer0pts、2 位が superflip、3  位が LEADER!! という結果に終わりました。おめでとうございます！#ctf4b #seccon https://t.co/ihTkU8nTga', '2019-05-26 06:05:51');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132527672637083649', '終了後のアンケートにご協力ください！アンケートはスコアサーバにリンクがありますのでご確認ください。 #ctf4b #seccon', '2019-05-26 06:03:33');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132527202610794496', 'お疲れ様でした！ #ctf4b #seccon', '2019-05-26 06:01:41');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132525959658188800', '残り5分です！ #seccon #ctf4b', '2019-05-26 05:56:44');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132521234132946944', '終了まであと30分を切りました！最後までBeginners CTFをお楽しみください！ #seccon #ctf4b', '2019-05-26 05:37:58');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132513813662535680', '長かったCTFですが、あと残り1時間となりました！スコアサーバにて今回のCTFについてアンケートのご協力をお願いしております。お手すきの際にご回答よろしくお願いいたします！ #seccon #ctf4b', '2019-05-26 05:08:28');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132482626571259909', 'あと3時間となりました！ #ctf4b #seccon', '2019-05-26 03:04:33');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132451320235868160', 'RT @ctf4b: 問題 secconpass は不具合によりフラグの一部を正常に得ることができないため、提出されたフラグの先頭 30 文字が正しければ、正解とします。フラグを一度送信しているものの正答とならなかった場合には、再度送信してください。ご迷惑をおかけしてしまい申し…', '2019-05-26 01:00:09');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132364002246373376', '問題 secconpass は不具合によりフラグの一部を正常に得ることができないため、提出されたフラグの先頭 30 文字が正しければ、正解とします。フラグを一度送信しているものの正答とならなかった場合には、再度送信してください。ご… https://t.co/6p5vqe000M', '2019-05-25 19:13:11');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132330386321580033', 'Cryptジャンルで使用しているサーバーの増強が完了しました。今後さらにサーバー負荷が上がると不必要に大量に投げているリクエストのブロックを行います。 #ctf4b #seccon', '2019-05-25 16:59:36');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132327657826205696', 'Cryptジャンルで使用しているサーバーの増強のため、Cryptサーバを数分間シャットダウンします。復旧次第、再度アナウンスを行う予定です。 #seccon #ctf4b', '2019-05-25 16:48:45');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132272479114747904', 'katsudonについて､問題にミスがあることが判明いたしました。修正版をkatsudon-okawariとして追加致しました。この度は申し訳ありませんでした。 #ctf4b #seccon', '2019-05-25 13:09:30');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132268622615678978', '問題にKatsudon-okawariが追加されました！ #ctf4b #seccon', '2019-05-25 12:54:10');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132267512886833157', 'Statusバッジは正常に動作するようになりました！ #ctf4b #seccon', '2019-05-25 12:49:46');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132246657075109888', '現在Statusバッジのシステムのメンテナンスを行っております、failed表示になっていますが、問題は正常に可動していますのでご確認ください。 #seccon #ctf4b', '2019-05-25 11:26:53');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132166609852698625', '今回から、サーバーが必要な問題に対して参加者のみなさんが見える形で「サーバーにアクセスできるか」と「作問者が用意しているソルバーで現在問題が解けるか」を見えるようにしています！ぜひ参考にしてください！  #seccon #ctf4b https://t.co/5jmqIVr7WQ', '2019-05-25 06:08:49');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132166123485421568', 'Webの問題に「Himitsu」と「Katsudon」を追加しました！ #ctf4b #seccon', '2019-05-25 06:06:53');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1132165468251279361', '競技開始しました！24時間頑張ってください！ #ctf4b #seccon https://t.co/0nOiZJ1pl3', '2019-05-25 06:04:16');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1131769932570435584', 'SECCON Beginners CTF 2019 のスコアサーバーをオープンしました。競技は 5/25 (土) 15:00 から開催されます。https://t.co/FzzxzTba3J  #ctf4b #seccon https://t.co/S3Gt8gpXxQ', '2019-05-24 03:52:33');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1131525811591192576', 'SECCON Beginners CTF 2019 のスコアサーバーは 5/24 (金) 12:00 以下の URL にてオープンの予定です。登録は当日の競技時間中も可能です。今しばらくお待ちください！ https://t.co/FzzxzTba3J #ctf4b #seccon', '2019-05-23 11:42:30');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1126088247325519872', 'SECCON Beginners CTF 2019 開催まで残り 2 週間弱となりました。開催は 5/25 (土) 15:00 から 24 時間の予定です。SECCON 公式 Web サイトにも案内が掲載されましたので、ぜひご確認… https://t.co/AX08h9cGDI', '2019-05-08 11:35:34');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1115791960495644672', 'SECCON BeginnersCTF 2019の開催が5/25(土) 15:00 ~ 5/26(日) 15:00(JST) で開催されます、登録方法や詳細なアナウンスなどは後ほど公開いたします。皆様奮ってご参加ください！… https://t.co/ihI68pPwap', '2019-04-10 01:41:48');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1066280531661217793', 'CTF 演習を終え、SECCON Beginners 2018 名古屋は無事終了しました。ご参加いただいた皆様、ありがとうございました！ #ctf4b #seccon', '2018-11-24 10:41:03');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1066226547428777984', '本日 3 つ目は Reversing の講義です。実際にアセンブリを眺めたり、書いたりしながら、プログラムの動作の仕組みや解析の仕方を学んでいきます。 #ctf4b #seccon https://t.co/uj9sodukRk', '2018-11-24 07:06:33');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1066186822479339521', '休憩の後、Web の講義が始まりました。XSS を主題としたこの講義では、実際に手を動かしながら XSS の原理・手法・CTF での出題例を学びます。 #ctf4b #seccon https://t.co/CDcAtnVjEE', '2018-11-24 04:28:41');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1066164109253595136', '午前中は Crypto の講義が行われています。この講義は暗号や RSA の基礎・原理を学ぶところからスタートし、最終的には RSA の脆弱な場合の復号にチャレンジします。 #ctf4b #seccon https://t.co/2GgUYeWp1r', '2018-11-24 02:58:26');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1066163865161879552', '初めにセキュリティを学ぶ際の倫理面について学んだのち、CTF とは一体どのようなものか・どう取り組んでいくと良いかについての紹介をしました。 #ctf4b #seccon https://t.co/E0lYTeAuO3', '2018-11-24 02:57:28');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1066161209315999744', '本日は SECCON Beginners 2018 名古屋が開催されています! 今回も 1 日を通して Crypto / Web / Reversing の 3 つのジャンルを学んだ後、演習として CTF に挑戦する予定です。… https://t.co/RZmg7fBwlw', '2018-11-24 02:46:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1058309348588380160', '11/24 (土) 開催 SECCON Beginners 2018 名古屋の参加登録が始まりました。奮ってご応募ください! https://t.co/BdXjayENuJ #ctf4b #seccon', '2018-11-02 10:46:25');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048472221650276353', 'NEXTでは、引き続きfuchuのpwnパートの解説が始まっています。 #ctf4b #seccon https://t.co/TtfZTWEBDC', '2018-10-06 07:17:11');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048462520208392192', '正しくはSECCON CTF 2017の国内決勝の問題でした。失礼いたしました。 https://t.co/WFdkkL3wfh', '2018-10-06 06:38:38');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048459050478788609', '通常開催では最後のプログラムであるCTF演習が始まりました。 #seccon #ctf4b https://t.co/SSafwrKktt', '2018-10-06 06:24:51');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048449496722489345', 'NEXTでは後半のWeb＋Pwnの講義が始まりました。この講義では、SECCON CTF 2018の国内決勝で出題されたfuchuを解説していきます。 #seccon #ctf4b https://t.co/5PU4uQp7GC', '2018-10-06 05:46:53');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048435022456283136', 'なお、Pwn講義はBeginners史上初の開催です！ #ctf4b #seccon https://t.co/wvYIfTnJzs', '2018-10-06 04:49:22');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048432252374343681', '後半のpwnの講義がスタートしました！まずは、pwnについての説明、pwnに必要なバイナリ解析の入門を行います。 #ctf4b #seccon https://t.co/vSTCgIoDan', '2018-10-06 04:38:22');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048427217972125701', 'NEXTでは現在Reversingの講義中です。Reversingに関する知識の説明と事前に配った問題の解説を中心に講義を行います。 #ctf4b #seccon https://t.co/WE2MJyvpog', '2018-10-06 04:18:22');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048420974771617794', 'Beginners NEXT始まりました！
こちらもほぼ満員で多くの方にご来場頂いています。ありがとうございます！ #ctf4b #seccon https://t.co/1yH47TrUvr', '2018-10-06 03:53:33');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048414772809490432', 'Beginners NEXTの会場にも徐々に参加者が集まり始めました！ #ctf4b #seccon https://t.co/0ew5a2FWgj', '2018-10-06 03:28:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048400435399127040', '講義では公開鍵暗号の理解に必要な剰余計算について理解を深めるため、剰余計算を実際に解いてみる演習中です。 #ctf4b #seccon https://t.co/fslxk27gJ7', '2018-10-06 02:31:56');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048393902904901632', 'Cryptoの講義が始まりました！まずはCryptoとは何か、どのようなスキルが必要なのか、暗号の種類を説明しています。 #ctf4b #seccon https://t.co/cEoLc253Cm', '2018-10-06 02:05:59');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048387292669173762', 'SECCON Beginners 東京の通常講義始まりました！今回も多くの方にご来場頂いております。ありがとうございます！ #ctf4b #seccon https://t.co/yj2NLb4pQq', '2018-10-06 01:39:43');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048382171591786497', 'おはようございます！今回のBeginnersでは会場提供をしてくださっている産業技術高専の学生さんに構築していただきました！！ #ctf4b #seccon https://t.co/rAr1StcPWB', '2018-10-06 01:19:22');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1048146238875394049', 'いよいよ明日に迫ってまいりましたSECCON Beginners Tokyo 2018とNEXTの開催は東京都立産業技術高専の品川キャンパスでの開催となります。最寄り駅は品川シーサイド駅です。荒川キャンパスではないのでご注意ください！ #ctf4b #seccon', '2018-10-05 09:41:51');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1046664642854932480', '本日SECCON Beginners 2018東京 (NEXT含む)について、当選の連絡をメールいたしました。メールを必ずご確認頂き、ご準備いただくようお願い致します。メールが届いていない方は落選となりますが、キャンセルが複数出た… https://t.co/7JcL7lkI8D', '2018-10-01 07:34:31');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1035905883924918272', 'SECCON Beginners 2018 広島、無事終了致しました。皆様ありがとうございました! #seccon #ctf4b', '2018-09-01 15:03:03');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1035788174012776449', '3 つの講義を終えた後は CTF 演習です。妨害コンテンツとも戦いながら、黙々と取り組んでいます。 #seccon #ctf4b', '2018-09-01 07:15:19');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1035762062784229377', 'Web に続き、Reversing の講義が行われています。本講義はアセンブリ入門からスタートして、最終的には関数呼び出し時のスタック周りの挙動までを追いかけていきます。 #seccon #ctf4b https://t.co/gS63WWbW97', '2018-09-01 05:31:33');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1035744343091363840', '現在は Web の講義が行われています。この講義では、XSS についてを初歩から掘り下げていきます。 #seccon #ctf4b https://t.co/Nc1M41dZ7m', '2018-09-01 04:21:09');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1035742950234284032', '午前中には Beginners 初めての Crypto の講義が行われました。数学の基礎から RSA 暗号の仕組み、その他公開鍵暗号の概観などまでを 60 分に凝縮して学びました。  #seccon #ctf4b https://t.co/miqllsInqH', '2018-09-01 04:15:37');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1035680825193332736', '本日 SECCON Beginners 2018 広島が開催されます! あいにくのお天気ですので、足元にはお気をつけてご来場ください。#seccon #ctf4b', '2018-09-01 00:08:45');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1031501070923522048', 'SECCON Beginners 2018 広島にご応募いただいた皆様のうち、厳正なる抽選の結果当選した方にのみ、本日メールを送付させていただきました。席数の都合上落選となってしまった皆様につきましては、ご希望に添えず申し訳ござい… https://t.co/a9c1zfDVXS', '2018-08-20 11:19:54');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1026675746050166785', '2018 年度初の SECCON Beginners を 9月1日(土) 広島にて開催致します。奮ってご応募ください! #ctf4b #seccon https://t.co/NqgFugeU69', '2018-08-07 03:45:47');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1013083326767218689', 'アナウンスの通り本日Beginners CTF 2018の問題・スコアサーバを停止いたします。多くの方のご参加ありがとうございました！ #ctf4b #seccon', '2018-06-30 15:34:21');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1011177255387533312', 'Beginners CTF 2018のスコアサーバーと問題サーバーは7/1に停止します。 #ctf4b #seccon', '2018-06-25 09:20:19');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000594232531300352', '@srpkdyy お問い合わせありがとうございます。こちらがご回答となります。ご確認ください。
https://t.co/h2KYrCgIb2', '2018-05-27 04:27:09');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000592263301316608', 'スコアサーバおよび問題サーバは競技終了後から1ヶ月稼働させておく予定ですが、予告なく稼働停止をする場合もございます。ご了承ください。 #ctf4b #seccon https://t.co/HPGac23qBW', '2018-05-27 04:19:20');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000590110772027392', 'なお、スコアサーバおよび問題サーバはしばらく稼働させておく予定です。Writeupの作成や解けなかった問題の復習にご活用ください #seccon #ctf4b', '2018-05-27 04:10:47');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000587610014662657', 'SECCON Beginners CTF、これにて競技終了となります。参加された皆様、お疲れ様でした！ #ctf4b #seccon', '2018-05-27 04:00:50');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000573545305456640', '残り一時間です！ #ctf4b #seccon', '2018-05-27 03:04:57');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000374141680156672', 'チームTWDBがすべての問題を制覇しました！おめでとうございます！ #seccon #ctf4b https://t.co/bgJBRv3OXR', '2018-05-26 13:52:36');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000300866316910592', '先ほど、全ての問題が１チーム以上に解かれました。引き続き皆さん頑張ってください！ #ctf4b #seccon', '2018-05-26 09:01:25');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000250690696921088', '現在スコアサーバへのある1IPからDDoSが試行されているように見えます。引き続き試行が繰り返される場合はブロックを行いますのでご注意ください。 #ctf4b #seccon', '2018-05-26 05:42:03');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000237510386044928', 'てけいさんえくすとりーむずですが、一時的にサーバダウンしていました。現在では接続可能になっております。 #ctf4b #seccon', '2018-05-26 04:49:40');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000227246525759489', '申し訳ありません。 [Warmup] Veni, vidi, vici のフラグ登録が間違っていました。解けている方は再度submitをお願いします。 #ctf4b #seccon', '2018-05-26 04:08:53');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000225516681494528', 'BeginnersCTF 2018開始いたしました！ #ctf4b #seccon', '2018-05-26 04:02:01');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000222556534620161', 'まもなくBeginnersCTF 2018が開催になります！IRCに参加されていない方は、そちらのご参加もよろしくお願いします！ #seccon #ctf4b', '2018-05-26 03:50:15');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000163643772710912', '本日13:00よりBeginnersCTF 2018を開催致します！開始直後はアクセス増加に伴いスコアサーバへの接続が難しくなる場合もございます。まだ、登録がお済みでない方はお早めにご登録ください！また、登録済みの方はルールに記載… https://t.co/oMZLTmFQqU', '2018-05-25 23:56:09');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/1000007442770505728', 'Beginners CTF 2018のルールをアップデートしました。競技に関するお問い合わせはIRCでのみとさせて頂きます。IRCへのアクセスはルールページをご確認ください。 #seccon #ctf4b', '2018-05-25 13:35:28');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/999651710090985472', 'RT @ctf4b: 他にもメールが配送されていない方は、再度お使いのメールサービスで https://t.co/VPV0SrJRw6 からのメールを受け取れるようになっているかご確認ください。もし、解決が難しい用であればご登録時に利用したメールアドレスから info2018@…', '2018-05-24 14:01:54');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/999651696920870912', 'RT @ctf4b: メールに関しましてご迷惑をおかけしております。メールが届いていない方でhttps://t.co/04gCBrYKUt などを利用している方に関しましては、受信セーフリストに追加の上パスワードの再設定から再度ステップを行って頂ければと思います。 #secco…', '2018-05-24 14:01:51');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/999181507544547328', '他にもメールが配送されていない方は、再度お使いのメールサービスで https://t.co/VPV0SrJRw6 からのメールを受け取れるようになっているかご確認ください。もし、解決が難しい用であればご登録時に利用したメールアドレ… https://t.co/807KME6V94', '2018-05-23 06:53:29');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/999180556838490112', 'メールに関しましてご迷惑をおかけしております。メールが届いていない方でhttps://t.co/04gCBrYKUt などを利用している方に関しましては、受信セーフリストに追加の上パスワードの再設定から再度ステップを行って頂ければと思います。 #seccon #ctf4b', '2018-05-23 06:49:43');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/998414961742200832', '大変お待たせいたしました、本日よりSECCON Beginners CTF 2018の参加登録を開始いたしました。皆様のご参加心よりお待ちしております！ #seccon #ctf4b https://t.co/SaMzP4cSYQ', '2018-05-21 04:07:31');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/997837487362920452', 'お待たせして大変申し訳ありません。SECCON Beginners CTFへの登録開始も延期となって大変申し訳ありませんが、登録は5/21(月)より開始いたします。競技日時に変更はございません。お待ち頂いている皆さんにはご迷惑をお… https://t.co/dKipN0Vkrs', '2018-05-19 13:52:50');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/992021944571645952', '大変お待たせいたしました！SECCON Beginnersとしては初のオンラインCTFを5/26(土) 13:00より24時間開催いたします！チームでの参加となりますが、お1人でのご参加も大歓迎です！レジストレーションに関しまして… https://t.co/7Sr05EPqPY', '2018-05-03 12:43:57');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/969091806376869888', '先日ご案内した SECCON Beginners CTF ですが、諸般の事情につき延期させていただきます。
新しい日程に関しましては再度ウェブサイトおよび twitter にてご案内いたします。
ご迷惑をおかけいたしますがご理解の… https://t.co/3sYKlEnXwf', '2018-03-01 06:07:46');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/968014271404638208', 'みなさま大変長らくお待たせいたしました。
SECCON Beginners CTFを3/17（土）に開催いたします。オンラインからも参加可能となる予定ですので、是非ご参加ください。なお、登録につきましてはもう少々お待ち下さい。… https://t.co/HaF6kBJAoC', '2018-02-26 06:46:01');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/957603371048824832', 'SECCON Beginnersでは、2017年度に引き続き2018年度の開催地を募集しています。開催地として立候補可能な団体様いらっしゃいましたら以下のサイトをご覧頂き、ご連絡ください。
#ctf4b #seccon
https://t.co/w0zKhiCkCr', '2018-01-28 13:16:49');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936867119966199809', '皆さんCTF演習集中しています…… #seccon #ctf4b https://t.co/R9XaKsI0oE', '2017-12-02 07:58:22');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936845869743083521', 'CTFオリエンテーションが始まりました！！ #seccon #ctf4b https://t.co/4U697KmIVg', '2017-12-02 06:33:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936828461456302080', 'Binaryの講義では、実行ファイルに注目しそれがどのように動作しているのかに迫っていきます。 #seccon #ctf4b', '2017-12-02 05:24:45');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936827944726536192', 'Binaryの講義が始まりました！ #seccon #ctf4b https://t.co/p5Cn7NcIRP', '2017-12-02 05:22:42');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936818186380443648', 'Forensicsの講義では、Wiresharkの使い方から実際のパケット内にある様々なプロトコルのデータに注目し、隠されたデータを探します。 #seccon #ctf4b', '2017-12-02 04:43:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936810001011634176', '今回のSECCON Beginners 長崎開催でもスポンサーのYAMAHA様より機材提供を頂き、ネットワークの提供を行っています！！ #seccon #ctf4b https://t.co/4tsYfIrIFq', '2017-12-02 04:11:23');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936809211660341249', 'Forensicsの講義が始まりました！！ #seccon #ctf4b https://t.co/qcN03o9mCP', '2017-12-02 04:08:15');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936784053088301056', 'Webの講義では、GoogleChromeに存在する開発者ツールを利用し、HTTPの様々なログを閲覧する方法について学び、その後Webアプリケーションの脆弱性にフォーカスし知識を深めます。 #seccon #ctf4b', '2017-12-02 02:28:17');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936778661973258240', 'Webの講義が始まりました！ #seccon #ctf4b https://t.co/aVoc9FBaWs', '2017-12-02 02:06:52');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936775388289581056', 'オリエンテーションが始まりました！倫理やCTFとはに注目して、CTFのみならず重要なことについて共有を行っています。 #seccon #ctf4b https://t.co/ZfO2xHxAMF', '2017-12-02 01:53:51');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936761497388060672', '受付開始しました！ #seccon #ctf4b https://t.co/IlO73asQw4', '2017-12-02 00:58:39');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/936759587633426432', '本日は SECCON Beginners 長崎開催です！会場の長崎県立大学シーボルト校でお待ちしております！！ #seccon #ctf4b https://t.co/JSyaxWhtag', '2017-12-02 00:51:04');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931783563690655744', 'CTFオリエンテーションが始まりました！ #ctf4b #seccon https://t.co/4H0kHYWJBD', '2017-11-18 07:18:07');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931762714988355585', 'Reversiongの講義では、アセンブリ言語の各命令に注目しどのような意味を持つのかを理解し、数命令ほどの処理でレジスタにどのような値が入っているのかを演習を通して理解を目指します。 #seccon #ctf4b', '2017-11-18 05:55:17');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931759234743545858', '今回もSECCON BeginnersではYAMAHA様の機材提供を頂いて演習環境のネットワークを提供しています。いつも柔軟かつ迅速に構築が出来るので非常に助かっております！ #ctf4b #seccon https://t.co/2KsiQGIhKw', '2017-11-18 05:41:27');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931755756696756225', 'Reversingの講義が始まりました！ #ctf4b #seccon https://t.co/nPepS1YqPR', '2017-11-18 05:27:38');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931745701364563968', 'Forensicsの講義では、Wiresharkの使い方からパケットがどのような形で保存されどのように効率的に調べていくかについて講義・演習を行っています。 #ctf4b #seccon', '2017-11-18 04:47:40');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931710538362511360', 'Webの講義ではブラウザが持つ開発者ツールを利用し、ブラウザとサーバ間の通信を見てどのようなやり取りをしているのかなど表面上は見えない情報に注目し解説を行っています。 #ctf4b #seccon', '2017-11-18 02:27:57');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931705271608422400', 'Webの講義が始まりました！ #ctf4b #seccon https://t.co/1HafzjDqwc', '2017-11-18 02:07:01');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931703159659020288', '初めのオリエンテーションでは、セキュリティなどの様々な技術を利用し問題を解いていくCTFとはどのような競技であるのか、またなぜセキュリティエンジニアは不正をしないのかをテーマに倫理の問題について解説を行っています。 #ctf4b #seccon', '2017-11-18 01:58:38');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931700355460501504', '開催しました！西川さまよりご挨拶をいただきました #ctf4b #seccon https://t.co/uLRhjwTt5M', '2017-11-18 01:47:29');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931700049280614400', '間もなく開催です！ #ctf4b #seccon https://t.co/0thKEa4ANZ', '2017-11-18 01:46:16');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/931632107675250688', '本日はSECCON Beginners鹿児島開催です！雨でお足元が悪い中だとは思いますが、本日お越しの方はお気をつけてご来場いただければと思います！ #ctf4b #seccon', '2017-11-17 21:16:18');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/927782038479843328', '11/18に開催されるSECCON Beginners鹿児島の締め切りが迫っています。ご応募お待ちしております！https://t.co/DoVuc5u7ji', '2017-11-07 06:17:29');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/927781298038390784', 'SECCON Beginners長崎を12月2日（土）に開催します！奮ってご応募ください！ https://t.co/A6uGVw9NYb', '2017-11-07 06:14:33');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/921318916005040129', 'SECCON Beginners鹿児島を11月18日（土）に開催します！奮ってご応募ください！https://t.co/DoVuc5u7ji', '2017-10-20 10:15:21');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916567773848809472', '今回のBeginners CTFでは高専の方にも問題を作って頂いています！壇上ではヒントを聞いています！！ #ctf4b #ctf4bnext #seccon https://t.co/6XWq6vTQBe', '2017-10-07 07:36:00');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916563465795362816', 'SECCON Beginnersの受付では、SECCONグッズの販売を行っています！ステッカー、パーカー、Tシャツがございます。記念に宜しければどうぞ！ #ctf4b #ctf4bnext #seccon https://t.co/CQ46kO5S2N', '2017-10-07 07:18:53');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916562552703688710', 'Beginners NEXTは、ちょうど休憩中です！演習に取り組む方、雑談する方様々な方がいらっしゃいます！ #ctf4b #ctf4bnext #seccon https://t.co/spEt4yLKgQ', '2017-10-07 07:15:16');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916559064787595264', '今回のトーク担当です #ctf4b #ctf4bnext #seccon https://t.co/g5tDB8Ev7d', '2017-10-07 07:01:24');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916558823300595712', '皆さん集中しています！！ #ctf4b #ctf4bnext #seccon https://t.co/PAYHS8S2nJ', '2017-10-07 07:00:26');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916554429997785089', 'Beginnersでは、CTFが始まりました！！！ #ctf4b #ctf4bnext #seccon https://t.co/0SKP4y0apN', '2017-10-07 06:42:59');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916553657432252416', 'Beginnersでは、間もなく開始するCTFについてCTFをどう進めるか、またCTF中の諸注意を確認しています。 #ctf4b #ctf4bnext #seccon https://t.co/l7fMmpCAkn', '2017-10-07 06:39:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916552969780322304', 'Beginners NEXTのPwn講義では、heapのデータ構造に注目し、利用中や解放済みの時には構造が異なる事など順を追って学んでいきます。 #ctf4b #ctf4bnext #seccon', '2017-10-07 06:37:11');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916551250363031552', 'Beginners NEXTでは、Pwnの講義が始まります！ #ctf4b #ctf4bnext #seccon https://t.co/ETlXymW2yS', '2017-10-07 06:30:21');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916540484830371840', 'Beginners NEXTのWebの演習では、リクエストの間に入っているWAFをなんとか回避できないか挑戦をしています。 #ctf4b #ctf4bnext #seccon', '2017-10-07 05:47:34');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916539075322011649', 'Beginners NEXTでは、ちょうどWeb講義の休憩に入りましたが、皆さん休憩中にも関わらず演習に集中しています……。皆さんの好きなタイミングでしっかり休憩を取っていただければと思います！
#ctf4b #ctf4bnext… https://t.co/OFVmalXaUK', '2017-10-07 05:41:58');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916537372778545152', 'BeginnersのReversing講義では、実は身近な存在であるバイナリについて解説し、マジックナンバーや様々なデータの形を紹介した上で、機械語の存在、アセンブリ言語に触れていきます。 #ctf4b #ctf4bnext #seccon', '2017-10-07 05:35:12');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916535932244733952', 'Reversingの講義が始まりました！！ #ctf4b #ctf4bnext #seccon https://t.co/rJsccFlpoV', '2017-10-07 05:29:29');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916534742979289088', '会場ネットワーク担当による講義が休み時間に行われてました #ctf4b #seccon https://t.co/RV7lxVLGI0', '2017-10-07 05:24:45');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916522988698165248', 'スタッフはみなさんを食べたりしないので、質問がある方はぜひ手を上げてご質問ください！　#ctf4b #seccon', '2017-10-07 04:38:03');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916520446362755072', 'BeginnersのForensicsの講義では、WireSharkの使い方を触れた上で実際にpcapファイルに隠されたflagを探す演習を行っています。 #ctf4b #ctf4bnext #SECCON', '2017-10-07 04:27:57');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916518416462123008', 'Beginners NEXTのWeb講義では激辛curlと題する講義が始まりました。Web問題の現状のCTFにおける位置づけを理解した上で、問題サーバーに対する自分の攻撃が成り立つかなどの切り分けを行える事などを目標に、Begin… https://t.co/X0INq7ZKY2', '2017-10-07 04:19:53');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916515977554157568', 'Beginners NEXTでは、Webの講義が始まりました！！ #ctf4b #ctf4bnext  #seccon https://t.co/57LzBdHxa0', '2017-10-07 04:10:11');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916514641794162688', 'BeginnersではForensicsの講義がはじまりました！！ #ctf4b #ctf4bnext  #seccon https://t.co/tiny9xNV1d', '2017-10-07 04:04:53');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916512136171470849', 'SECCON Beginners NEXT多くの方にご参加頂いています！ #ctf4b #ctf4bnext  #seccon https://t.co/9JzOUEcanV', '2017-10-07 03:54:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916511680237940738', 'SECCON Beginners NEXTのオリエンテーション始まりました！！ #ctf4b #ctf4bnext  #seccon https://t.co/8GsFcsPjk5', '2017-10-07 03:53:07');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916510757218484224', '今回はインフラスポンサーであるさくらインターネット様の、さくらのクラウドにVPN接続して演習環境を提供していただいてます。 #ctf4b #ctf4bnext  #seccon', '2017-10-07 03:49:27');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916510103301267456', '今回もSECCON Beginners東京、Beginners NEXTではYAMAHA様の機材提供を頂きまして、ネット環境を提供しています！ #ctf4b #ctf4bnext  #seccon https://t.co/C6elie2fes', '2017-10-07 03:46:51');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916509319843127296', 'SECCON Beginners NEXT間もなく開催です！！ #ctf4b #ctf4bnext  #seccon https://t.co/qWaxxUxKxu', '2017-10-07 03:43:44');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916489735379460096', 'SQLインジェクションについて、それが発生してしまう文字列を知り、特別な意味を持つ文字を駆使しどのようにSQL文を書き換えられるのかについて講義を行っています。 #ctf4b #ctf4bnext #seccon', '2017-10-07 02:25:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916489103662710784', 'Webの講義を行っています！ #ctf4b #ctf4bnext  #seccon https://t.co/TbFeOiXB5I', '2017-10-07 02:23:24');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916476848892125184', 'オリエンテーションが始まりました！！ #ctf4b #ctf4bnext  #seccon https://t.co/xdJH8hsuUk', '2017-10-07 01:34:42');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916465170729246721', '今日は高専の学生さんにサポートしてもらっています！！ #ctf4b #ctf4bnext #seccon https://t.co/Z2s72Sl5H6', '2017-10-07 00:48:18');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916439970037997568', 'なお、重ねてのご案内になりますが会場は東京都立産業技術高等専門学校 品川キャンパスとなります。一部の検索では荒川キャンパスが先にヒットし、会場を間違えてしまう場合もございますので、ご確認の上お気をつけてお越しください！ #ctf4b #ctf4bnext #seccon', '2017-10-06 23:08:10');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/916438876788473856', '今日はSECCON Beginners東京開催、並びにNEXT開催です！東京都立産業技術高等専門学校の品川キャンパスにて開催致します！生憎の天気ですが皆様のご参加お待ちしております！！ #ctf4b #ctf4bnext #seccon', '2017-10-06 23:03:49');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/911493106385563648', '現在CTF中です
一位の方は既に1000点越えとなっているようです
#seccon #ctf4b https://t.co/rptoJ464b2', '2017-09-23 07:31:05');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/911479248518389761', 'リーダーによる CTF オリエンテーションです
#seccon #ctf4b https://t.co/dyJQIAsXl7', '2017-09-23 06:36:01');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/911460771703398400', 'Binaryの講義が始まりました
#seccon #ctf4b https://t.co/Seb8c4HhfT', '2017-09-23 05:22:36');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/911441318651379712', 'Forensics の講義が始まりました！
#seccon #ctf4b https://t.co/t3HVE3akfG', '2017-09-23 04:05:18');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/911412603309629441', 'SECCON Beginners 仙台始まりました！ただいまWebの講義中です #ctf4b #seccon https://t.co/BjiJswME5F', '2017-09-23 02:11:12');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/910450185649643520', 'SECCON Beginners NEXTがSECCON Beginners 東京と同時開催されます！初心者から脱したい方、Beginners以外のCTFでは太刀打ちできない方は奮ってご応募ください！
https://t.co/HWm2zpP7Zg', '2017-09-20 10:26:54');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/910450050748186624', 'SECCON Beginners 東京を10/7（土）に開催します！是非皆様ご参加ください！
https://t.co/H6Jy7i5K8f', '2017-09-20 10:26:22');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/903159078129614849', 'SECCON Beginners 仙台を9/23（土）に開催致します。是非皆様ご参加ください！ #seccon #ctf4b https://t.co/wpqGpziKDA', '2017-08-31 07:34:38');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/901361325686759425', 'CTFの答え合わせをしています！ #seccon #ctf4b https://t.co/2qimsaTQfx', '2017-08-26 08:31:01');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/901333919840092161', 'CTF始まりました！ #seccon #ctf4b', '2017-08-26 06:42:07');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/901269976144781312', '現在Webの演習中です！ #seccon #ctf4b https://t.co/ISRiwiC69E', '2017-08-26 02:28:01');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/889734238412578817', 'SECCON Beginners 広島を8/26（土）に開催致します。是非皆様ご参加ください！ #SECCON #ctf4b https://t.co/rVdvTsIQyJ', '2017-07-25 06:29:07');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/881058675388895232', 'CTF演習も終盤です！！ #ctf4b #seccon https://t.co/DYu0VWkCLS', '2017-07-01 07:55:32');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/881031253813739520', 'すごく集中して取り組んでいます……… #ctf4b #seccon https://t.co/ABFgTiepZ5', '2017-07-01 06:06:34');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/881030725641752576', 'Reversingの講義ではアセンブリ言語入門として，現在関数呼び出しについて解説を行っています！ #ctf4b #seccon', '2017-07-01 06:04:28');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/881023968781647873', 'Reversingの講義は15:20までになります！ #ctf4b #seccon', '2017-07-01 05:37:37');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/881023621577097216', 'Reversingの講義始まってます！！ #ctf4b #seccon https://t.co/lDNT7Blxgp', '2017-07-01 05:36:14');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/881007303289217024', 'Forensicsの講義は14:05まで行います！ #ctf4b #seccon', '2017-07-01 04:31:24');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/881004094466400257', 'パケットの詳細を見るツールとしてWireSharkの使い方に注目しています．例えば，DisplayFilterという機能を利用するとプロトコルなどに注目してパケットを見ることが出来ます． #ctf4b #seccon', '2017-07-01 04:18:39');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/881002326307217408', 'Forensicsの講義が始まりました！！ #ctf4b #seccon https://t.co/2LtTicLmH7', '2017-07-01 04:11:37');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/880977720305369088', 'Webの実践演習中です。前までの内容を踏まえて実際にflagを探します。 #ctf4b #seccon https://t.co/7CLUDtDt7T', '2017-07-01 02:33:50');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/880974486710792192', 'Web問題のよくあるパターンとして，不正な入力に対して無防備である場合などがあります． #ctf4b #seccon', '2017-07-01 02:21:00');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/880973831921217537', 'Webの講義が始まりました！！ #ctf4b #seccon https://t.co/4eqSBuCw6g', '2017-07-01 02:18:23');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/880961974649016320', '間もなく開始です！！ #ctf4b #seccon https://t.co/WKspTe05xQ', '2017-07-01 01:31:16');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/880957042176835584', '受付開始しました！！ #ctf4b #seccon https://t.co/nkGwFHf5Bz', '2017-07-01 01:11:40');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/880748397346471940', '準備完了です！明日はお待ちしております！ #ctf4b #seccon https://t.co/skzJT7LZuZ', '2017-06-30 11:22:36');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878528032729776128', 'CTF演習が終了して、問題解説を行っています！ #ctf4b #seccon https://t.co/bfsQSwM0BU', '2017-06-24 08:19:39');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878495892784824320', 'アセンブリ言語を学びCTFに結びつけるには，ファイルの挙動を確認し，逆アセンブル結果から解析し，知らない命令は検索するというのが重要 #ctf4b #seccon', '2017-06-24 06:11:57');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878486235433402368', 'レジスタ，汎用レジスタ，特殊レジスタに関する解説が行われています． #ctf4b #seccon', '2017-06-24 05:33:34');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878484866655895552', 'Reversingの講義は，Linuxでの実行ファイルの動作原理を解明していくことを目的として進めていきます． #ctf4b #seccon', '2017-06-24 05:28:08');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878484560387883008', 'Reversingの講義が始まりました！ #ctf4b #seccon https://t.co/uR9Mt6JVMj', '2017-06-24 05:26:55');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878468434169798656', 'Forensicsの講義ではTCP flagsの解説を行っています #ctf4b #seccon', '2017-06-24 04:22:50');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878467311681675264', 'Forensicsの講義ではネットワークの通信の仕組みから紐解いています #ctf4b #seccon', '2017-06-24 04:18:22');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878466610092163074', 'Forensicsの講義が始まりました！ #ctf4b #seccon https://t.co/insLHPPEwt', '2017-06-24 04:15:35');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878456479337938945', '今回もYAMAHA様の機材提供を受けましてネットワークを構築しております！ #ctf4b #seccon https://t.co/KOYnO41kix', '2017-06-24 03:35:20');
	insert into tweets(url, text, tweeted_at) values ('https://twitter.com/user/status/878451701107220480', 'お昼休憩です！ #ctf4b #seccon', '2017-06-24 03:16:21');
EOSQL

