package internal

type Validatable interface {
	Validate() error
}

func CheckValid(v Validatable) error {
	if err := v.Validate(); err != nil {
		return err
	}
	return nil
}
