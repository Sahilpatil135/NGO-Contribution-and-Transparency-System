async function main() {
  const DonationLedger = await ethers.getContractFactory("DonationLedger");
  const donationLedger = await DonationLedger.deploy();

  await donationLedger.waitForDeployment();

  console.log("DonationLedger Contract address:", await donationLedger.getAddress());
}

main();
