require("@nomicfoundation/hardhat-toolbox");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  gasReporter: {
    enabled: process.env.REPORT_GAS === "true",
    currency: "USD",
    showTimeSpent: true,
  },
};
