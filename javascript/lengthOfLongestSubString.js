var lengthOfLongestSubstring = function (s) {
  const len = s.length;
  // Base condition
  if (!s) {
    return 0;
  }

  let start = 0;
  let end = 0;
  let maxLength = 0;
  const hashMap = {};
  while (end < len) {
    if (!hashMap[s[end]]) {
      hashMap[s[end]] = true;
      end++;
      maxLength = Math.max(maxLength, Object.keys(hashMap).length);
    } else {
      delete hashMap[s[start]];
      start++;
    }
  }
  return maxLength;
};
