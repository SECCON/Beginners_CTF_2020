pragma solidity >= 0.5.0 < 0.7.0;

import "./C4B.sol";

contract ChallengeManager {

  event ChallengeDeployed(address indexed player, address challenge);

  function deploy(address player) public {
    bytes memory base = abi.encodePacked(blockhash(block.number - 1));
    bytes8 password = bytes8(keccak256(base));
    C4B chal = new C4B(player, password);
    emit ChallengeDeployed(player, address(chal));
  }
}
