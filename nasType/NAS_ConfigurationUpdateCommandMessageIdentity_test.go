package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/machi12/nas"
	"github.com/machi12/nas/nasType"
)

type nasTypeConfigurationUpdateCommandMessageIdentityData struct {
	in  uint8
	out uint8
}

var nasTypeConfigurationUpdateCommandMessageIdentityTable = []nasTypeConfigurationUpdateCommandMessageIdentityData{
	{nas.MsgTypeConfigurationUpdateCommand, nas.MsgTypeConfigurationUpdateCommand},
}

func TestNasTypeNewConfigurationUpdateCommandMessageIdentity(t *testing.T) {
	a := nasType.NewConfigurationUpdateCommandMessageIdentity()
	assert.NotNil(t, a)
}

func TestNasTypeGetSetConfigurationUpdateCommandMessageIdentity(t *testing.T) {
	a := nasType.NewConfigurationUpdateCommandMessageIdentity()
	for _, table := range nasTypeConfigurationUpdateCommandMessageIdentityTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
