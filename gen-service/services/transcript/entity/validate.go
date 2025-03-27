package entity

func checkContent(s string) error {
	if s == "" {
		return ErrContentNotValid
	}

	return nil
}
