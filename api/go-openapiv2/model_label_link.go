/*
 * Label Studio API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type LabelLink struct {
	Id               int32 `json:"id,omitempty"`
	AnnotationsCount int32 `json:"annotations_count,omitempty"`
	// Tag name
	FromName string `json:"from_name"`
	Project  int32  `json:"project"`
	Label    int32  `json:"label"`
}
