# beats 80%, O(n) time complexity but does 2 passes, O(n) space complexity rather than O(1)
# gather peaks and obtain 
class Solution:
    def longestMountain(self, arr: List[int]) -> int:
        if len(arr) < 3:
            return 0

        max_length = 0

        direction = 0
        if arr[1] > arr[0]:
            direction = 1
        elif arr[0] > arr[1]:
            direction = -1
        peaks = []

        for i in range(2, len(arr)):
            if direction == 0:
                if arr[i] > arr[i - 1]:
                    direction = 1
                elif arr[i - 1] > arr[i]:
                    direction = -1

            elif direction == 1:
                if arr[i] < arr[i - 1]:
                    peaks.append(i - 1)
                    direction = -1
                elif arr[i] == arr[i - 1]:
                    direction = 0

            elif direction == -1:
                if arr[i] > arr[i - 1]:
                    direction = 1
                elif arr[i] == arr[i - 1]:
                    direction == 0

        for peak_idx in peaks:
            l = peak_idx - 1
            l_prev = arr[peak_idx]
            l_length = 0

            r = peak_idx + 1
            r_prev = arr[peak_idx]
            r_length = 0

            while l >= 0 and arr[l] < l_prev:
                l_length += 1

                l_prev = arr[l]
                l -= 1
            while r < len(arr) and arr[r] < r_prev:
                r_length += 1

                r_prev = arr[r]
                r += 1

            mount_length = 1 + l_length + r_length

            if mount_length > max_length:
                max_length = mount_length

        return max_length
