package types

type SpecialForm Symbol

func (s SpecialForm) IsSpecialForm() bool {
	return true
}
