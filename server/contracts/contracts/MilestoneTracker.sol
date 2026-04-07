// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

/**
 * @title MilestoneTracker
 * @dev Tracks donation milestones WITHOUT holding or transferring ETH
 * Used purely for milestone calculation and event emission
 * Actual fund disbursement happens off-chain via backend
 */
contract MilestoneTracker {
    struct Cause {
        uint256 goal;           // Funding goal in smallest currency unit
        uint256 collected;      // Total collected amount
        uint8 milestonesPaid;   // Number of milestones reached (0-4)
        bool exists;            // Whether cause is registered
    }

    mapping(bytes16 => Cause) public causes;

    event CauseRegistered(bytes16 indexed causeId, uint256 goal);
    event DonationRecorded(bytes16 indexed causeId, uint256 amount, uint256 newCollected);
    event MilestoneReached(
        bytes16 indexed causeId,
        uint8 milestone,           // 1=25%, 2=50%, 3=75%, 4=100%
        uint256 amountToDiburse    // Amount to disburse for this milestone
    );

    /**
     * @dev Register a cause with its funding goal and initial collected amount
     * Can only be called once per causeId (idempotent after first call)
     */
    function registerOrUpdateCause(bytes16 causeId, uint256 goal, uint256 initialCollected) external {
        require(goal > 0, "Goal must be greater than zero");
        
        if (!causes[causeId].exists) {
            causes[causeId] = Cause({
                goal: goal,
                collected: initialCollected,
                milestonesPaid: 0,
                exists: true
            });
            emit CauseRegistered(causeId, goal);
            
            // Check if initial amount already crosses milestones
            if (initialCollected > 0) {
                _processMilestones(causeId, causes[causeId]);
            }
        } else {
            // Update goal if cause already exists (allows goal updates)
            causes[causeId].goal = goal;
        }
    }

    /**
     * @dev Record a donation amount and check for milestone progress
     * Does NOT receive or transfer ETH - just tracks the amount
     */
    function recordDonation(bytes16 causeId, uint256 amount) external {
        require(causes[causeId].exists, "Cause not registered");
        require(amount > 0, "Amount must be greater than zero");

        Cause storage c = causes[causeId];
        c.collected += amount;

        emit DonationRecorded(causeId, amount, c.collected);

        // Check and process milestones
        _processMilestones(causeId, c);
    }

    /**
     * @dev Internal function to check and emit milestone events
     */
    function _processMilestones(bytes16 causeId, Cause storage c) internal {
        // Process up to 4 milestones (25%, 50%, 75%, 100%)
        while (c.milestonesPaid < 4) {
            uint8 nextMilestone = c.milestonesPaid + 1;
            uint256 nextThreshold = (uint256(nextMilestone) * c.goal) / 4;

            // If we haven't reached the next milestone, stop
            if (c.collected < nextThreshold) {
                break;
            }

            // Calculate amount for this milestone (always 25% of goal)
            uint256 milestoneAmount = c.goal / 4;

            // Mark milestone as paid
            c.milestonesPaid = nextMilestone;

            // Emit event for backend to process
            emit MilestoneReached(causeId, nextMilestone, milestoneAmount);
        }
    }

    /**
     * @dev Get cause information
     */
    function getCause(bytes16 causeId) external view returns (
        uint256 goal,
        uint256 collected,
        uint8 milestonesPaid,
        bool exists
    ) {
        Cause storage c = causes[causeId];
        return (c.goal, c.collected, c.milestonesPaid, c.exists);
    }

    /**
     * @dev Check if a specific milestone has been reached
     */
    function isMilestoneReached(bytes16 causeId, uint8 milestone) external view returns (bool) {
        require(milestone >= 1 && milestone <= 4, "Milestone must be 1-4");
        return causes[causeId].milestonesPaid >= milestone;
    }
}
