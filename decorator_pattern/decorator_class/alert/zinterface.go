package alert

// Alert :
type Alert interface {
	Notify(msg string)
}

// Logger :
type Logger interface {
	Log(msg string)
}
