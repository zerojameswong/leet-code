// brute force O(n^2) solution
// also go does not have a builtin min function for ints...
func maxArea(height []int) int {
    maxArea := 0

    for i := 0; i < len(height) - 1; i++ {
        for j := i + 1; j < len(height); j++ {
            minHeight := height[i]
            if height[i] > height[j] {
                minHeight = height[j]
            }
            breadth := j - i
            area := minHeight * breadth
            if area > maxArea {
                maxArea = area
            }
        }
    }

    return maxArea
}

// o(n) solution whose optimality is not obvious or intuitive
// the idea is that we have 2 pointers at the ends of the slice and the lower of the two boundaries is the bottleneck and we relieve that by moving the lower pointer inwards
func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    
    maxArea := 0

    for right > left {
        lower := "left"
        minHeight := height[left]
        if height[right] < height[left] {
            lower = "right"
            minHeight = height[right]
        }
        breadth := right - left
        area := minHeight * breadth
        if area > maxArea {
            maxArea = area
        }

        if lower == "left" {
            left++
        } else {
            right--
        }
    }

    return maxArea
}

