package payment

// PaymentStrategy 策略接口
type PaymentStrategy interface {
	Pay(amount int)
}
