// File: frontend/src/utils/helpers.js
export const formatCurrency = (amount) => {
  // Ensure amount is a number before formatting
  const price = parseFloat(amount);
  if (isNaN(price)) {
    return 'N/A';
  }
  
  // Formats the number as Kenyan Shillings. You can change 'KES' to 'USD' or other currencies.
  return price.toLocaleString('en-US', {
    style: 'currency',
    currency: 'KES', 
  });
};