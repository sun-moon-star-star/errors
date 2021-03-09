package caller

type Http struct{}

func (h *Http) Call(*Caller) error {
	return nil
}
