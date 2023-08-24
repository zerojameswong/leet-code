// i think this is really readable and O(n) but leetcode thinks its only better than 30% or sth
func removeDuplicates(nums []int) int {
    sorted := 0
    
    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        nums[sorted] = nums[i]
        sorted++
    }

    return sorted
}
