package mail

type Mail struct {
	From    string
	To      []string
	Title   []byte
	Message []byte
}

type Service interface {
	Send(mail Mail) (err error)
	Fetch() (mails []Mail, err error) // 阅后即焚
}
