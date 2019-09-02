/*
班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 的朋友。所谓的朋友圈，是指所有朋友的集合。

给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你必须输出所有学生中的已知的朋友圈总数。

示例 1:

输入:
[[1,1,0],
 [1,1,0],
 [0,0,1]]
输出: 2
说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
第2个学生自己在一个朋友圈。所以返回2。
示例 2:

输入:
[[1,1,0],
 [1,1,1],
 [0,1,1]]
输出: 1
说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
注意：

N 在[1,200]的范围内。
对于所有学生，有M[i][i] = 1。
如果有M[i][j] = 1，则有M[j][i] = 1。
*/

package five

// 并查集基础
// 1. 初始化一个并查集对象，遍历所有学生，将互为朋友的同学调用 Join 方法连接起来
// 2. 再依次遍历所有同学，调用 Find 方法找到其与其他共同的朋友，将该朋友标记为 1
// 3. 统计数组中被标记为 1 的个数即为答案

func findCircleNum(M [][]int) int {
	circle := NewFriendCircle(len(M))
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			if i != j && M[i][j] == 1 {
				circle.Join(i, j)
			}
		}
	}
	root := make([]int, len(M))
	for i := 0; i < len(M); i++ {
		root[circle.Find(i)] = 1
	}
	cnt := 0
	for _, n := range root {
		if n == 1 {
			cnt++
		}
	}
	return cnt
}

// FriendCircle ...
type FriendCircle struct {
	nums []int
}

// NewFriendCircle ...
func NewFriendCircle(n int) *FriendCircle {
	// nums 记录该学生对应的朋友索引
	nums := make([]int, n)
	// 初始化，每个人都是自个的朋友，没毛病
	for i := range nums {
		nums[i] = i
	}
	return &FriendCircle{nums}
}

// Find ...
func (fc *FriendCircle) Find(n int) int {
	ta := n
	// 沿着朋友链找，最终的那个一定是 nums[x] = x
	for ta != fc.nums[ta] {
		ta = fc.nums[ta]
	}
	// 此时 ta 就是最终那个点，需要把这条链上所有节点指向 ta
	// 压缩路径
	for i := n; i != ta; {
		j := fc.nums[i]
		fc.nums[i], i = ta, j
	}
	return ta
}

// Join ...
func (fc *FriendCircle) Join(a, b int) {
	x, y := fc.Find(a), fc.Find(b)
	if x != y {
		fc.nums[y] = x
	}
}
