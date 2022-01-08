var removeDuplicates = function (nums) {
  let count = 0,
    prev;
  for (let i = 0; i < nums.length; i++) {
    if (prev !== nums[i]) {
      nums[count] = nums[i];
      prev = nums[i];
      count++;
    }
  }
  return count;
};
