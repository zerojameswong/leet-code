func removeElement(nums []int, val int) int {
    notVal := 0

    for i := range nums {
        if nums[i] != val {
            if nums[notVal] != nums[i] {
                nums[notVal] = nums[i]
            }
            notVal++
        } 
    }

    return notVal
}