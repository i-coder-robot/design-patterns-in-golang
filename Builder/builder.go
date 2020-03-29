package Builder

type MessageBuilder interface {
	SetRecipient(recipient string) MessageBuilder
	SetMessage(message string) MessageBuilder
	//Represent
	//vt. 代表；表现；描绘；回忆；再赠送
	//vi. 代表；提出异议
	Build() (*MessageRepresented,error)
}
