package entity

func checkMindmapData(s string) error {
	if s == "" {
		return ErrMindmapDataNotValid
	}

	return nil
}
