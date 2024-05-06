package nasType_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/machi12/nas/nasType"
)

func TestNasTypeNewPlain5GSNASMessage(t *testing.T) {
	a := nasType.NewPlain5GSNASMessage()
	assert.NotNil(t, a)
}
