package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	got := Add(1,1)
	want := 2
	assert.Equal(t, got, want, "got and want should be the same")
}
