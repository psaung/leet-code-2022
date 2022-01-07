/**
 * @param {string} s
 * @return {string[][]}
 */

function isPalindrome(str, l, r) {
  while (l < r) {
    if (str[l] != str[r]) {
      return false;
    }
    l++;
    r--;
  }
  return true;
}

function range(start, end) {
  var len = end - start + 1;
  var a = new Array(len);

  for (let i = 0; i < len; i++) a[i] = start + i;
  return a;
}

var partition = function (str) {
  let res = [],
    part = [];

  const len = str.length;

  function dfs(i) {
    if (i >= str.length) {
      res.push([...part]);
      return;
    }
    range(i, len).forEach((j) => {
      if (isPalindrome(str, i, j)) {
        part.push(str.slice(i, j + 1));
        dfs(j + 1);
        part.pop();
      }
    });
  }
  dfs(0);
  return res;
};

partition("aab");
