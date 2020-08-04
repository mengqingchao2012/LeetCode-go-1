# 173. Binary Search Tree Iterator
## Problem

Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling `next()` will return the next smallest number in the BST.

```
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return true
iterator.next();    // return 15
iterator.hasNext(); // return true
iterator.next();    // return 20
iterator.hasNext(); // return false
```

 

**Note:**

- `next()` and `hasNext()` should run in average O(1) time and uses O(*h*) memory, where *h* is the height of the tree.
- You may assume that `next()` call will always be valid, that is, there will be at least a next smallest number in the BST when `next()` is called.

## Solution

- 难点在于要求 next() 和 hasNext() 方法的平均复杂度是 O(1) 和 O(h)，h 是树高

- 常规做法：通过中序遍历，将节点值存入一个数组中，保存该数组后，就可以实现 next 和 hasNext 方法，但是这种方法的空间复杂度是 O(n)，不满足条件

- 由中序遍历：

  ```c
  void inorder(TreeNode root) {
      if (root == null) return;
      inorder(root.left);
      list.add(root.val);
      inorder(root.right);
  }
  ```

  可知：因为需要递归左右子树，因此保存了所有节点，空间复杂度是 O(n)，要想实现空间复杂度为 O(h)，则可以用栈保存左子树，每次弹出栈顶元素，此时该元素就是调用 next 的返回值，然后对该元素，将其右分支的左子树再次加入栈中。hasNext 方法通过判断栈中元素是否为空来实现