package main

import (
	"fmt"
)

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 10339 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	visted := make(map[byte]bool)
	maxLen := 0
	n := len(s)
	rk := 0
	for i := 0; i < n; i++ { // 左指针
		if i != 0 {
			//左指针往右移动
			delete(visted, s[i-1])
		}
		for rk < n && !visted[s[rk]] {
			fmt.Println(string(s[i]), string(s[rk]), i, rk+1)
			visted[s[rk]] = true
			rk++
		}
		fmt.Println("===")
		maxLen = max(maxLen, rk-i)
	}

	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	maxLen := lengthOfLongestSubstring("pwwkew")
//	fmt.Println(maxLen)
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
