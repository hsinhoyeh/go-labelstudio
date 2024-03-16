/*
 * Label Studio API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Task creation response
type TaskCreationResponse struct {
	// Number of tasks added
	TaskCount int32 `json:"task_count,omitempty"`
	// Number of annotations added
	AnnotationCount int32 `json:"annotation_count,omitempty"`
	// Number of predictions added
	PredictionsCount int32 `json:"predictions_count,omitempty"`
	// Time in seconds to create
	Duration float32 `json:"duration,omitempty"`
	// Database IDs of uploaded files
	FileUploadIds []int32 `json:"file_upload_ids,omitempty"`
	// Whether uploaded files can contain lists of tasks, like CSV/TSV files
	CouldBeTasksList bool `json:"could_be_tasks_list,omitempty"`
	// The list of found file formats
	FoundFormats []string `json:"found_formats,omitempty"`
	// The list of found data columns
	DataColumns []string `json:"data_columns,omitempty"`
}
