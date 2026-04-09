/**
 * Utility functions for handling different goal types based on aid_type
 */

// Aid types that should show non-monetary goals
const AID_TYPE_CONFIG = {
  "Volunteering": {
    type: "volunteers",
    goalLabel: "Volunteers Needed",
    collectedLabel: "Volunteers Registered",
    unit: "",
    singularUnit: "",
    showCurrency: false,
  },
  "Blood Donations": {
    type: "blood_units",
    goalLabel: "Units Needed",
    collectedLabel: "Units Collected",
    unit: "units",
    singularUnit: "unit",
    showCurrency: false,
  },
  // Default for all other aid types (monetary)
  _default: {
    type: "monetary",
    goalLabel: "Goal Amount",
    collectedLabel: "Amount Raised",
    unit: "₹",
    singularUnit: "₹",
    showCurrency: true,
  },
};

/**
 * Get goal configuration based on aid type name
 * @param {string} aidTypeName - The name of the aid type (e.g., "Volunteering", "Blood Donations")
 * @returns {Object} Configuration object with type, labels, and units
 */
export const getGoalConfig = (aidTypeName) => {
  return AID_TYPE_CONFIG[aidTypeName] || AID_TYPE_CONFIG._default;
};

/**
 * Format the goal amount based on aid type
 * @param {number} amount - The goal amount
 * @param {string} aidTypeName - The name of the aid type
 * @returns {string} Formatted goal string
 */
export const formatGoal = (amount, aidTypeName) => {
  const config = getGoalConfig(aidTypeName);
  const value = parseFloat(amount) || 0;

  if (config.showCurrency) {
    return `₹${value.toLocaleString()}`;
  }

  // For non-monetary goals
  const count = Math.floor(value);
  const unit = count === 1 ? config.singularUnit : config.unit;
  return `${count} ${unit}`;
};

/**
 * Format the collected amount based on aid type
 * @param {number} amount - The collected amount
 * @param {string} aidTypeName - The name of the aid type
 * @returns {string} Formatted collected string
 */
export const formatCollected = (amount, aidTypeName) => {
  return formatGoal(amount, aidTypeName);
};

/**
 * Check if an aid type uses monetary goals
 * @param {string} aidTypeName - The name of the aid type
 * @returns {boolean} True if monetary, false otherwise
 */
export const isMonetary = (aidTypeName) => {
  const config = getGoalConfig(aidTypeName);
  return config.showCurrency;
};

/**
 * Get the appropriate label for the goal field
 * @param {string} aidTypeName - The name of the aid type
 * @returns {string} Label text
 */
export const getGoalLabel = (aidTypeName) => {
  const config = getGoalConfig(aidTypeName);
  return config.goalLabel;
};

/**
 * Get the appropriate label for the collected field
 * @param {string} aidTypeName - The name of the aid type
 * @returns {string} Label text
 */
export const getCollectedLabel = (aidTypeName) => {
  const config = getGoalConfig(aidTypeName);
  return config.collectedLabel;
};

/**
 * Get placeholder text for goal input field
 * @param {string} aidTypeName - The name of the aid type
 * @returns {string} Placeholder text
 */
export const getGoalPlaceholder = (aidTypeName) => {
  const config = getGoalConfig(aidTypeName);

  if (config.showCurrency) {
    return "e.g., 100000";
  }

  if (config.type === "volunteers") {
    return "e.g., 50";
  }

  if (config.type === "blood_units") {
    return "e.g., 100";
  }

  return "e.g., 100";
};
