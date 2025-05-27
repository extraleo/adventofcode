# 二分
> 关键不在于区间内的数具有什么性质, 而是外部的元素具有什么性质
> 循环不定式
```go
// 就记住一个点, 区间不能为空
// 对于开区间 (-1, n) 循环的条件是 left +1 < right  return right
// 对于闭区间 [0, n-1] 循环的条件是 left <= right return left
```


- 34 find-first-and-last-position-of-element-in-sorted-array 开区间
- 2529. 正整数和负整数的最大计数 https://leetcode.cn/problems/maximum -count-of-positive-integer-and-negative-integer/
  - 这道题可以看作是找第一个 >=0 的数和 第一个大于等于1的数
- [TBD] 2300. 咒语和药水的成功对数 https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/
  - > 排序后二分, 因为只是要求数量,没要求 index, 所以先排序 https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/solutions/1595712/by-endlesscheng-1kbp/
  - > 不过在题解中, 灵神提到了**上取整** 转换**下取整**的公式, 这样就避免了每次都去做乘法
  > a/b 上取整  == ((a-1)/b) 下取整 + 1
- 2563. 统计公平数对的数目 https://leetcode.cn/problems/count-the-number-of-fair-pairs/
  - 可以看作
- [TBD] 275. H 指数 II https://leetcode.cn/problems/h-index-ii/
- [TBD] 875. 爱吃香蕉的珂珂 https://leetcode.cn/problems/koko-eating-bananas/
- [TBD] 2187. 完成旅途的最少时间 https://leetcode.cn/problems/minimum-time-to-complete-trips/
- [TBD] 2861. 最大合金数 https://leetcode.cn/problems/maximum-number-of-alloys/
- [TBD] 2439. 最小化数组中的最大值 https://leetcode.cn/problems/minimize-maximum-of-array/
- [TBD] 2517. 礼盒的最大甜蜜度 https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/
- [TBD] 162 f ind-peak-element