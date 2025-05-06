# 双指针和滑动窗口
## 题单
https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/
 
## 双向双指针
> 用 O(1) 的时间获得 O(N) 的信息
### 两数之和
- 167 two-sum-ii-input-array-is-sorted
- 15 3sum
- 2824 count-pairs-whose-sum-is-less-than-target

## 定长滑动窗口
### 基础
> 核心思想是 +1 和 -1, 因为窗口滑动的过程中, 中间的数是不变的
> 其中滑动窗口开始循环的下标, 既可以从 0 开始, 也可以从k 开始, 从k 开始的话, 就要先计算好一个base值, 在这个值上进行 +1 和 -1 的操作. insert -> update -> delete
> https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/solutions/2809359/tao-lu-jiao-ni-jie-jue-ding-chang-hua-ch-fzfo/

- 1456 maximum-number-of-vowels-in-a-substring-of-given-length medium
- 643 maximum-average-subarray-i simple
- 1343 number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold medium
- 2090 k-radius-subarray-averages medium
- 2841 maximum-sum-of-almost-unique-subarray
- 2461 maximum-sum-of-distinct-subarrays-with-length-k
  - same as 2841, m == k
- [TBD] 1423 maximum-points-you-can-obtain-from-cards
  1. 环形数组的滑动窗口
  2. 计算长为 n−k 的连续子数组和的最小值

### 进阶 - 选做 TBD

## 不定长滑动窗口
只要是连续的子数组, 都可以用滑动窗口
### 基础
- 3 longest-substring-without-repeating-characters
- 3090 maximum-length-substring-with-two-occurrences
  - 和 3 是异曲同工, 也和 2461 2841 很像. 2道定长. 重新写一下3
