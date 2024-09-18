package manager

import (
	"fmt"
	"github.com/learning-go/ex_04_design/behavior/strategy/pkg/payment"
	"strings"
)

type PaymentStrategyManager struct {
	strategyMap map[string]payment.PaymentStrategy
}

// NewPaymentStrategyManager 无参构造函数
func NewPaymentStrategyManager() *PaymentStrategyManager {
	return &PaymentStrategyManager{
		strategyMap: map[string]payment.PaymentStrategy{
			"ALIPAY":      &payment.AlipayStrategy{},
			"WECHAT":      &payment.WechatPayStrategy{},
			"CREDIT_CARD": &payment.CreditCardStrategy{},
		},
	}
}

// Pay 根据支付类型执行支付
func (p *PaymentStrategyManager) Pay(paymentType string, amount int) {
	strategy, exists := p.strategyMap[strings.ToUpper(paymentType)] // 类似Java的Map.get()
	if exists {
		strategy.Pay(amount)
	} else {
		fmt.Println("支付方式未找到")
	}
}
