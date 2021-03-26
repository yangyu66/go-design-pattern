package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserProxy_Login(t *testing.T) {
	proxy := NewUserProxy(&User{})

	err := proxy.Login("test", "password")
	n := proxy.GetName()
	assert.Equal(t, "ca", n)
	require.Nil(t, err)
}
