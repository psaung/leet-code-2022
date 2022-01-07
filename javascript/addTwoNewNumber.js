/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */

const l2 = {
  val: 5,
  next: {
    val: 6,
    next: {
      val: 4,
      next: null,
    },
  },
};

const l1 = {
  val: 2,
  next: {
    val: 4,
    next: {
      val: 3,
      next: null,
    },
  },
};

function reverse(list) {
  var curr = list;
  var prev;

  while (curr) {
    let next = curr.next;
    curr.next = prev || null;
    prev = curr;
    curr = next;
  }

  return prev;
}

function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}

var addTwoNumbers = function (l1, l2) {
  let head = new ListNode(0);
  let result = head;

  let carry = 0;
  let sum = 0;

  while (l1 !== null || l2 !== null || carry !== 0) {
    sum = 0;
    if (l1 !== null) {
      sum += l1.val;
      l1 = l1.next;
    }
    if (l2 !== null) {
      sum += l2.val;
      l2 = l2.next;
    }
    sum += carry;
    carry = parseInt(sum / 10);

    result.next = new ListNode(sum % 10);
    result = result.next;
  }

  return head.next;
};

addTwoNumbers(l2, l1);
