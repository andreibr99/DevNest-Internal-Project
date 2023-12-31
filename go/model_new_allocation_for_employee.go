/*
 * DevNest Portal API (OpenAPI 3.0)
 *
 * The DevNest Portal API endpoints definitions based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.0
 * Contact: contact@devnest.ro
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type NewAllocationForEmployee struct {

	ProjectId float64 `json:"projectId,omitempty"`

	StartingDate time.Time `json:"startingDate"`

	EndingDate time.Time `json:"endingDate,omitempty"`

	Type_ string `json:"type"`

	Hours float64 `json:"hours"`

	HourlyRate float64 `json:"hourlyRate"`

	FlatRate float64 `json:"flatRate"`

	FlatRateType string `json:"flatRateType,omitempty"`
}
