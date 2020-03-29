package Builder

import "encoding/xml"

type XMLMessageBuilder struct {
	message Message
}

func (b *XMLMessageBuilder) SetRecipient(reciipient string) MessageBuilder {
	b.message.Recipient=reciipient
	return b
}

func (b *XMLMessageBuilder) SetMessage(text string) MessageBuilder  {
	b.message.Text=text
	return b
}

func (b *XMLMessageBuilder) Build() (*MessageRepresented ,error) {
	data, err := xml.Marshal(b.message)
	if err != nil {
		return nil,err
	}
	return &MessageRepresented{
		Body:   data,
		Format: "XML",
	},nil
}

