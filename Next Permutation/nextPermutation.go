// i did it!
// but ugly and only beats 17% on time
func nextPermutation(nums []int)  {
    if len(nums) <= 1 {
        return
    }

    if nums[len(nums) - 1] > nums[len(nums) - 2] {
        nums[len(nums) - 1], nums[len(nums) - 2] = nums[len(nums) - 2], nums[len(nums) - 1]
        return
    }

    for i := len(nums) - 3; i >= 0 ; i-- {
        seenHigher := false
        for j := i + 1; j < len(nums); j++ {
            if nums[j] > nums[i] {
                seenHigher = true
            }
            if seenHigher && (nums[j] <= nums[i] || j == len(nums) - 1) {
                var toSwitch int
                if nums[j] <= nums[i] {
                    toSwitch = j - 1
                } else {
                    toSwitch = j
                }
                nums[toSwitch], nums[i] = nums[i], nums[toSwitch]

                swaps := (len(nums) - 1 - (i + 1) - 1) / 2 + 1
                for k := 0; k < swaps; k++ {
                    nums[i + 1 + k], nums[len(nums) - 1 - k] = nums[len(nums) - 1 - k], nums[i + 1 + k]
                }
                return
            }
        }
    }

    swaps := (len(nums) - 2) / 2 + 1
    for k := 0; k < swaps; k++ {
        nums[k], nums[len(nums) - 1 - k] = nums[len(nums) - 1 - k], nums[k]
    }
}