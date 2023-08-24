# works but exceeds memory limit on large inputs
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        max_value = 0
        max_value_length = 0
        
        bitwise_vals = [[0 for j in range(len(nums))] for i in range(len(nums))]

        for i in range(len(nums)):
            bitwise_vals[i][i] = nums[i]
            if nums[i] > max_value:
                max_value = nums[i]
                max_value_length = 1

        for length in range(2, len(nums) + 1):
            for i in range(0, len(nums) - length + 1):
                j = i + length - 1
                val = bitwise_vals[i][j - 1] & bitwise_vals[j][j]
                bitwise_vals[i][j] = val
                if val >= max_value:
                    max_value = val
                    max_value_length = length

        return max_value_length
        
# works but exceeds memory limit
# important to recurse on the smaller problem first for val1 and val2 
# if not you would need to go backwards on the max length which there is no logic for in this code
# you can ignore sub-problem order by taking the max of lengths if value is the same
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        max_val = 0
        max_val_len = 0
        memo = {}

        def find(l, r):
            nonlocal max_val, max_val_len

            if (l, r) in memo:
                return memo[l, r]

            val = 0
            if l == r:
                memo[l, r] = nums[l]
                val = nums[l]
            else:
                val1 = find(l, l) & find(l + 1, r)
                val2 = find(r, r) & find(l, r - 1)
                val = val1

            length = r - l + 1
            if val >= max_val:
                max_val = val
                if length > max_val_len:
                    max_val_len = length

            memo[l, r] = val

            return val

        find(0, len(nums) - 1)

        return max_val_len

# beats 95%, makes use of the fact that the max bitwise AND is always a number and itself so
# the problem reduces to finding the longest subarray of the maximum value
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        max_val = 0
        max_val_len = 0
        current_len = 0

        for num in nums:
            if num < max_val:
                current_len = 0
            elif num > max_val:
                max_val = num
                max_val_len = 1
                current_len = 1
            elif num == max_val:
                current_len += 1
                if current_len > max_val_len:
                    max_val_len = current_len

        return max_val_len
