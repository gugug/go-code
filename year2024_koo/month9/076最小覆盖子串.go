package main

import (
	"math"
)

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//
//
//
// 注意：
//
//
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
//
//
//
// 示例 1：
//
//
//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
//
// 示例 2：
//
//
//输入：s = "a", t = "a"
//输出："a"
//解释：整个字符串 s 是最小覆盖子串。
//
//
// 示例 3:
//
//
//输入: s = "a", t = "aa"
//输出: ""
//解释: t 中两个字符 'a' 均应包含在 s 的子串中，
//因此没有符合条件的子字符串，返回空字符串。
//
//
//
// 提示：
//
//
// m == s.length
// n == t.length
// 1 <= m, n <= 10⁵
// s 和 t 由英文字母组成
//
//
//
//进阶：你能设计一个在
//o(m+n) 时间内解决此问题的算法吗？
//
// Related Topics 哈希表 字符串 滑动窗口 👍 3005 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {

	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, rigth := 0, 0
	window := make(map[byte]int)

	// valid 表示窗口中满足条件的字符个数
	valid := 0
	// 记录最小覆盖子串的起始索引及长度
	minLen := math.MaxInt32
	start := 0

	for rigth < len(s) {
		// 扩大窗口
		c := s[rigth]
		rigth++

		// 更新窗口中的数据
		if _, exist := need[c]; exist {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断是否需要收缩窗口
		for valid == len(need) {
			// 更新最小字符串
			if rigth-left < minLen {
				start = left
				minLen = rigth - left
			}
			// 缩小窗口
			d := s[left]
			left++
			if _, exist := need[d]; exist {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}

//leetcode submit region end(Prohibit modification and deletion)

//func main() {
//	var s = "ADOBECODEBANC"
//	var t = "ABC"
//
//	window := minWindow(s, t)
//	fmt.Println(window)
//}
