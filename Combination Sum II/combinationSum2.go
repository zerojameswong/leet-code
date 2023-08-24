// only beats 23.5% apparently
func combinationSum2(candidates []int, target int) [][]int {
    var combis [][]int
    var backtracking func(int, int, []int)

    sort.Ints(candidates)

    backtracking = func(currentSum, currentIndex int, combi []int) {
        if currentSum == target {
            combis = append(combis, combi)
            return
        }
        if currentSum > target || currentIndex >= len(candidates) {
            return
        }

        newCombi := make([]int, len(combi), len(combi) + 1)
        copy(newCombi, combi)
        newCombi = append(newCombi, candidates[currentIndex])
        newSum := currentSum + candidates[currentIndex]
        backtracking(newSum, currentIndex + 1, newCombi)
        for currentIndex + 1 < len(candidates) && candidates[currentIndex + 1] == candidates[currentIndex] {
            currentIndex++
        }
        backtracking(currentSum, currentIndex + 1, combi)
    }

    backtracking(0, 0, []int{})

    return combis
}

// same leh
// not using double problem instead using spawn all sub problems with unique numbers
func combinationSum2(candidates []int, target int) [][]int {
    var combis [][]int
    var backtracking func(int, int, []int)

    sort.Ints(candidates)

    backtracking = func(currentSum, currentIndex int, combi []int) {
        if currentSum == target {
            combis = append(combis, combi)
            return
        }
        if currentSum > target || currentIndex >= len(candidates) {
            return
        }

        for i := currentIndex; i < len(candidates); i++ {
            if i > currentIndex && candidates[i] == candidates[i - 1] {
                continue
            }
            newCombi := make([]int, len(combi), len(combi) + 1)
            copy(newCombi, combi)
            newCombi = append(newCombi, candidates[i])
            newSum := currentSum + candidates[i]
            backtracking(newSum, i + 1, newCombi)
        }
    }
    
    backtracking(0, 0, []int{})

    return combis
}
