// manually doing ksum where k=4
// can also be more selective where to end loop for k=2
func fourSum(nums []int, target int) [][]int {
    solns := [][]int{}

    sort.Ints(nums)

    for i := 0; i < len(nums) - 3; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        for j := i + 1; j < len(nums) - 2; j++ {
            if j > i + 1 && nums[j] == nums[j - 1] {
                continue
            }
            left := j + 1
            right := len(nums) - 1
            for right - left > 0 {
                sum := nums[i] + nums[j] + nums[left] + nums[right]
                diff := target - sum
                var moveLeft, moveRight bool
                if diff == 0 {
                    solns = append(solns, []int{nums[i], nums[j], nums[left], nums[right]})
                    moveLeft = true
                    moveRight = true
                } else if diff > 0 {
                    moveLeft = true

                } else {
                    moveRight = true
                }
                if moveLeft {
                    ref := nums[left]
                    for left < len(nums) - 1 && ref == nums[left] {
                        left++
                    }
                }
                if moveRight {
                    ref := nums[right]
                    for right > j + 1 && ref == nums[right] {
                        right--
                    }
                }
            }
        }
    }

    return solns
}
