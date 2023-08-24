// i did it i beat 100%!!! the main tricky thing is the boundaries
// one thing is when pivot/ mid is last we cant index pivot + 1 / mid + 1
// and we cant spawn 2 sub problems with overlapping boundaries
func search(nums []int, target int) int {
    if len(nums) == 0 {
        return -1
    }

    pivot := 0
    var binarySearchPivot func(int, int)
    binarySearchPivot = func(from int, to int) {
        if to - from == 1 && nums[from] > nums[to] {
            pivot = from
            return
        }
        mid := from + (to - from) / 2
        if nums[from] > nums[mid] {
            binarySearchPivot(from, mid)
        }
        if nums[mid] > nums[to] {
            binarySearchPivot(mid, to)
        }
    }
    binarySearchPivot(0, len(nums) - 1)

    found := -1
    var binarySearchTarget func(int, int)
    binarySearchTarget = func(from int, to int) {
        if nums[from] == target {
            found = from
            return
        }
        mid := from + (to - from) / 2
        if target >= nums[from] && target <= nums[mid] {
            binarySearchTarget(from, mid)
        }
        if mid < to && target >= nums[mid + 1] && target <= nums[to] {
            binarySearchTarget(mid + 1, to)
        }

    }
    if target >= nums[0] && target <= nums[pivot] {
        binarySearchTarget(0, pivot)
    }
    if pivot < len(nums) - 1 && target >= nums[pivot + 1] && target <= nums[len(nums) - 1] {
        binarySearchTarget(pivot + 1, len(nums) - 1)
    }

    return found
}
