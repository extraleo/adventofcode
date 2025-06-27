# 常用数据结构
https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/

## 前缀和

## 差分 https://leetcode.cn/problems/car-pooling/solutions/2550264/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/
```
对于数组 a，定义其差分数组（difference array）为

d[i]={ 
a[0];  i=0
a[i]−a[i−1];  i≥1
​
性质 1：从左到右累加 d 中的元素，可以得到数组 a。
性质 2：如下两个操作是等价的。
把 a 的子数组 a[i],a[i+1],…,a[j] 都加上 x。
把 d[i] 增加 x，把 d[j+1] 减少 x。
```
- 1094 拼车
只需要记住在第 i 站上车, j 站下车, 中间是不会有变化的, 所以可以做累加.
就是差分. i 站上, j 站下, 那么就是 [i ,j-1] 这段都要加上一个数 == i + n, j-n

