package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/machi12/nas"
	"github.com/machi12/nas/nasType"
)

func TestNasTypeNewPDUSESSIONMODIFICATIONREJECTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONREJECTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONMODIFICATIONREJECTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONMODIFICATIONREJECTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONMODIFICATIONREJECTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionModificationReject, nas.MsgTypePDUSessionModificationReject},
}

func TestNasTypeGetSetPDUSESSIONMODIFICATIONREJECTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONMODIFICATIONREJECTMessageIdentity()
	for _, table := range nasTypePDUSESSIONMODIFICATIONREJECTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
