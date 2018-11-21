package health

import (
	"reflect"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Healthcheck represents the actual health of some portion of the system
type Healthcheck struct {
	Component   string
	Measurement string
	Status      string
	Value       interface{}
	Unit        string
}

// ComponentHealthchecker is an interface that any component can implement to indicate that it perform a health check
type ComponentHealthchecker interface {
	// CheckHealth will check the health of this component
	CheckHealth(c chan<- Healthcheck)
}

// Healthchecker is the means to perform healthchecks on the system
type Healthchecker struct {
	checkers []ComponentHealthchecker
}

// New will create a new Healthchecker for the given ComponentHealthchecker instances
func New(componentHealthcheckers ...ComponentHealthchecker) Healthchecker {
	return Healthchecker{
		checkers: componentHealthcheckers,
	}
}

// CheckHealth will execute every healthcheck that is registered and return the composite of them all
func (h *Healthchecker) CheckHealth() []Healthcheck {
	results := []Healthcheck{}

	log.Info("Starting healthchecks")
	var wg sync.WaitGroup

	for i := range h.checkers {
		checker := h.checkers[i]
		log.WithField("checker", reflect.TypeOf(checker)).
			Info("Performing healthcheck")

		// This works asynchronously.
		// We create a channel to receive all healthcheck results - allowing each checker to return 0 or many results
		// We then run the checker in a new goroutine, and we run a goroutine per checker to receive the results and append
		// them into the response list
		wg.Add(2)
		checkResults := make(chan Healthcheck)
		go func() {
			defer wg.Done()
			checker.CheckHealth(checkResults)
			close(checkResults)
		}()
		go func() {
			defer wg.Done()
			for result := range checkResults {
				log.WithField("result", result).
					WithField("checker", reflect.TypeOf(checker)).
					Info("Received Healthcheck result")
				results = append(results, result)
			}
		}()
	}
	wg.Wait()

	return results
}
