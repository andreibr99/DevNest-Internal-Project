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
	"os"
)

type CreatedTimesheet struct {

	Id float64 `json:"id"`

	Status string `json:"status"`

	Date string `json:"date"`

	ProjectId float64 `json:"projectId"`

	File **os.File `json:"file,omitempty"`
}
