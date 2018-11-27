package health

import (
	"net/http"
	"strings"

	"github.com/unrolled/render"
)

var renderer *render.Render

func init() {
	renderer = render.New()
}

// HealthcheckDetailModel is the API Model representation of a single HealthCheck value
type healthcheckDetailModel struct {
	ObservedValue interface{} `json:"observedValue,omitempty"`
	ObservedUnit  string      `json:"observedUnit,omitempty"`
	Status        string      `json:"status"`
}

// HealthcheckModel is the API Model representation of the Healthcheck results
type healthcheckModel struct {
	Status  string                              `json:"status"`
	Details map[string][]healthcheckDetailModel `json:"details"`
}

// Get will perform the actual healthchecks and return the response
func (h *Healthchecker) get(w http.ResponseWriter, r *http.Request) {
	healthcheckResults := h.checkHealth()

	response := healthcheckModel{
		Status:  Pass,
		Details: make(map[string][]healthcheckDetailModel),
	}

	for _, healthcheckResult := range healthcheckResults {
		component := strings.Join([]string{healthcheckResult.Component, healthcheckResult.Measurement}, ":")

		response.Details[component] = append(response.Details[component], healthcheckDetailModel{
			Status:        healthcheckResult.Status,
			ObservedValue: healthcheckResult.Value,
			ObservedUnit:  healthcheckResult.Unit,
		})
		response.Status = combineHealth(response.Status, healthcheckResult.Status)
	}

	statusCode := 200
	if response.Status == Fail {
		statusCode = 500
	}

	renderer.JSON(w, statusCode, response)
}
