package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/machi12/nas"
	"github.com/machi12/nas/nasType"
)

func TestNasTypeNewPDUSESSIONESTABLISHMENTACCEPTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTACCEPTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionEstablishmentAccept, nas.MsgTypePDUSessionEstablishmentAccept},
}

func TestNasTypeGetSetPDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONESTABLISHMENTACCEPTMessageIdentity()
	for _, table := range nasTypePDUSESSIONESTABLISHMENTACCEPTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
