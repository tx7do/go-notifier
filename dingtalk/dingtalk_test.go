package dingtalk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNotifier(t *testing.T) {
	c := NewNotifier()
	assert.NotNil(t, c)
}
