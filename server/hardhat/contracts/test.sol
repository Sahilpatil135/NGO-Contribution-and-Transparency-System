// SPDX-License-Identifier: Unidentified
pragma solidity ^0.8.28;

contract Counter {
	uint public x = 0;
	
	function inc() public {
	    x += 1;
	}

	function getX() public view returns (uint) {
		return x;
	}
}
