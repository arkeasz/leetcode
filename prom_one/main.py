class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        num_index = []
        for i in range(len(nums)):
            for j in range(len(nums)):
                if i != j and nums[i] + nums[j] == target and i < j:
                    num_index = [i, j]
        return num_index

nums = [1, 4, 5, 2]
target = 5

res = Solution.twoSum(Solution, nums=nums, target=target)

print(res)
