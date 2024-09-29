package main

import (
	"sort"
)

//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//
//
//
// 示例 1:
//
//
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
//
//
// 示例 2:
//
//
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
//
//
//
//
// 提示:
//
//
// 1 <= s.length, p.length <= 3 * 10⁴
// s 和 p 仅包含小写字母
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 1512 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func GetAnagrams(p string) string {
	pbytes := []byte(p)
	sort.Slice(pbytes, func(i, j int) bool {
		return pbytes[i] < pbytes[j]
	})
	sortedPStr := string(pbytes)
	return sortedPStr
}
func findAnagrams2(s string, p string) []int {
	var ans []int
	pAnagrams := GetAnagrams(p)

	n := len(s)

	for index := range s {
		right := index + len(p)
		if right > n {
			break
		}

		subStr := s[index:right]
		subAnagrms := GetAnagrams(subStr)
		if subAnagrms == pAnagrams {
			ans = append(ans, index)
		}
	}

	return ans
}

func findAnagrams(s string, p string) []int {
	var ans []int

	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}

	var sCount, pCount [26]int
	for i := range p {
		sCount[s[i]-'a']++ //记录s中前pLen个字母的词频
		pCount[p[i]-'a']++ //记录要寻找的字符串中每个字母的词频(只用进行一次来确定)
	}

	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] { //i是滑动前的首位
		sCount[ch-'a']--        //将滑动前首位的词频删去
		sCount[s[i+pLen]-'a']++ //增加滑动后最后一位的词频（以此达到滑动的效果）
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
//func main() {
//	anagrams := findAnagrams("abab", "ab")
//	fmt.Println(anagrams)
//}
