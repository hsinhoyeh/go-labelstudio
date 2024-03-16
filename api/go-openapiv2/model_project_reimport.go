/*
 * Label Studio API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProjectReimport struct {
	Id               int32       `json:"id,omitempty"`
	Status           string      `json:"status,omitempty"`
	Error_           string      `json:"error,omitempty"`
	TaskCount        int32       `json:"task_count,omitempty"`
	AnnotationCount  int32       `json:"annotation_count,omitempty"`
	PredictionCount  int32       `json:"prediction_count,omitempty"`
	Duration         int32       `json:"duration,omitempty"`
	FileUploadIds    interface{} `json:"file_upload_ids,omitempty"`
	FilesAsTasksList bool        `json:"files_as_tasks_list,omitempty"`
	FoundFormats     interface{} `json:"found_formats,omitempty"`
	DataColumns      interface{} `json:"data_columns,omitempty"`
	Traceback        string      `json:"traceback,omitempty"`
	Project          int32       `json:"project,omitempty"`
}
