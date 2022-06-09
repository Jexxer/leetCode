class Solution:
    def two_sum(self, nums: list[int], target: int) -> list[int]:
        num_dict = {}
        for i, num in enumerate(nums):
            if target - num not in num_dict:
                num_dict[num] = i
            else:
                return [num_dict.get(target - nums[i]), i]

instance = Solution()
print(instance.two_sum([2, 7, 11, 15], 9))