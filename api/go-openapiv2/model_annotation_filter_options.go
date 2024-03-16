/*
 * Label Studio API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AnnotationFilterOptions struct {
	// Include not skipped and not ground truth annotations
	Usual bool `json:"usual,omitempty"`
	// Include ground truth annotations
	GroundTruth bool `json:"ground_truth,omitempty"`
	// Include skipped annotations
	Skipped bool `json:"skipped,omitempty"`
}
