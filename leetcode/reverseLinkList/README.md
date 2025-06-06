# 反转链表
```
nxt := cur.Next
cur.Next = pre
pre = cur
cur = nxt
```
> 在反转后, pre 是反转前链表的最后一个节点, 而 cur 是最后一个节点的 next 节点


- 206. 反转链表 https://leetcode.cn/problems/reverse-linked-list/solution/you-xie-cuo-liao-yi-ge-shi-pin-jiang-tou-o5zy/
- 92. 反转链表 II https://leetcode.cn/problems/reverse-linked-list-ii/solution/you-xie-cuo-liao-yi-ge-shi-pin-jiang-tou-teqq/
- 25. K 个一组翻转链表 [美团社招] https://leetcode.cn/problems/reverse-nodes-in-k-group/solution/you-xie-cuo-liao-yi-ge-shi-pin-jiang-tou-plfs/
  - 首先有个大循环, 有几次反转
  - 在每次反转之后, 需要更新 p0 节点

课后作业：
- 24. 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/
  - `k = 2` 的 25题
445. 两数相加 II https://leetcode.cn/problems/add-two-numbers-ii/
2816. 翻倍以链表形式表示的数字 https://leetcode.cn/problems/double-a-number-represented-as-a-linked-list/