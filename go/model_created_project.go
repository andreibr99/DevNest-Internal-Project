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

type CreatedProject struct {

	Id float64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ClientName string `json:"clientName,omitempty"`

	Description string `json:"description,omitempty"`

	StartingDate time.Time `json:"startingDate,omitempty"`

	EndingDate time.Time `json:"endingDate,omitempty"`

	Currency string `json:"currency,omitempty"`

	Status bool `json:"status,omitempty"`
}
