function checkExceed(number, sign) {
  const pow = Math.pow(2, 31);
  number = number * sign;
  if (number > pow - 1 || number < -1 * pow) {
    return 0;
  }
  return number;
}

function reverse(x) {
  let result = 0;
  let str = x + "";
  let len = str.length;
  let sign = 1;
  if (str[0] === "-") {
    str = str.substr(1, len);
    sign = -1;
  }
  str = str.split("").reverse().join("");
  len = str.length;
  for (let i = 0; i < len; i++) {
    result = result * 10 + parseInt(str[i]);
  }

  return checkExceed(result, sign);
}
