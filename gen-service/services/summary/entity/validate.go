package entity

func checkText(s string) error {
	if s == "" {
		return ErrContentNotValid
	}

	return nil
}
