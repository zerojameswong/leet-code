// essentially splitting the matrix into rounds like a running track and work from out to in
// for each round we take one less than the length of a side and perform a cycle of swaps with its corresponding entries
func rotate(matrix [][]int)  {
    n := len(matrix)
    nRings := n / 2

    for i := 0; i < nRings; i++ {
        ringSize := n - 2 * i - 1
        for j := 0; j < ringSize; j++ {
            temp := matrix[i][i + j]
            matrix[i][i + j] = matrix[i + ringSize - j][i]
            matrix[i + ringSize - j][i] = matrix[i + ringSize][i + ringSize - j]
            matrix[i + ringSize][i + ringSize - j] = matrix[i + j][i + ringSize]
            matrix[i + j][i + ringSize] = temp
        }
    }
}
