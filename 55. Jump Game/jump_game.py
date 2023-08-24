class Solution:
    def canJump(self, nums: List[int]) -> bool:
        can_reach_last = [False] * len(nums)
        can_reach_last[-1] = True

        for i in range(len(nums) - 2, -1, -1):
            for j in range(1, nums[i] + 1):
                if i + j <= len(nums) - 1 and can_reach_last[i + j]:
                    can_reach_last[i] = True
                    break

        return can_reach_last[0]

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        first_can_reach = len(nums) - 1

        for i in range(len(nums) - 2, -1, -1):
            if i + nums[i] >= first_can_reach:
                first_can_reach = i

        return first_can_reach == 0
