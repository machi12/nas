package nasType

// NOTE: AuthenticationParameterSNMAC
// SNMACValue Row, sBit, len = [0, 7], 8 , 64
type AuthenticationParameterSNMAC struct {
	Iei   uint8
	Octet [8]uint8
}

func NewAuthenticationParameterSNMAC(iei uint8) (authenticationParameterSNMAC *AuthenticationParameterSNMAC) {
	authenticationParameterSNMAC = &AuthenticationParameterSNMAC{}
	authenticationParameterSNMAC.SetIei(iei)
	return authenticationParameterSNMAC
}

// Iei Row, sBit, len = [], 8, 8
func (a *AuthenticationParameterSNMAC) GetIei() (iei uint8) {
	return a.Iei
}

// Iei Row, sBit, len = [], 8, 8
func (a *AuthenticationParameterSNMAC) SetIei(iei uint8) {
	a.Iei = iei
}

// SNMACValue Row, sBit, len = [0, 7], 8, 64
func (a *AuthenticationParameterSNMAC) GetSNMACValue() (sNMACValue [8]uint8) {
	copy(sNMACValue[:], a.Octet[0:8])
	return sNMACValue
}

// AuthenticationParameterSNMAC
// SNMACValue Row, sBit, len = [0, 7], 8, 64
func (a *AuthenticationParameterSNMAC) SetSNMACValue(sNMACValue [8]uint8) {
	copy(a.Octet[0:8], sNMACValue[:])
}
