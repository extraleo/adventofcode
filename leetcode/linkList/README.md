# 链表

## 反转链表
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
- 2 两数相加 I https://leetcode.cn/problems/add-two-numbers/solutions/2327008/dong-hua-jian-ji-xie-fa-cong-di-gui-dao-oe0di/
  - 尾插法, 要哨兵节点
- 445. 两数相加 II https://leetcode.cn/problems/add-two-numbers-ii/
  - 反转 -> 2 两数相加 -> 反转
- 2816. 翻倍以链表形式表示的数字 https://leetcode.cn/problems/double-a-number-represented-as-a-linked-list/
  - 翻倍就是两个一样的数相加

## 快慢指针

- 876. 链表的中间结点 https://leetcode.cn/problems/middle-of-the-linked-list/solution/mei-xiang-ming-bai-yi-ge-shi-pin-jiang-t-wzwm/
  - for fast != nil && fast.Next != nil  return slow
141. 环形链表 https://leetcode.cn/problems/linked-list-cycle/solution/mei-xiang-ming-bai-yi-ge-shi-pin-jiang-t-c4sw/
  - for fast != nil && fast.Next != nil    if fast == slow  return true
142. 环形链表 II https://leetcode.cn/problems/linked-list-cycle-ii/solution/mei-xiang-ming-bai-yi-ge-shi-pin-jiang-t-nvsq/
143. 重排链表 https://leetcode.cn/problems/reorder-list/solution/mei-xiang-ming-bai-yi-ge-shi-pin-jiang-t-u66q/

课后作业：
234. 回文链表 https://leetcode.cn/problems/palindrome-linked-list/
  - 中间节点 + 反转链表
2130. 链表最大孪生和 https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/
  - 中间节点 + 反转

## 删除节点/链表去重  
237. 删除链表中的节点 https://leetcode.cn/problems/delete-node-in-a-linked-list/solution/ru-he-shan-chu-jie-dian-liu-fen-zhong-ga-x3kn/
19. 删除链表的倒数第 N 个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/solution/ru-he-shan-chu-jie-dian-liu-fen-zhong-ga-xpfs/
83. 删除排序链表中的重复元素 https://leetcode.cn/problems/remove-duplicates-from-sorted-list/solution/ru-he-qu-zhong-yi-ge-shi-pin-jiang-tou-p-98g7/
82. 删除排序链表中的重复元素 II https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/solution/ru-he-qu-zhong-yi-ge-shi-pin-jiang-tou-p-2ddn/

课后作业：
203. 移除链表元素 https://leetcode.cn/problems/remove-linked-list-elements/
3217. 从链表中移除在数组中存在的节点 https://leetcode.cn/problems/delete-nodes-from-linked-list-present-in-array/
2487. 从链表中移除节点 https://leetcode.cn/problems/remove-nodes-from-linked-list/
1669. 合并两个链表 https://leetcode.cn/problems/merge-in-between-linked-lists/  