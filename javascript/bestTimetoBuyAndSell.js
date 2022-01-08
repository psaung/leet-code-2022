const maxProfit = function (prices) {
  let result = 0,
    stock = prices[0];
  const len = prices.length;

  for (let i = 1; i < len; i++) {
    if (stock < prices[i] && prices[i] > (prices[i + 1] || 0)) {
      result += prices[i] - stock;
      stock = prices[i + 1];
    }
    if (prices[i] < stock) {
      stock = prices[i];
    }
  }
  return result;
};

maxProfit([1, 2, 3, 4, 5]);
