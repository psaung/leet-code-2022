function check32bit(number, sign = 1) {
  const pow = Math.pow(2, 31);
  if (number >= pow) {
    return sign < 0 ? pow : pow - 1;
  }
  return number;
}

function checkNull(val) {
  return isNaN(val) || val === null || val === undefined;
}

var myAtoi = function (s) {
  const len = s.length;
  let result = null,
    occ = 0,
    sign = 1;

  for (let i = 0; i < len; i++) {
    if (result === null && s[i] === "-") {
      if (result === null) {
        sign = -1;
      }
      result = 0;
      continue;
    }

    if (result === null && s[i] === "+") {
      result = 0;
      continue;
    }

    const inttt = s[i] === " " ? null : parseInt(s[i]);

    if (result !== null && checkNull(inttt)) {
      break;
    }
    if (inttt !== null) {
      result = result * 10 + parseInt(s[i]);
    }
  }

  if (!result) {
    return 0;
  }

  return check32bit(result, sign) * sign;
};
