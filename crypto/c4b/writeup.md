# C4B - Writeup

## 問題のポイント
1. Ethereumのクライアントを適切に使えるかどうか
2. `private` 修飾子が着いた変数の読み出し方がわかるかどうか
3. `blockhash(block.number - 1)` を計算できるかどうか
の主に3つある。

2, 3はそれぞれが単純な問題になっていて、それを組み合わせたものが問題になっている。
難易度調整の都合でMedium1つだと少々難しすぎるためEasy2つにしてある。
このあたりは今後の修正により変わるかもしれない。

## 方針
### `private` 変数の読み取り
まず、privateといってもEthereumでは全てのデータがチェーン上に載っていて公開されているので読み出せる。

Contractのデータ領域はStorageと呼ばれていて`getStrageAt(addr, num)`を使うと読み取れる。
`web3.eth.getStrageAt(addr, 0)` を実行すると `player` の値が返ってくる。
次に、`web3.eth.getStrageAt(addr, 1)` を実行するとランダムっぽい値8バイト分とその後ろに0がついたデータが返ってくる。
入手したい `password` は `bytes8` 型なのでこの8バイトをパスワードとして使えばいい。

### `blockhash(block.number - 1)` の計算
1つ前のブロックのブロックハッシュを入れるだけに見えるが、
1つ前のブロックハッシュを手に入れるのは難しい。

想定解は、問題を解くコントラクトを自作してその中の関数から動的に1つ前のブロックハッシュを手に入れる。
試してないけど行けるかもしれない解法はマイニング完了のイベントを受け取って即checkを叩くようなイベントハンドラを書く方法。
どちらにしてもEthereum上でのプログラミングが必要になるので目的としてはOK。

## コード
↓こんな感じのを[Remix](https://remix.ethereum.org/)につっこんでポチッとすると解ける。
```
pragma solidity >= 0.5.0 < 0.7.0;

contract C4B {

  address public player;
  bytes8 password;
  bool public success;

  constructor(address _player, bytes8 _password) public {
    player = _player;
    password = _password;
    success = false;
  }

  function check(bytes8 _password, uint256 pin) public returns(bool) {
    uint256 hash = uint256(blockhash(block.number - 1));
    if (hash%1000000 == pin) {
      if (keccak256(abi.encode(password)) == keccak256(abi.encode(_password))) {
        success = true;
      }
    }
    return success;
  }
}


contract SolveC4B {
    address addr;
    constructor(address _addr) public {
        addr = _addr;
    }
    function solve() public returns(bool) {
        C4B chal = C4B(addr);
        uint256 hash = uint256(blockhash(block.number - 1));
        bytes8 b = bytes8(0xd97150ee12d47357);

        bool result = chal.check(b, hash % 1000000);
        return result;
    }
}
```
