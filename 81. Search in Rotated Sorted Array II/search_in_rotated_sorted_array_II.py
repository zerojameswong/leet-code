class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        prev = nums[0]

        for i in nums:
            if i == target:
                return True
            if i < prev:
                if target > nums[0]:
                    return False
            prev = i

        return False
