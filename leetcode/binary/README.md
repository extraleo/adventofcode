# 二分
> 关键不在于区间内的数具有什么性质, 而是外部的元素具有什么性质
> 理解二分，请牢记区间的定义！区间内的数（下标）都是还未确定与 target 的大小关系的，有的是 < target，有的是 ≥ target；区间外的数（下标）都是确定与 target 的大小关系的。
> 对于本题（递增数组），区间左侧外面的都是 < target，区间右侧外面的都是 ≥ target。从这个定义可以知道，找到了 ≥ target 的数之后，要把这个数（下标）放在区间外面，而不是区间里面！
> 所以对于闭区间写法，当 nums【mid】 >= target 时，要把 mid 放在区间外面，代码就自然是 right = mid - 1 了
> 循环不定式
```go
// 就记住一个点, 区间不能为空
// 对于开区间 (-1, n) 循环的条件是 left +1 < right  return right
// 对于闭区间 [0, n-1] 循环的条件是 left <= right return left
```
> 二分一个是外面的循环, 一个是里面的判断, 外面的循环就是看自己怎么定义区间, 开区间还是闭区间
> 判断就是看对题目的理解了
> **最小的最大，最大的最小，基本都是二分**


- 34 find-first-and-last-position-of-element-in-sorted-array 开区间
- 2529. 正整数和负整数的最大计数 https://leetcode.cn/problems/maximum -count-of-positive-integer-and-negative-integer/
  - 这道题可以看作是找第一个 >=0 的数和 第一个大于等于1的数
- 2300. 咒语和药水的成功对数 https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/
  - > 排序后二分, 因为只是要求数量,没要求 index, 所以先排序 https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/solutions/1595712/by-endlesscheng-1kbp/
  - > 不过在题解中, 灵神提到了**上取整** 转换**下取整**的公式, 这样就避免了每次都去做乘法
  > a/b 上取整  == ((a-1)/b) 下取整 + 1
- [TBD] 2563. 统计公平数对的数目 https://leetcode.cn/problems/count-the-number-of-fair-pairs/
  - 可以看作
- [TBD] 275. H 指数 II
  - 题目没看懂, 什么玩意
- [TBD] 875. 爱吃香蕉的珂珂 https://leetcode.cn/problems/koko-eating-bananas/
  - 转换成向上取整的平均数, 一定是每一个数的向上取整, 而不是 sum 之后再向上取整
- [TBD] 2187. 完成旅途的最少时间 
  - 感觉自己是个智障, 在二分上, 二分的区间不是参数数组的长度, 可以自己定义一个. 这个和 875 是一样的
- [TBD] 2861. 最大合金数 https://leetcode.cn/problems/maximum-number-of-alloys/
  - 看不懂
- 2439. 最小化数组中的最大值 
  - 看分类讨论来做这一题
- [TBD] 2517. 礼盒的最大甜蜜度 https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/
  - 最小值的最大值, 就是用二分, 但是这个题的解法看不懂
- [TBD] 162 f ind-peak-element