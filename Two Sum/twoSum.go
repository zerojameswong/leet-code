// O(n^2)
func twoSum(nums []int, target int) []int {
    var indexPair []int
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++  {
            if nums[i] + nums[j] == target {
                indexPair = []int{i, j}
                break
            }
        }
    }
    return indexPair
}

// O(n)
func twoSum(nums []int, target int) []int {
    var indexPair []int
    var complements = make(map[int]int)
    
    for idx, num := range(nums) {
        complementValue := target - num
        complements[complementValue] = idx
    }

    for idx, num := range(nums) {
        complementIdx, exists := complements[num]
        if exists && idx != complementIdx {
            if complementIdx < idx {
                indexPair = []int{complementIdx, idx}
            } else {
                indexPair = []int{idx, complementIdx}
            }
        }
    }

    return indexPair
}

// O(nlog(n))
Sort then use two pointers at ends of the array, push right pointer to the left if sum of two pointers > target and push left pointer to the right if sum of two pointers < target
