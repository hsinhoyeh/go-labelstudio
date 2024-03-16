/*
 * Label Studio API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MlInteractiveAnnotatingRequest struct {
	// ID of task to annotate
	Task int32 `json:"task"`
	// Context for ML model
	Context interface{} `json:"context,omitempty"`
}