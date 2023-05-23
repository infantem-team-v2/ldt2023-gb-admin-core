package terrors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorsMessage(t *testing.T) {
	t.Run("Testing parse file", TestInit)
	t.Run("Testing getting external message w/ code", TestGetExternalMessage)
}

func TestInit(t *testing.T) {
	Init()
	assert.NotEqual(t, nil, externalMessagesMap, "Is map w/ external msgs nil")
	t.Logf("Result from yaml: %v", externalMessagesMap)
	t.Logf("Random errorsMessage: %v", externalMessagesMap[100000])
}

func TestGetExternalMessage(t *testing.T) {
	Init()
	msg, code, err := getExternalMessage(100000)
	assert.Equal(t, nil, err, "Is err w/ external msgs nil")
	t.Logf("Code %d, Message: %v", code, msg)
}
