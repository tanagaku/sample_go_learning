package meander_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanagaku/sample_go_learning/chapter7/meander"
)

func TestCostValues(t *testing.T) {
	assert.Equal(t, int(meander.Cost1), 1)
	assert.Equal(t, int(meander.Cost2), 2)
	assert.Equal(t, int(meander.Cost3), 3)
	assert.Equal(t, int(meander.Cost4), 4)
	assert.Equal(t, int(meander.Cost5), 5)
}

func TestCostString(t *testing.T) {
	assert.Equal(t, meander.Cost1.String(), "$")
	assert.Equal(t, meander.Cost2.String(), "$$")
	assert.Equal(t, meander.Cost3.String(), "$$$")
	assert.Equal(t, meander.Cost4.String(), "$$$$")
	assert.Equal(t, meander.Cost5.String(), "$$$$$")
}

func TestParseCost(t *testing.T) {
	assert.Equal(t, meander.Cost1, meander.ParseCost("$"))
	assert.Equal(t, meander.Cost2, meander.ParseCost("$$"))
	assert.Equal(t, meander.Cost3, meander.ParseCost("$$$"))
	assert.Equal(t, meander.Cost4, meander.ParseCost("$$$$"))
	assert.Equal(t, meander.Cost5, meander.ParseCost("$$$$$"))
}

func TestParseCostRange(t *testing.T) {
	var l *meander.CostRange
	l = meander.ParseCostRange("$$...$$$")
	assert.Equal(t, l.From, meander.Cost2)
	assert.Equal(t, l.To, meander.Cost3)
	l = meander.ParseCostRange("$...$$$$$")
	assert.Equal(t, l.From, meander.Cost1)
	assert.Equal(t, l.To, meander.Cost5)
}

func TestCostRangeString(t *testing.T) {
	assert.Equal(t, "$$...$$$$", (&meander.CostRange{
		From: meander.Cost2,
		To:   meander.Cost4,
	}).String())
}
