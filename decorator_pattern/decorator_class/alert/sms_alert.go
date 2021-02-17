package alert

import "fmt"

// SMSAlert :
type SMSAlert struct {
	ReceiverPhone string
}

// Notify :
func (sms SMSAlert) Notify(msg string) {
	fmt.Printf("send sms message: %v to %v\n", msg, sms.ReceiverPhone)
}
