# use the fact that we can calculate the penalty once then incrementally calculate it
# for each preceding closing time
# beats 67%
# actually can build from the front as the well so no need for the first O(n) pass
# (using negative penalty for Y and positive for N, the actual penalty doesnt matter,
# just the minimum closing time)
# just that this solution was most intuitive to me to begin with
class Solution:
    def bestClosingTime(self, customers: str) -> int:
        min_penalty = len(customers)
        closing_time = -1

        penalty = 0
        for i in customers:
            if i == 'N':
                penalty += 1
        if penalty < min_penalty:
            min_penalty = penalty
            closing_time = len(customers)
        
        for i in range(len(customers) - 1, -1, -1):
            if customers[i] == 'Y':
                penalty += 1
            else:
                penalty -= 1

            if penalty <= min_penalty:
                min_penalty = penalty
                closing_time = i

        return closing_time
