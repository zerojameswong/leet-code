// not too bad, another way to do this is to do depth first search such that all calls share the same used slice and only do the copying of the solution at the end
// thus no need to do so much copying
func permute(nums []int) [][]int {
    var perms [][]int

    var sub func([]int, []bool)

    sub = func(perm []int, used []bool) {
        if len(perm) == len(nums) {
            perms = append(perms, perm)
        }
        for i := range used {
            if !used[i] {
                newPerm := make([]int, len(perm), len(perm) + 1)
                copy(newPerm, perm)
                newPerm = append(newPerm, nums[i])

                newUsed := make([]bool, len(nums))
                copy(newUsed, used)
                newUsed[i] = true

                sub(newPerm, newUsed)
            }
        }
    }

    sub([]int{}, make([]bool, len(nums)))

    return perms
}
