// too slow for some problmes
func jump(nums []int) int {
    var shortest = len(nums)
    var backtrack func(int, int)

    backtrack = func(idx, jumps int) {
        if idx >= len(nums) {
            return
        }
        if idx == len(nums) - 1 {
            if jumps < shortest {
                shortest = jumps
            }
            return
        }
        for i := 1; i <= nums[idx]; i++ {
            backtrack(idx + i, jumps + 1)
        }
    }

    backtrack(0, 0)

    return shortest
}

// this also takes too long
func jump(nums []int) int {
    var shortest = len(nums)
    var backtrack func(int, int)

    backtrack = func(idx, jumps int) {
        if idx == 0 {
            if jumps < shortest {
                shortest = jumps
            }
            return
        }
        for i := idx - 1; i >= 0; i-- {
            if idx - i <= nums[i] {
                backtrack(i, jumps + 1)
            }
        }
    }

    backtrack(len(nums) - 1, 0)

    return shortest
}

// this completes in time, O(n)
// all about expanding the reachable region
func jump(nums []int) int {
    var (
        shortest int
        left int 
        right int
    )

    for right < len(nums) - 1 {
        furthest := right
        for i := left; i <= right; i++ {
            if i + nums[i] > furthest {
                furthest = i + nums[i]
            }
        }
        left = right + 1
        right = furthest
        shortest++
    }


    return shortest
}
