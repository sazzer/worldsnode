package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testHealthcheck struct {
	result string
}

func (d testHealthcheck) CheckHealth(c chan<- Healthcheck) {
	c <- Healthcheck{
		Component:   "dummy",
		Measurement: "check",
		Status:      d.result,
	}
}

func TestNoHealthchecks(t *testing.T) {
	healthchecker := New()

	results := healthchecker.CheckHealth()

	assert.Equal(t, 0, len(results))
}

func TestSingleHealthcheck(t *testing.T) {
	tests := []string{Pass, Warn, Fail}

	for _, tt := range tests {
		healthchecker := New(testHealthcheck{tt})

		results := healthchecker.CheckHealth()

		assert.Equal(t, 1, len(results))
		assert.Equal(t, Healthcheck{Component: "dummy", Measurement: "check", Status: tt}, results[0])
	}
}

func TestMultipleHealthchecks(t *testing.T) {
	healthchecker := New(
		testHealthcheck{Pass},
		testHealthcheck{Warn},
		testHealthcheck{Fail},
	)

	results := healthchecker.CheckHealth()

	assert.Equal(t, 3, len(results))
	// TODO: Check the actual results
}
