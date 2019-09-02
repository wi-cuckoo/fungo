/*
给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。

然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的最短时间。

示例 1：

输入: tasks = ["A","A","A","B","B","B"], n = 2
输出: 8
执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
注：

任务的总个数为 [1, 10000]。
n 的取值范围为 [0, 100]。
*/

package six

func leastInterval(tasks []byte, n int) int {
	// 计算每种任务数量
	table, max := [26]int{}, 0
	for _, t := range tasks {
		i := int(t - 'A')
		table[i]++
		if table[i] > max {
			max = table[i]
		}
	}
	schema := make([]int, max) // 记录插排的每个段任务数量
	for _, t := range table {
		for i := 0; i < t; i++ {
			schema[i]++
		}
	}

	// 依据每段饱和的原则，max-1 段的任务数应为 n+1，再加上最后一段长度
	cnt := (n+1)*(max-1) + schema[max-1]
	// 为何此处有如此判断呢，主要是因为如果前 max-1 段的坑位不够，或没有坑的极端情况下，则所有任务都排满，不需要任何待命
	// 换个方式也可以理解，最短时间至少是不小于任务数的
	if cnt < len(tasks) {
		return len(tasks)
	}
	return cnt
}
