package main

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 5226 👎 0

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func trap(height []int) int {
	// 双指针

	ans := 0
	left, right := 0, len(height)-1

	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}

	}

	return ans
}

//leetcode submit region begin(Prohibit modification and deletion)
func trap2(height []int) int {

	// DP
	length := len(height)

	left := make([]int, length)
	right := make([]int, length)

	left[0] = height[0]
	right[length-1] = height[length-1]

	for i := 1; i < length; i++ {
		left[i] = max(left[i-1], height[i])

	}

	for i := length - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	ans := 0

	for i := 0; i < length; i++ {
		ans += min(left[i], right[i]) - height[i]
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//
//	ans := trap([]int{
//		0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
//	})
//	fmt.Println(ans)
//}
