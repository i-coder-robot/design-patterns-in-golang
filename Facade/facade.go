package Facade

type Facade struct {
	to string
	comment string
	validatorToken ValidatorToken
	validatorPermission ValidatorPermission
	store Storage
	notificator Email
}

func NewFacade(to, comment,token,user string) Facade{
	return Facade{
		to:                  to,
		comment:             comment,
		validatorToken:      NewValidatorToken(token),
		validatorPermission: NewValidatorPermission(user),
		store:               NewStorage("mysql"),
		notificator:         NewEmail(),
	}
}

func (f *Facade) Comment() error{
	err := f.validatorToken.Validate()
	if err != nil {
		return err
	}
	err = f.validatorPermission.Validate()
	if err != nil {
		return err
	}
	f.store.Save(f.comment)
	f.notificator.Send(f.to,f.comment)
	return nil
}

func (f *Facade) Notify() {
	f.notificator.Send(f.to,f.comment)
}