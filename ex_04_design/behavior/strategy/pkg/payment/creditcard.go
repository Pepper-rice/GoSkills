package payment

import "fmt"

// CreditCardStrategy 信用卡支付策略
type CreditCardStrategy struct{}

func (c *CreditCardStrategy) Pay(amount int) {
	fmt.Printf("使用信用卡支付了 %d 元\n", amount)
}
