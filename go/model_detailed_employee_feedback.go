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

type DetailedEmployeeFeedback struct {

	Id float64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	OverallRating float32 `json:"overallRating,omitempty"`

	TeamAtmosphere float32 `json:"teamAtmosphere,omitempty"`

	Workload float32 `json:"workload,omitempty"`

	ClientCommunication float32 `json:"clientCommunication,omitempty"`

	MyInvolvementInTheProject float32 `json:"myInvolvementInTheProject,omitempty"`

	TechnicalResults float32 `json:"technicalResults,omitempty"`
}
