// we sort and skip over duplicates once a sub problem is spawned with a duplicate number
func permuteUnique(nums []int) [][]int {
    var perms [][]int

    var dfs func([]int, []bool)

    dfs = func(perm []int, used []bool) {
        if len(perm) == len(nums) {
            perms = append(perms, perm)
            return
        }
        
        i := 0
        for i < len(nums) {
            if !used[i] {
                newPerm := make([]int, len(perm), len(nums))
                copy(newPerm, perm)
                newPerm = append(newPerm, nums[i])

                newUsed := make([]bool, len(nums))
                copy(newUsed, used)
                newUsed[i] = true

                dfs(newPerm, newUsed)

                i++
                for i < len(nums) && nums[i] == nums[i - 1] {
                    i++
                }
            } else {
                i++
            }
        }
    }

    sort.Ints(nums)
    dfs([]int{}, make([]bool, len(nums)))

    return perms
}

// same just phrased differently
func permuteUnique(nums []int) [][]int {
    var perms [][]int

    var dfs func([]int, []bool)

    dfs = func(perm []int, used []bool) {
        if len(perm) == len(nums) {
            perms = append(perms, perm)
            return
        }
        
        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue
            }
            if i > 0 && nums[i - 1] == nums[i] && !used[i - 1] {
                continue
            }

            newPerm := make([]int, len(perm), len(nums))
            copy(newPerm, perm)
            newPerm = append(newPerm, nums[i])

            newUsed := make([]bool, len(nums))
            copy(newUsed, used)
            newUsed[i] = true

            dfs(newPerm, newUsed)
        }
    }

    sort.Ints(nums)
    dfs([]int{}, make([]bool, len(nums)))

    return perms
}
