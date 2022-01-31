package meander_test

import (
	"testing"
	"go get github.com/cheekybits/is"
	"github.com/tanagaku/sample_go_learning/chapter7/meander"
)

func TestCostValues(t *testing.T) {
	is := is.New(t)

	is.Equal(int(meander.Cost1), 1)
	is.Equal(int(meander.Cost2), 2)
	is.Equal(int(meander.Cost3), 3)
	is.Equal(int(meander.Cost4), 4)
	is.Equal(int(meander.Cost5), 5)
}
