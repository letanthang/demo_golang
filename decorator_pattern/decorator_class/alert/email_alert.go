package alert

import "fmt"

// EmailAlert :
type EmailAlert struct {
	ReceiverEmail string
}

// Notify :
func (email EmailAlert) Notify(msg string) {
	fmt.Printf("send email message: %v to %v\n", msg, email.ReceiverEmail)
}
