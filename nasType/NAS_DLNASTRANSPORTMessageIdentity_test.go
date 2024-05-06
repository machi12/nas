package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/machi12/nas"
	"github.com/machi12/nas/nasType"
)

type nasTypeDLNASTRANSPORTMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeDLNASTRANSPORTMessageIdentityTable = []nasTypeDLNASTRANSPORTMessageIdentityData{
	{nas.MsgTypeDLNASTransport, nas.MsgTypeDLNASTransport},
}

func TestNasTypeNewDLNASTRANSPORTMessageIdentity(t *testing.T) {
	a := nasType.NewDLNASTRANSPORTMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetDLNASTRANSPORTMessageIdentity(t *testing.T) {
	a := nasType.NewDLNASTRANSPORTMessageIdentity()
	for _, table := range nasTypeDLNASTRANSPORTMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
