package Builder

type Sender struct {
	builder MessageBuilder
}

func (s *Sender) SetBuilder(b MessageBuilder) {
	s.builder = b
}

func (s *Sender) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	s.builder.SetRecipient(recipient).SetMessage(message)
	return s.builder.Build()
}
