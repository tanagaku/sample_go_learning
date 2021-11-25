package option

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Open(t *testing.T) {
	l := log.Logger{}
	r, err := Open(WithFlag(false), WithParam("test"), WithLogger(&l))
	assert.Nil(t, err)

	expected := options{param: "test", flag: false, logger: &l}
	assert.Equal(t, expected, r)
}
