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

type CreatedAllocation struct {

	Id float64 `json:"id,omitempty"`

	ProjectId float64 `json:"projectId,omitempty"`

	EmployeeId float64 `json:"employeeId,omitempty"`

	StartingDate time.Time `json:"startingDate,omitempty"`

	EndingDate time.Time `json:"endingDate,omitempty"`

	Type_ string `json:"type,omitempty"`

	Hours float64 `json:"hours,omitempty"`

	HourlyRate float64 `json:"hourlyRate,omitempty"`

	FlatRate float64 `json:"flatRate,omitempty"`

	FlatRateType string `json:"flatRateType,omitempty"`
}
