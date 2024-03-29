/*
 * Label Studio API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

type ProjectImport struct {
	Id                     int32       `json:"id,omitempty"`
	PreannotatedFromFields interface{} `json:"preannotated_from_fields,omitempty"`
	CommitToProject        bool        `json:"commit_to_project,omitempty"`
	ReturnTaskIds          bool        `json:"return_task_ids,omitempty"`
	Status                 string      `json:"status,omitempty"`
	Url                    string      `json:"url,omitempty"`
	Traceback              string      `json:"traceback,omitempty"`
	Error_                 string      `json:"error,omitempty"`
	// Creation time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Updated time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Complete or fail time
	FinishedAt       time.Time   `json:"finished_at,omitempty"`
	TaskCount        int32       `json:"task_count,omitempty"`
	AnnotationCount  int32       `json:"annotation_count,omitempty"`
	PredictionCount  int32       `json:"prediction_count,omitempty"`
	Duration         int32       `json:"duration,omitempty"`
	FileUploadIds    interface{} `json:"file_upload_ids,omitempty"`
	CouldBeTasksList bool        `json:"could_be_tasks_list,omitempty"`
	FoundFormats     interface{} `json:"found_formats,omitempty"`
	DataColumns      interface{} `json:"data_columns,omitempty"`
	Tasks            interface{} `json:"tasks,omitempty"`
	TaskIds          interface{} `json:"task_ids,omitempty"`
	Project          int32       `json:"project,omitempty"`
}
