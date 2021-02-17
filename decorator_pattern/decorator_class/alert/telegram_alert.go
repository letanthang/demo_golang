package alert

import "fmt"

// TelegramAlert :
type TelegramAlert struct {
	ReceiverAccount string
}

// Notify :
func (telegram TelegramAlert) Notify(msg string) {
	fmt.Printf("send telegram message: %v to %v\n", msg, telegram.ReceiverAccount)
}
