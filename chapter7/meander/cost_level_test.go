package meander_test

import (
	"sample_go_learning/chapter7/meander"
	"testing"

	"github.com/cheekybits/is"
)

func TestCostValues(t *testing.T) {
	is := is.New(t)

	is.Equal(int(meander.Cost1), 1)
	is.Equal(int(meander.Cost2), 2)
	is.Equal(int(meander.Cost3), 3)
	is.Equal(int(meander.Cost4), 4)
	is.Equal(int(meander.Cost5), 5)
}

func TestCostString(t *testing.T) {
	is := is.New(t)

	is.Equal(int(meander.Cost1.String()), "$")
	is.Equal(int(meander.Cost2.String()), "$$")
	is.Equal(int(meander.Cost3.String()), "$$$")
	is.Equal(int(meander.Cost4.String()), "$$$$")
	is.Equal(int(meander.Cost5.String()), "$$$$$")
}
