package Builder
type Message struct {
	Recipient string `json:"recipient" xml:"recipient"`
	Text      string `json:"text" xml:"text"`
}

type MessageRepresented struct {
	Body []byte
	Format string
}
