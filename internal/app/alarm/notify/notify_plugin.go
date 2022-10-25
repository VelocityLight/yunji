package notify

type AlarmPlugin interface {
	SendAlarm(email string, content NotifyContent) error
}

type NotifyContent struct {
	Header string
	// Receiver email
	Blocks []Block
}

type Block struct {
	Text  string

	Links []Link
}

type Link struct {
	Href string
	Text string
}
