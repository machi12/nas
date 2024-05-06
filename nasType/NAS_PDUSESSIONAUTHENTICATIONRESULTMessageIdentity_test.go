package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/machi12/nas"
	"github.com/machi12/nas/nasType"
)

func TestNasTypeNewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageTypeTable = []nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType{
	{nas.MsgTypePDUSessionAuthenticationResult, nas.MsgTypePDUSessionAuthenticationResult},
}

func TestNasTypeGetSetPDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewPDUSESSIONAUTHENTICATIONRESULTMessageIdentity()
	for _, table := range nasTypePDUSESSIONAUTHENTICATIONRESULTMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
