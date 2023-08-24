// O(n^3) and horribly ugly
func threeSum(nums []int) [][]int {
    allTriplets := [][]int{}

    for i := 0; i < len(nums) - 2; i++ {
        for j := i + 1; j < len(nums) - 1; j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] + nums[j] + nums[k] == 0 {
                    solution := []int{nums[i], nums[j], nums[k]} 
                    sort.Ints(solution)

                    found := false
                    for i := range(allTriplets) {
                        j := 0
                        for ; j < len(solution); j++ {
                            if allTriplets[i][j] != solution[j] {
                                break
                            }
                        }
                        if j == 3 {
                            found = true
                            break
                        }
                    }

                    if !found {
                        allTriplets = append(allTriplets, solution)
                    }
                }
            }
        }
    }

    return allTriplets
}

/* 
 beautiful online solution O(n^2)
 by sorting the slice we incur O(nlogn) time for sorting
 but we test against the previous first element and skip if it has already been processed
 this prevents us from repeating solutions

 for example -3 -3 -3 0 1 2 6 (sorted already) has (-3, -3, 6) and (-3, 1, 2) as solutions
 if we just let this slide, the same solutions will be found if a solution exists for that first element
 look at -3 | -3 -3 0 1 2 6 and see that for the next loop it is the exact same problem -3 -3 | -3 0 1 2 6
 the next loop is -3 -3 -3 | 0 1 2 6 
 essentially for any solution i, j, k if there exists some l = i, then l, j, k is a valid but repeated solution

 once the first element is fixed, we are left with the O(n) 2 pointer solution for twosum
 once we find a solution dont break and keep going as there is still the possible for more solutions in between
 for example [-4, -3, -2, 6, 7] has 2 solutions (-4, -3, -7) and (-4, -2, -6)
 
 i added a check for nums[i] <= 0 since if the first element under consideration is positive
 since in a sorted array there is no chance that summation with any subsequent entries can bring the sum any closer to zero
*/
func threeSum(nums []int) [][]int {
    sort.Ints(nums)

    sols := [][]int{}

    for i := 0; i < len(nums) - 2 && nums[i] <= 0; i++ {
        if i >= 1 && nums[i] == nums[i - 1] {
            continue
        }

        left := i + 1
        right := len(nums) - 1

        complement := -nums[i]

        for right - left >= 1 {
            if nums[left] + nums[right] == complement {
                sols = append(sols, []int{nums[i], nums[left], nums[right]})
                left++
                right--
            } else if nums[left] + nums[right] < complement {
                left++
            } else {
                right--
            }
        }
    }

    return sols
}
