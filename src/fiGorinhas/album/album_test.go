package album

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPacote(t *testing.T) {
	pacote := NewPacote(5, 550)
	assert.Equal(t, 5, len(pacote))
}
