package main

import (
	Deque "container/list"
) // 使用 Go 的 container/list 包作为双端队列

//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 2886 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	// 创建一个双端队列
	deq := Deque.New()
	maxValues := []int{}

	for i := 0; i < len(nums); i++ {
		// 移除不在窗口内的索引
		if deq.Len() > 0 && deq.Front().Value.(int) < i-k+1 {
			deq.Remove(deq.Front())
		}

		// 移除小于当前元素的索引
		for deq.Len() > 0 && nums[deq.Back().Value.(int)] < nums[i] {
			deq.Remove(deq.Back())
		}

		// 添加当前元素的索引
		deq.PushBack(i)

		// 记录当前窗口的最大值
		if i >= k-1 {
			maxValues = append(maxValues, nums[deq.Front().Value.(int)])
		}
	}
	return maxValues
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	slidingWindow := maxSlidingWindow([]int{
//		1, 3, -1, -3, 5, 3, 6, 7,
//	}, 3)
//	fmt.Println(slidingWindow)
//}
