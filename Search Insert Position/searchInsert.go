// spawn every candidate slice of 2 elements and return upper, lower if they match exactly or upper if the number is in between
func searchInsert(nums []int, target int) int {
    position := -1

    if target < nums[0] {
        return 0
    }
    if target > nums[len(nums) - 1] {
        return len(nums)
    }
    
    var binarySearch func(int, int)
    binarySearch = func(from, to int) {
        fmt.Println(from, to)
        if nums[from] == target {
            position = from
            return
        }
        if nums[to] == target {
            position = to
            return
        }
        if to - from == 1 && target > nums[from] && target < nums[to] {
            position = to
            return
        }

        mid := from + (to - from) / 2
        if target >= nums[from] && target <= nums[mid] {
            binarySearch(from, mid)
        }
        if target >= nums[mid] && target <= nums[to] {
            binarySearch(mid, to)
        }
    }

    binarySearch(0, len(nums) - 1)

    return position
}
