package payment

import "fmt"

// WechatPayStrategy 微信支付策略
type WechatPayStrategy struct{}

func (w *WechatPayStrategy) Pay(amount int) {
	fmt.Printf("使用微信支付了 %d 元\n", amount)
}
