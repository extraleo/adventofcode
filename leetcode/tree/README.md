# Tree

## 递归
104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/solution/kan-wan-zhe-ge-shi-pin-rang-ni-dui-di-gu-44uz/

课后作业：
111. 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/
112. 路径总和 https://leetcode.cn/problems/path-sum/
129. 求根节点到叶节点数字之和 https://leetcode.cn/problems/sum-root-to-leaf-numbers/
1448. 统计二叉树中好节点的数目 https://leetcode.cn/problems/count-good-nodes-in-binary-tree/
[hard - TBD] 987. 二叉树的垂序遍历 https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree/


## 相同, 平衡, 右视图

100. 相同的树 https://leetcode.cn/problems/same-tree/solution/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-empk/
101. 对称二叉树 https://leetcode.cn/problems/symmetric-tree/solution/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-6dq5/
110. 平衡二叉树 https://leetcode.cn/problems/balanced-binary-tree/solution/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-c3wj/
199. 二叉树的右视图 https://leetcode.cn/problems/binary-tree-right-side-view/solution/ru-he-ling-huo-yun-yong-di-gui-lai-kan-s-r1nc/
> 第一次使用go的闭包

课后作业[all TBD]：
965. 单值二叉树 https://leetcode.cn/problems/univalued-binary-tree/
951. 翻转等价二叉树 https://leetcode.cn/problems/flip-equivalent-binary-trees/
226. 翻转二叉树 https://leetcode.cn/problems/invert-binary-tree/
617. 合并二叉树 https://leetcode.cn/problems/merge-two-binary-trees/
2331. 计算布尔二叉树的值 https://leetcode.cn/problems/evaluate-boolean-binary-tree/
508. 出现次数最多的子树元素和 https://leetcode.cn/problems/most-frequent-subtree-sum/
1026. 节点与其祖先之间的最大差值 https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/
1372. 二叉树中的最长交错路径 https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/
1080. 根到叶路径上的不足节点 https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths/


## 二叉搜索树
// 中序遍历修改值而不是返回值
98. 验证二叉搜索树 https://leetcode.cn/problems/validate-binary-search-tree/solution/qian-xu-zhong-xu-hou-xu-san-chong-fang-f-yxvh/

课后作业：
700. 二叉搜索树中的搜索 https://leetcode.cn/problems/search-in-a-binary-search-tree/
938. 二叉搜索树的范围和 https://leetcode.cn/problems/range-sum-of-bst/
530. 二叉搜索树的最小绝对差 https://leetcode.cn/problems/minimum-absolute-difference-in-bst/
  - 中序遍历
2476. 二叉搜索树最近节点查询 https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/
  - 中序遍历+二分查找
501. 二叉搜索树中的众数 https://leetcode.cn/problems/find-mode-in-binary-search-tree/
230. 二叉搜索树中第 K 小的元素 https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
1373. 二叉搜索子树的最大键值和 https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree/
105. 从前序与中序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
106. 从中序与后序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
889. 根据前序和后序遍历构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
1110. 删点成林 https://leetcode.cn/problems/delete-nodes-and-return-forest/




## 最近公共祖先
236. 二叉树的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/solutions/2023872/fen-lei-tao-lun-luan-ru-ma-yi-ge-shi-pin-2r95/
235. 二叉搜索树的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/solutions/2023873/zui-jin-gong-gong-zu-xian-yi-ge-shi-pin-8h2zc/

课后作业：【TBD】
1123. 最深叶节点的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/
2096. 从二叉树一个节点到另一个节点每一步的方向 https://leetcode.cn/problems/step-by-step-directions-from-a-binary-tree-node-to-another/


## BFS
102. 二叉树的层序遍历 https://leetcode.cn/problems/binary-tree-level-order-traversal/solution/bfs-wei-shi-yao-yao-yong-dui-lie-yi-ge-s-xlpz/
两个数组, 一个 cur 一个ans 
有 cur 必有 nxt
ans 里面 append vals, vals 是cur 里面node 的 val
103. 二叉树的锯齿形层序遍历 https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/solution/bfs-wei-shi-yao-yao-yong-dui-lie-yi-ge-s-xlv3/
  偶数层的时候， append ans 时，反转一下再 append
513. 找树左下角的值 https://leetcode.cn/problems/find-bottom-left-tree-value/solution/bfs-wei-shi-yao-yao-yong-dui-lie-yi-ge-s-f34y/
  层序遍历, ans[-1][0]
课后作业：
107. 二叉树的层序遍历 II https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
层序遍历, return的时候 ans[::-1]
104. 二叉树的最大深度 https://leetcode.cn/problems/maximum-depth-of-binary-tree/
111. 二叉树的最小深度 https://leetcode.cn/problems/minimum-depth-of-binary-tree/
2583. 二叉树中的第 K 大层和 https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/
199. 二叉树的右视图 https://leetcode.cn/problems/binary-tree-right-side-view/
116. 填充每个节点的下一个右侧节点指针 https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
117. 填充每个节点的下一个右侧节点指针 II https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
1302. 层数最深叶子节点的和 https://leetcode.cn/problems/deepest-leaves-sum/
1609. 奇偶树 https://leetcode.cn/problems/even-odd-tree/
2415. 反转二叉树的奇数层 https://leetcode.cn/problems/reverse-odd-levels-of-binary-tree/
2641. 二叉树的堂兄弟节点 II https://leetcode.cn/problems/cousins-in-binary-tree-ii/