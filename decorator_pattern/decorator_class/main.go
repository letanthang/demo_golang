package main

import (
	"app/decorator_pattern/decorator_class/alert"
)

func main() {

	logClient := alert.LogrusLogger{}
	alertClient := alert.SMSAlert{ReceiverPhone: "0933932173"}

	logAndAlert := alert.AlertLogDecorator{LogClient: logClient, AlertClient: alertClient}
	logAndAlert.Notify("This is Prime Optimus to all Autobot")

	logAndAlertEmail := alert.AlertLogDecorator{logClient, alert.EmailAlert{ReceiverEmail: "letanthang@gmail.com"}}
	logAndAlertEmail.Notify("I'm Yoo and she is Mi")
}
