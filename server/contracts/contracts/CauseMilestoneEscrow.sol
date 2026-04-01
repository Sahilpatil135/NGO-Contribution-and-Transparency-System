// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title CauseMilestoneEscrow
 * @notice Holds native ETH per cause and automatically sends 25% of the declared goal
 *         to the organization's wallet each time cumulative deposits cross 25%, 50%, 75%, and 100%
 *         of that goal. Amounts are tracked in wei; `goal` is total funding target in wei.
 */
contract CauseMilestoneEscrow {
    address public owner;

    struct Cause {
        uint256 goal;
        uint256 collected;
        uint256 released;
        uint8 milestonesPaid;
        address payable beneficiary;
        bool exists;
    }

    mapping(bytes16 => Cause) private _causes;

    uint256 private _locked;

    event CauseRegistered(bytes16 indexed causeId, uint256 goal, address beneficiary);
    event DonationReceived(bytes16 indexed causeId, address indexed donor, uint256 amount, uint256 newCollected);
    event MilestoneReleased(bytes16 indexed causeId, uint8 milestone, uint256 amount, address indexed beneficiary);

    error NotOwner();
    error CauseAlreadyExists();
    error CauseNotFound();
    error InvalidGoal();
    error InvalidBeneficiary();
    error ZeroDonation();
    error TransferFailed();

    modifier onlyOwner() {
        if (msg.sender != owner) revert NotOwner();
        _;
    }

    modifier nonReentrant() {
        if (_locked == 1) revert();
        _locked = 1;
        _;
        _locked = 0;
    }

    constructor() {
        owner = msg.sender;
    }

    function transferOwnership(address newOwner) external onlyOwner {
        owner = newOwner;
    }

    /// @notice Register a cause before accepting ETH donations. `goal` is the full target in wei.
    function registerCause(bytes16 causeId, uint256 goal, address payable beneficiary) external onlyOwner {
        if (_causes[causeId].exists) revert CauseAlreadyExists();
        if (goal == 0) revert InvalidGoal();
        if (beneficiary == address(0)) revert InvalidBeneficiary();

        _causes[causeId] = Cause({
            goal: goal,
            collected: 0,
            released: 0,
            milestonesPaid: 0,
            beneficiary: beneficiary,
            exists: true
        });

        emit CauseRegistered(causeId, goal, beneficiary);
    }

    /// @notice Donate native ETH to a cause; milestones are settled in the same transaction.
    function donate(bytes16 causeId) external payable nonReentrant {
        if (msg.value == 0) revert ZeroDonation();
        Cause storage c = _causes[causeId];
        if (!c.exists) revert CauseNotFound();

        c.collected += msg.value;
        emit DonationReceived(causeId, msg.sender, msg.value, c.collected);

        _processMilestones(causeId, c);
    }

    function _processMilestones(bytes16 causeId, Cause storage c) internal {
        while (c.milestonesPaid < 4) {
            uint256 nextThreshold = ((uint256(c.milestonesPaid) + 1) * c.goal) / 4;
            if (c.collected < nextThreshold) {
                break;
            }
            uint256 payout = nextThreshold - c.released;
            if (payout == 0) {
                break;
            }
            c.released += payout;
            c.milestonesPaid += 1;

            (bool ok, ) = c.beneficiary.call{value: payout}("");
            if (!ok) revert TransferFailed();

            emit MilestoneReleased(causeId, c.milestonesPaid, payout, c.beneficiary);
        }
    }

    function getCause(bytes16 causeId)
        external
        view
        returns (
            uint256 goal,
            uint256 collected,
            uint256 released,
            uint256 milestonesPaid,
            address beneficiary,
            bool exists
        )
    {
        Cause storage c = _causes[causeId];
        return (c.goal, c.collected, c.released, c.milestonesPaid, c.beneficiary, c.exists);
    }
}
