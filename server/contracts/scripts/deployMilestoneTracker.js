async function main() {
  console.log("Deploying MilestoneTracker contract...");

  const MilestoneTracker = await ethers.getContractFactory("MilestoneTracker");
  const milestoneTracker = await MilestoneTracker.deploy();

  await milestoneTracker.waitForDeployment();

  const address = await milestoneTracker.getAddress();
  console.log("MilestoneTracker deployed to:", address);
  console.log("\nAdd this to your .env file:");
  console.log(`MILESTONE_TRACKER_ADDRESS=${address}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
