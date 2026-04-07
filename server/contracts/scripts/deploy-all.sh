#!/bin/bash

# Deploy all contracts script
# Usage: ./deploy-all.sh [network]
# Example: ./deploy-all.sh localhost
# Example: ./deploy-all.sh sepolia

NETWORK=${1:-localhost}

echo "========================================="
echo "Deploying all contracts to: $NETWORK"
echo "========================================="

# Deploy DonationLedger
echo -e "\nDeploying DonationLedger..."
npx hardhat run scripts/deployDonationLedger.js --network $NETWORK

# Deploy MilestoneTracker
echo -e "\nDeploying MilestoneTracker..."
npx hardhat run scripts/deployMilestoneTracker.js --network $NETWORK

echo -e "\n========================================="
echo "All contracts deployed!"
