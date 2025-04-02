package entity

func checkTextContent(s string) error {
	if s == "" {
		return ErrTextContentCannotNotBlank
	}

	return nil
}
