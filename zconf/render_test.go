package zconf

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRender(t *testing.T) {
	_ = os.Setenv("TEST_KEY1", "VAL1")
	out, err := Render([]byte(`{{osHostname}}{{.Env.TEST_KEY1}}{{jsonMarshal "hello world"}}`))
	require.NoError(t, err)
	hostname, _ := os.Hostname()
	require.Equal(t, hostname+"VAL1"+`"hello world"`, string(out))
}
