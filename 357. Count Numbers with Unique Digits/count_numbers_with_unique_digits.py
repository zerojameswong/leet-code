# brute force method exceeds time limit for n = 8
class Solution:
    def countNumbersWithUniqueDigits(self, n: int) -> int:
        number = 1
        for i in range(n):
            number *= 10

        unique_count = 0
        for i in range(number):
            to_str = str(i)
            count = {}
            for char in to_str:
                if char not in count:
                    count[char] = 1
                else:
                    break
            else:
                unique_count += 1

        return unique_count
        
# my idea to cache previous results for 10^(i-1) but exceeds time limit for n = 8
class Solution:
    def countNumbersWithUniqueDigits(self, n: int) -> int:
        unique_count = 1

        repeats = {}
        new_repeats = {}
        for i in range(1, n + 1):
            if i == 1:
                unique_count = 10
                continue
            if i == 2:
                for j in range(10, 100):
                    if j % 11 == 0:
                        repeats[j] = True
                    else:
                        unique_count += 1
                continue

            for j in range(int(10 ** (i - 1)), 10 ** i):
                rem = j % 10 ** (i - 1)
                if rem in repeats:
                    new_repeats[j] = True
                else:
                    to_str = str(j)
                    count = {}
                    for char in to_str:
                        if char not in count:
                            count[char] = 1
                        else:
                            new_repeats[j] = True
                            break
                    else:
                        unique_count += 1

                repeats = new_repeats
                new_repeats = {}

        return unique_count

# using some dp and combinatorics
# pay attention to calculating the increment only after recursing so that each recursive call returns
# before multiplying by the increment
# nonlocal incremement required since int is not mutable, another way to do it is to store dicts in memo
class Solution:
    def countNumbersWithUniqueDigits(self, n: int) -> int:
        memo = {}
        memo[0] = 1
        memo[1] = 10

        increment = 9 * 9
        memo[2] = memo[1] + increment

        def count_unique(n):
            unique_count = 0

            if n in memo:
                return memo[n]

            nonlocal increment
            prev = count_unique(n - 1)
            increment = (11 - n) * increment
            unique_count = increment + count_unique(n - 1)
            memo[n] = unique_count

            return unique_count

        return count_unique(n)
