package four

/*
基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。
给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。


示例 1：
输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
输出：2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-genetic-mutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 用回溯的思路，对于基因序列的任一位置，都进行四种变化的可能性尝试。只要在基因库中存在就继续下探
// 对于已经在基因库中访问过的序列，进行标记，不再重复寻找
func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}

	table := map[string]struct{}{start: {}}
	for _, b := range bank {
		table[b] = struct{}{}
	}

	// 确认能不能找到对应基因库序列
	minStep := len(bank) + 1
	visited := map[string]struct{}{}
	var backtrack func(cur []byte, step int)
	backtrack = func(cur []byte, step int) {
		if _, ok := table[string(cur)]; !ok {
			return
		}
		if _, ok := visited[string(cur)]; ok {
			return
		}
		visited[string(cur)] = struct{}{}
		// fmt.Println("accept: ", string(cur))
		if string(cur) == end {
			if step < minStep {
				minStep = step
			}
			return
		}
		for i := 0; i < 8; i++ {
			for _, c := range []byte{'A', 'C', 'G', 'T'} {
				pre := cur[i]
				if pre == c {
					continue
				}
				cur[i] = c
				backtrack(cur, step+1)
				cur[i] = pre
			}
		}
	}

	backtrack([]byte(start), 0)

	if minStep < len(bank)+1 {
		return minStep
	}
	return -1
}
