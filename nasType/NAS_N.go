package nasType

// N
// N, sBit, len = [0, 15], 8 , 128
type N struct {
	Iei   uint8
	Octet [16]uint8
}

func NewN(iei uint8) (n *N) {
	n = &N{}
	n.SetIei(iei)
	return n
}

// N
// Iei Row, sBit, len = [], 8, 8
func (n *N) GetIei() (iei uint8) {
	return n.Iei
}

// N
// Iei Row, sBit, len = [], 8, 8
func (n *N) SetIei(iei uint8) {
	n.Iei = iei
}

// N
// NValue Row, sBit, len = [0, 15], 8 , 128
func (a *N) GetNValue() (n [16]uint8) {
	copy(n[:], a.Octet[0:16])
	return n
}

// N
// NValue Row, sBit, len = [0, 15], 8 , 128
func (a *N) SetNValue(n [16]uint8) {
	copy(a.Octet[0:16], n[:])
}
