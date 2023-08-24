// O(n^3) brute force
func threeSumClosest(nums []int, target int) int {
    var (
        closest int
        smallestDifference int = math.MaxInt32
    )

    for i := 0; i < len(nums) - 2; i++ {
        for j := i + 1; j < len(nums) - 1; j++ {
            for k := j + 1; k < len(nums); k++ {
                sum := nums[i] + nums[j] + nums[k]
                diff := sum - target
                if diff < 0 {
                    diff = -diff
                }
                if diff < smallestDifference {
                    closest = sum
                    smallestDifference = diff
                }
            }
        }
    }

    return closest
}

// using the three sum method but still kind of ugly
func threeSumClosest(nums []int, target int) int {
    var (
        closestSum int
        smallestDifference int = math.MaxInt32
    )

    sort.Ints(nums)

    out:
    for i := 0; i < len(nums) - 2; i++ {
        left := i + 1
        right := len(nums) - 1

        for right - left >= 1 {
            sum := nums[i] + nums[left] + nums[right]
            if sum == target {
                closestSum = sum
                smallestDifference = 0
                break out
            }

            diff := sum - target

            if diff < 0 {
                diff = -diff
                left++
            } else {
                right--
            }

            if diff < smallestDifference {
                closestSum = sum
                smallestDifference = diff
            }
        }
    }

    return closestSum
}
