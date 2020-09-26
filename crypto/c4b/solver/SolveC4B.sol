pragma solidity >= 0.5.0 < 0.7.0;

contract ChallengeManager {
  event ChallengeDeployed(address indexed player, address challenge);

  function deploy(address player) public {
    bytes memory base = abi.encodePacked(blockhash(block.number - 1));
    bytes8 password = bytes8(keccak256(base));
    C4B chal = new C4B(player, password);
    emit ChallengeDeployed(player, address(chal));
  }
}

pragma solidity >= 0.5.0 < 0.7.0;

contract C4B {

  address public player;
  bytes8 password;
  bool public success;

  event CheckPassed(address indexed player);

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
    emit CheckPassed(player);
    return success;
  }
}


contract SolveC4B {
  C4B chal;
  bytes8 password;
  constructor(address addr) public {
    chal = C4B(addr);
  }

  function solve(bytes8 password) public returns(bool) {
    uint256 hash = uint256(blockhash(block.number - 1));
    bool result = chal.check(password, hash % 1000000);
    return result;
  }
}
