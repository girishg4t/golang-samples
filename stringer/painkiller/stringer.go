package painkiller

//Pill to use out side
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

//PillType used
type PillType struct {
	Pill
}

func (p PillType) String() string {
	return p.Pill.String()
}
