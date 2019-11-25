/*
 给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

 你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

 文本的最后一行应为左对齐，且单词之间不插入额外的空格。

 说明:

 单词是指由非空格字符组成的字符序列。
 每个单词的长度大于 0，小于等于 maxWidth。
 输入单词数组 words 至少包含一个单词。
 示例:

 输入:
 words = ["This", "is", "an", "example", "of", "text", "justification."]
 maxWidth = 16
 输出:
 [
    "This    is    an",
    "example  of text",
    "justification.  "
 ]
*/

package zero

import "strings"

// 思路挺简单的：就是上满一车就走，多退少补
// 1. 遍历所有单词分段，原则是一组单词长度和加上所需最少空格数不能超过 maxWidth
// 2. 对分好的每组，计算需要补多少空格，算法也简单
// 假设一组单词有 n 个，需要插入空格则是 w = maxWidth-sum(n), 其中sum(n)表示n个单词长度和
// 那么每个插入位置至少要分配得到 w/(n-1) 个空格，然后余下的 w%(n-1) 个空格可以随机分配
// 3. 需要特别处理的是最后一组，最后一组每个单词间隔一个空格，不足maxWidth的部分补上空格就行

func fullJustify(words []string, maxWidth int) []string {
	lines := make([]string, 0)
	width, left := 0, 0
	for i := 0; i < len(words); i++ {
		width += len(words[i])
		space := i - left
		// 如果加上该单词长度超出，则不能加上这个逼
		if width+space > maxWidth {
			lines = append(lines, jbspace(words[left:i], width-len(words[i]), maxWidth))
			width, left = len(words[i]), i
		}
	}
	if left < len(words) {
		last := strings.Join(words[left:], " ")
		lines = append(lines, last+nbspace(maxWidth-len(last)))
	}
	return lines
}

func jbspace(s []string, w, mw int) string {
	if len(s) == 1 {
		return s[0] + nbspace(mw-w)
	}
	n := (mw - w) / (len(s) - 1)
	m := (mw - w) % (len(s) - 1)
	ns, ts := nbspace(n), ""
	for i := 0; i < len(s)-1; i++ {
		ts += s[i] + ns
		if i < m {
			ts += " "
		}
	}
	return ts + s[len(s)-1]
}

func nbspace(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}
