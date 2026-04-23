const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Metrics gas tests", function () {
  it("records donation in DonationLedger", async function () {
    const DonationLedger = await ethers.getContractFactory("DonationLedger");
    const ledger = await DonationLedger.deploy();
    await ledger.waitForDeployment();

    const donationId = ethers.hexlify(ethers.randomBytes(16));
    const causeId = ethers.hexlify(ethers.randomBytes(16));
    const donorId = ethers.hexlify(ethers.randomBytes(16));

    const tx = await ledger.recordDonation(donationId, causeId, donorId, 1000n, "PAY-1");
    await tx.wait();

    const donation = await ledger.getDonation(donationId);
    expect(donation.amount).to.equal(1000n);
  });

  it("processes milestone thresholds in MilestoneTracker", async function () {
    const MilestoneTracker = await ethers.getContractFactory("MilestoneTracker");
    const tracker = await MilestoneTracker.deploy();
    await tracker.waitForDeployment();

    const causeId = ethers.hexlify(ethers.randomBytes(16));
    const goal = 10000n;
    await (await tracker.registerOrUpdateCause(causeId, goal, 0n)).wait();

    await (await tracker.recordDonation(causeId, 2500n)).wait();
    await (await tracker.recordDonation(causeId, 2500n)).wait();

    const cause = await tracker.getCause(causeId);
    expect(cause.milestonesPaid).to.equal(2n);
    expect(cause.collected).to.equal(5000n);
  });
});
