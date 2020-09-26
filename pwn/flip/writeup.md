# Flip - Writeup

## 方針
1. exit の GOT を \_start に向け，繰り返し反転が行えるようにする
2. setbuf の GOT を puts に書き換える & stderr をずらして libc リーク
	- \_\_stack\_chk\_fail の GOT と PLT を適宜利用
3. setbuf の GOT を system に書き換える & stderr を "/bin/sh" に向けてシェル奪取

