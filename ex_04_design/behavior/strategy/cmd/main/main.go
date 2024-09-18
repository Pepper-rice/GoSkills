package main

import (
	"github.com/learning-go/ex_04_design/behavior/strategy/internal/manager"
)

func main() {
	manager := manager.NewPaymentStrategyManager()
	manager.Pay("ALIPAY", 100)
	manager.Pay("WECHAT", 200)
	manager.Pay("CREDIT_CARD", 300)
	manager.Pay("UNKNOWN", 400) // 示例错误处理
}
