/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function (num) {
  const keys = [1, 5, 10, 50, 100, 500, 1000];
  const values = ["I", "V", "X", "L", "C", "D", "M"];

  let result = "",
    idx = keys.length - 1,
    t = num;

  while (t != 0) {
    if (idx > 0 && keys[idx] - keys[idx - 1] == keys[idx - 1]) {
      nxtIdx = idx - 2;
    } else {
      nxtIdx = idx - 1;
    }

    // console.log(nxtIdx >= 0 && t < keys[idx] && t >= keys[idx] - keys[nxtIdx], nxtIdx, keys[idx], t, keys[idx], idx)
    const shouldSkip =
      nxtIdx >= 0 ? t < keys[idx] && t >= keys[idx] - keys[nxtIdx] : false;

    // console.log(shouldSkip, nxtIdx, keys[idx], t)
    if (t - keys[idx] >= 0 || shouldSkip) {
      // 1 && 94 < 100 && 100 < 94 > 100 - 10
      if (nxtIdx >= 0 && t < keys[idx] && t >= keys[idx] - keys[nxtIdx]) {
        t = t - keys[idx] + keys[nxtIdx];
        result += values[nxtIdx] + values[idx];
      } else {
        t = t - keys[idx];
        result += values[idx];
      }
    }
    const innerShould =
      nxtIdx >= 0 ? !(t < keys[idx] && t >= keys[idx] - keys[nxtIdx]) : true;
    if (t < keys[idx] && innerShould) {
      idx--;
    }
  }
  return result;
};
