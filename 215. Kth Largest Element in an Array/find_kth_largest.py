# beats 88%, using heap, max or min depending on how close k is to either end
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        is_min = k > len(nums) / 2

        heap = []

        for num in nums:
            if is_min:
                heapq.heappush(heap, num)
            else:
                heapq.heappush(heap, -num)

        if is_min:
            while len(heap) > k:
                heapq.heappop(heap)
            return heap[0]
        else:
            for _ in range(k):
                res = heapq.heappop(heap)
            return -res
