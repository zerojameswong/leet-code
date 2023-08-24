// my brutish method, the biggest obstacle here is in making sure that the solutions are unique
func combinationSum(candidates []int, target int) [][]int {
    var toConvert [][]int
    var visited [][]int
    var backtrack func(int, []int)
    var isVisited func([]int) bool

    backtrack = func(target int, dist []int) {
        if target == 0 {
            toConvert = append(toConvert, dist)
        }

        for i := range candidates {
            if target - candidates[i] >= 0 {
                check := make([]int, len(candidates))
                copy(check, dist)
                check[i]++
                if !isVisited(check) {
                    visited = append(visited, check)
                    backtrack(target - candidates[i], check)
                }
            }
        }
    }

    isVisited = func(check []int) bool {
        for i := range visited {
            j := 0
            for ; j < len(candidates); j++ {
                if check[j] != visited[i][j] {
                    break
                }
            }
            if j == len(candidates) {
                return true
            }
        }
        return false
    }

    backtrack(target, make([]int, len(candidates)))

    var combis [][]int
    for i := range toConvert {
        var combi []int
        for j := range toConvert[i] {
            for k := 0; k < toConvert[i][j]; k++ {
            combi = append(combi, candidates[j])
            }
        }
        combis = append(combis, combi)
    }

    return combis
}

// a more elegant (copied) solution that uses a different decision tree that restricts the 
func combinationSum(candidates []int, target int) [][]int {
    var combis [][]int

    var dfs func(int, []int, int)

    dfs = func(currentSum int, soln []int, candidateIdx int) {
        if currentSum == target {
            combis = append(combis, soln)
            return
        }
        if currentSum > target || candidateIdx >= len(candidates) {
            return
        }

        newSoln := make([]int, len(soln), len(soln) + 1)
        copy(newSoln, soln)
        newSoln = append(newSoln, candidates[candidateIdx])
        dfs(currentSum + candidates[candidateIdx], newSoln, candidateIdx)
        candidateIdx++
        dfs(currentSum, soln, candidateIdx)
    }

    dfs(0, []int{}, 0)

    return combis
}
