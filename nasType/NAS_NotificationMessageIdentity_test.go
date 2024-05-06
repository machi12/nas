package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/machi12/nas"
	"github.com/machi12/nas/nasType"
)

func TestNasTypeNewNotificationMessageIdentity(t *testing.T) {
	a := nasType.NewNotificationMessageIdentity()
	assert.NotNil(t, a)
}

type nasTypeNotificationMessageIdentityMessageType struct {
	in  uint8
	out uint8
}

var nasTypeNotificationMessageIdentityMessageTypeTable = []nasTypeNotificationMessageIdentityMessageType{
	{nas.MsgTypeNotification, nas.MsgTypeNotification},
}

func TestNasTypeGetSetNotificationMessageIdentityMessageType(t *testing.T) {
	a := nasType.NewNotificationMessageIdentity()
	for _, table := range nasTypeNotificationMessageIdentityMessageTypeTable {
		a.SetMessageType(table.in)
		assert.Equal(t, table.out, a.GetMessageType())
	}
}
