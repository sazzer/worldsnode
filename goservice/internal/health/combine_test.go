package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type combineHealthTests struct {
	old      string
	new      string
	expected string
}

func TestCombineHealth(t *testing.T) {
	tests := []combineHealthTests{
		{Pass, Pass, Pass},
		{Pass, Warn, Warn},
		{Pass, Fail, Fail},
		{Warn, Pass, Warn},
		{Warn, Warn, Warn},
		{Warn, Fail, Fail},
		{Fail, Pass, Fail},
		{Fail, Warn, Fail},
		{Fail, Fail, Fail},
	}

	for _, tt := range tests {
		newScore := combineHealth(tt.old, tt.new)
		assert.Equal(t, tt.expected, newScore)
	}
}
