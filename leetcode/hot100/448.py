class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        ans = []
        tmp = [-1] * (len(nums)+1)
        print(len(tmp))
        for x in nums:
            tmp[x] = 1
        for i in range(1,len(nums)+1):
            if tmp[i] == -1:
                ans.append(i)
        return ans
