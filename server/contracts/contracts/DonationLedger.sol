// SPDX-License-Identifier: Unidentified
pragma solidity ^0.8.28;

contract DonationLedger {
	struct Donation {
	    bytes16 causeId;
	    bytes16 donorId;
	    uint256 amount;
	    uint256 timestamp;
	    string paymentRef;
	}

	uint256 donationsCount;

	mapping(bytes16 => Donation) public donations;
	mapping(bytes16 => bytes16[]) public donationsByCause;

	function recordDonation(
	    bytes16 donationId,
	    bytes16 causeId,
	    bytes16 donorId,
	    uint256 amount,
	    string memory paymentRef
	) public {
		donations[donationId] = Donation({
			causeId: causeId,
			donorId: donorId,
			amount: amount,
			timestamp: block.timestamp,
			paymentRef: paymentRef
		});

		donationsByCause[causeId].push(donationId);

		donationsCount++;
	}

	function getDonation(bytes16 donationId) public view returns (Donation memory) {
		return donations[donationId];
	}

	function getDonationsByCause(bytes16 causeId) public view returns (bytes16[] memory) {
		return donationsByCause[causeId];
	}
}
