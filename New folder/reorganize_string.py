# beats 96%
# idea is that we whittle down the most frequently occuring letters by interlacing them
# with each other
# when you have (1, 1, 1, 1, 1) of different letters there will always be a way to ensure
# that there are no repeats
# the use of temp is so that if the most recent character in the final string is 
# also the most frequently appearing character then it is ignored to avoid repeating it
class Solution:
    def reorganizeString(self, s: str) -> str:
        counts = {}
        for char in s:
            if char not in counts:
                counts[char] = 1
            else:
                counts[char] += 1
        
        heap = []
        for k, v in counts.items():
            heap.append((-v, k))
        heapq.heapify(heap)

        res = []
        prev = ''
        temp = None
        while heap:
            popped = heapq.heappop(heap)
            count, char = popped

            if char == prev:
                temp = popped
                continue

            res.append(char)
            prev = char
            if temp is not None:
                heapq.heappush(heap, temp)
                temp = None
            if count < -1:
                heapq.heappush(heap, (count + 1, char))

        if len(res) < len(s):
            return ''
        else:
            return ''.join(res)
