package entity

func checkTitle(s string) error {
	if s == "" {
		return ErrTitleIsBlank
	}

	return nil
}
