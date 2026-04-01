const hre = require("hardhat");

async function main() {
  const f = await hre.ethers.getContractFactory("CauseMilestoneEscrow");
  const escrow = await f.deploy();
  await escrow.waitForDeployment();
  console.log("CauseMilestoneEscrow:", await escrow.getAddress());
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
