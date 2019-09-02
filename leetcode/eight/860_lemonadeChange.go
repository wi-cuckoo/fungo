/*
 在柠檬水摊上，每一杯柠檬水的售价为 5 美元。

顾客排队购买你的产品，（按账单 bills 支付的顺序）一次购买一杯。

每位顾客只买一杯柠檬水，然后向你付 5 美元、10 美元或 20 美元。你必须给每个顾客正确找零，也就是说净交易是每位顾客向你支付 5 美元。

注意，一开始你手头没有任何零钱。

如果你能给每位顾客正确找零，返回 true ，否则返回 false 。
*/

package eight

func lemonadeChange(bills []int) bool {
	changes := [...]int{0, 0, 0, 0} // 5, 10, 15, 20
	for _, b := range bills {
		// 更新一下钱包中5块数量
		i := b/5 - 1
		changes[i]++
		// 如果是 5 块就不用找零
		if b == 5 {
			continue
		}
		b -= 5 // 需要找零的钱
		// fmt.Println("befor change", i, b, changes)
		for i--; i >= 0; i-- {
			val := 5 * (i + 1)
			if n := b / val; n > 0 {
				if n > changes[i] {
					n = changes[i]
				}
				b -= n * val
				changes[i] -= n
				if b <= 0 {
					break
				}
			}
		}
		if b > 0 {
			return false
		}
	}
	return true
}
