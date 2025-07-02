# 动态规划

198. 打家劫舍 https://leetcode.cn/problems/house-robber/solution/ru-he-xiang-chu-zhuang-tai-ding-yi-he-zh-1wt1/

课后作业：
70. 爬楼梯 https://leetcode.cn/problems/climbing-stairs/
如果最后一步爬了 1 个台阶，那么我们得先爬到 i−1，要解决的问题缩小成：从 0 爬到 i−1 有多少种不同的方法。
如果最后一步爬了 2 个台阶，那么我们得先爬到 i−2，要解决的问题缩小成：从 0 爬到 i−2 有多少种不同的方法。
由于这两种方法是互相独立的（爬的台阶个数不同），所以根据加法原理，从 0 爬到 i 的方法数等于这两种方法数之和，即

dfs(i)=dfs(i−1)+dfs(i−2)

746 使用最小花费爬楼梯 https://leetcode.cn/problems/min-cost-climbing-stairs/
和 70 一样

213. 打家劫舍 II https://leetcode.cn/problems/house-robber-ii/
740. 删除并获得点数 https://leetcode.cn/problems/delete-and-earn/
2466. 统计构造好字符串的方案数 https://leetcode.cn/problems/count-ways-to-build-good-strings/
377. 组合总和 Ⅳ https://leetcode.cn/problems/combination-sum-iv/
2266. 统计打字方案数 https://leetcode.cn/problems/count-number-of-texts/
64. 最小路径和 https://leetcode.cn/problems/minimum-path-sum/


## 0-1 背包
494. 目标和 https://leetcode.cn/problems/target-sum/solution/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-s1cx/
322. 零钱兑换 https://leetcode.cn/problems/coin-change/solution/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-21m5/

课后作业：
2915. 和为目标值的最长子序列的长度 https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/
416. 分割等和子集 https://leetcode.cn/problems/partition-equal-subset-sum/
2787. 将一个数字表示成幂的和的方案数 https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/
518. 零钱兑换 II https://leetcode.cn/problems/coin-change-ii/
279. 完全平方数 https://leetcode.cn/problems/perfect-squares/