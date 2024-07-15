package qutil

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendManyTransferPayload_Size(t *testing.T) {
	var payload SendManyTransferPayload
	b, err := payload.MarshallBinary()
	require.NoError(t, err, "binary marshalling payload")
	require.True(t, len(b) == SendManyInputSize)
}
