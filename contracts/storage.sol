pragma solidity 0.5.9;

contract StorageContract {
    
    uint val;
    
    function setVal(uint newval) public {
        val = newval;
    }
    
    function getVal() public view returns (uint) {
        return val;
    }
    
    function submitHeaderBlock(
    bytes calldata vote,
    bytes calldata sigs,
    bytes calldata txData)
    external
  {
  }
      
    function s(uint256 d, address a, bytes32 b, bytes memory b1) public {
        
    }
    
}