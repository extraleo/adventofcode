class NumArray:

    def __init__(self, nums: List[int]):
        s = [0] * (len(nums)+1)
        for i, x in enumerate(nums):
            s[i+1] = s[i] + x
        self.s = s

    def sumRange(self, left: int, right: int) -> int:
        return self.s[right+1] - self.s[left]


# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)