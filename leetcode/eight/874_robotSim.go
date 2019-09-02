/*
机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：

-2：向左转 90 度
-1：向右转 90 度
1 <= x <= 9：向前移动 x 个单位长度
在网格上有一些格子被视为障碍物。

第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])

如果机器人试图走到障碍物上方，那么它将停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。

返回从原点到机器人的最大欧式距离的平方。
*/

package eight

func robotSim(commands []int, obstacles [][]int) int {
	return 0
}

// https://leetcode-cn.com/problems/delete-columns-to-make-sorted/
func minDeletionSize(A []string) int {
	cnt := 0
	for i := 0; i < len(A[0]); i++ {
		for j := 1; j < len(A); j++ {
			if A[j-1][i] > A[j][i] {
				cnt++
				break
			}
		}
	}
	return cnt
}

// https://leetcode-cn.com/problems/string-without-aaa-or-bbb/
func strWithout3a3b(A int, B int) string {
	buf := make([]byte, 0, A+B)
	flag := A > B // true for adding 'a', false to add 'b'
	for ; A > 0 && B > 0; flag = !flag {
		// fmt.Println(flag, A, B, string(buf))
		// 如果 A 多，则当前添加两个'a'或者一个'b'
		if A > B {
			if flag {
				buf, A = append(buf, 'a', 'a'), A-2
			} else {
				buf, B = append(buf, 'b'), B-1
			}
		} else if A < B {
			if flag {
				buf, A = append(buf, 'a'), A-1
			} else {
				buf, B = append(buf, 'b', 'b'), B-2
			}
		} else {
			if flag {
				buf, A = append(buf, 'a'), A-1
			} else {
				buf, B = append(buf, 'b'), B-1
			}
		}
	}
	for i := 0; i < A; i++ {
		buf = append(buf, 'a')
	}
	for i := 0; i < B; i++ {
		buf = append(buf, 'b')
	}
	return string(buf)
}

func strWithout3a3bV2(A int, B int) string {
	buf := make([]byte, 0, A+B)
	for A > 0 || B > 0 {
		var writeA bool
		if len(buf) >= 2 && buf[len(buf)-1] == buf[len(buf)-2] {
			writeA = buf[len(buf)-1] == 'b'
		} else {
			writeA = A >= B
		}
		if writeA {
			buf = append(buf, 'a')
			A--
		} else {
			buf = append(buf, 'b')
			B--
		}
	}
	return string(buf)
}

// https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/
func largestSumAfterKNegations(A []int, K int) int {
	for ; K > 0; K-- {
		t := 0
		for i := 1; i < len(A); i++ {
			if A[i] > A[t] {
				t = i
			}
		}
		A[t] = -A[t]
	}
	sum := 0
	for _, n := range A {
		sum += n
	}
	return sum
}
