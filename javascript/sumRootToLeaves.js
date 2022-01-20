var sumRootToLeaf = function (root) {
  let result = 0,
    a = "";

  function sumOrder(node, a) {
    a += node.val;
    if (!node.left && !node.right) {
      result += parseInt(a, 2);
    } else {
      node.left && preOrder(node.left, a);
      node.right && preOrder(node.right, a);
    }
  }

  sumOrder(root, "");

  return result;
};
