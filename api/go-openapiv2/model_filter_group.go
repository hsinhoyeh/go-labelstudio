/*
 * Label Studio API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type FilterGroup struct {
	Id      int32    `json:"id,omitempty"`
	Filters []Filter `json:"filters"`
	// Type of conjunction
	Conjunction string `json:"conjunction"`
}