package payment

import "fmt"

// AlipayStrategy 阿里支付策略
type AlipayStrategy struct{}

func (a *AlipayStrategy) Pay(amount int) {
	fmt.Printf("使用支付宝支付了 %d 元\n", amount)
}
