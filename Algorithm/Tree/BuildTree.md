## BuildTree



### 迭代实现
```java

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        // 判断越界问题
        if(preorder == null || preorder.length == 0){
            return null;
        }
        // 初始化栈和头节点 ，前序遍历的第一个节点是二叉树的头节
        // 将头节点添加到队列中
        TreeNode root = new TreeNode(preorder[0]);
        Deque<TreeNode> stack = new LinkedList<TreeNode>();
        stack.push(root);
        // 指向中序遍历节点的指针
        int inorderIndex = 0;
        // 从1开始，因为头节点已经存入到stack中
        for(int i = 1; i < preorder.length; i++){
            // 将当前节点的值存入到临时变量中
            int preorderVal = preorder[i];
            TreeNode node = stack.peek();
            // 将当前节点与中序遍历数组指针指向的数据比较。
            // 如果当前节点的值不等于当前指针所指向的值，将当前前序数组中的值存入stack中，并将其添加为左子节点
            if(node.val != inorder[inorderIndex]){
                node.left = new TreeNode(preorderVal);
                stack.push(node.left);
            }else{
                // 如果当前栈顶的元素等于指针所指向的数据，将指针加1和stack.pop
                while(!stack.isEmpty() && stack.peek().val == inorder[inorderIndex]){
                    node = stack.pop();
                    inorderIndex++;
                }
                // 直到最后一个元素，则该元素是最后一个弹出元素的右子节点
                node.right = new TreeNode(preorderVal);
                // 将当前元素push到栈中
                stack.push(node.right);
            }
        }
        return root;
    }
}

```