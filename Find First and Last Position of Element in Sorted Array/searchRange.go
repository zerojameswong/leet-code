// noice
// the main thing i paid attention to was using nums[from] == target and else if mid < to && target >= nums[mid + 1] && target <= nums[to]
// to ensure that I got the first appearance of the target, then from there i can just linearly search
// another way to do this would have been to use 2 pointers from left and right and just keep going which is also O(n)
func searchRange(nums []int, target int) []int {
    found := []int{-1, -1}

    if len(nums) == 0 {
        return found
    }

    firstOccurence := -1
    var binarySearch func(int, int)
    binarySearch = func(from int, to int) {
        fmt.Println(from, to)
        if nums[from] == target {
            firstOccurence = from
            found[0] = from
            found[1] = from
            return
        }
        mid := from + (to - from) / 2
        if target >= nums[from] && target <= nums[mid] {
            binarySearch(from, mid)
        } else if mid < to && target >= nums[mid + 1] && target <= nums[to] {
            binarySearch(mid + 1, to)
        }
    }
    binarySearch(0, len(nums) - 1)

    if firstOccurence > -1 {
        i := firstOccurence + 1
        lastOccurence := firstOccurence
        for ; i < len(nums); i++ {
            if nums[i] != target {
                break
            }
            lastOccurence++
        }
        found[1] = lastOccurence
    }

    return found
}
