package main

import (
	"sort"
)

//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例: 
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//] 
//
// 说明： 
//
// 
// 所有输入均为小写字母。 
// 不考虑答案输出的顺序。 
// 
// Related Topics 哈希表 字符串 排序 
// 👍 772 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
type SortRunes []rune

func (s SortRunes) Len() int {
	return len(s)
}

func (s SortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var _ sort.Interface = SortRunes{}

func groupAnagrams(strs []string) [][]string {
	mapping := make(map[string][]string)
	ans := make([][]string, 0)

	for _, s := range strs {
		sortedString := SortRunes(s)
		sort.Sort(sortedString)
		value := mapping[string(sortedString)]
		value = append(value, s)
		mapping[string(sortedString)] = value
	}

	for _, v := range mapping {
		ans = append(ans, v)
	}
	return ans
}
//leetcode submit region end(Prohibit modification and deletion)
