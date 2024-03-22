package project

import (
	"encoding/json"
	"time"
)

type annotationExport interface {
	Unmarshal([]byte) error
}

type BatchAnnotatioExportImageClassification []AnnotatioExportImageClassification

func (b *BatchAnnotatioExportImageClassification) Unmarshal(blob []byte) error {
	return json.Unmarshal(blob, b)
}

func (b BatchAnnotatioExportImageClassification) At(idx int) AnnotatioExportImageClassification {
	return []AnnotatioExportImageClassification(b)[idx]
}

func (b BatchAnnotatioExportImageClassification) Len() int {
	return len([]AnnotatioExportImageClassification(b))
}

type EachFunc func(a AnnotatioExportImageClassification)

func (b BatchAnnotatioExportImageClassification) ForEach(do EachFunc) {
	for _, v := range []AnnotatioExportImageClassification(b) {
		do(v)
	}
}

type AnnotatioExportImageClassification struct {
	ID          int `json:"id,omitempty"`
	Annotations []struct {
		ID          int        `json:"id,omitempty"`
		CompletedBy int        `json:"completed_by,omitempty"`
		Result      []struct { // NOTE: this field could be variant by label config
			ID    string `json:"id,omitempty"`
			Type  string `json:"type,omitempty"`
			Value struct {
				Choices []string `json:"choices,omitempty"`
			} `json:"value,omitempty"`
			Origin   string `json:"origin,omitempty"`
			ToName   string `json:"to_name,omitempty"`
			FromName string `json:"from_name,omitempty"`
		} `json:"result,omitempty"`
		WasCancelled   bool      `json:"was_cancelled,omitempty"`
		GroundTruth    bool      `json:"ground_truth,omitempty"`
		CreatedAt      time.Time `json:"created_at,omitempty"`
		UpdatedAt      time.Time `json:"updated_at,omitempty"`
		DraftCreatedAt any       `json:"draft_created_at,omitempty"`
		LeadTime       float64   `json:"lead_time,omitempty"`
		Prediction     struct {
		} `json:"prediction,omitempty"`
		ResultCount      int    `json:"result_count,omitempty"`
		UniqueID         string `json:"unique_id,omitempty"`
		ImportID         any    `json:"import_id,omitempty"`
		LastAction       any    `json:"last_action,omitempty"`
		Task             int    `json:"task,omitempty"`
		Project          int    `json:"project,omitempty"`
		UpdatedBy        int    `json:"updated_by,omitempty"`
		ParentPrediction any    `json:"parent_prediction,omitempty"`
		ParentAnnotation any    `json:"parent_annotation,omitempty"`
		LastCreatedBy    any    `json:"last_created_by,omitempty"`
	} `json:"annotations,omitempty"`
	Drafts      []any `json:"drafts,omitempty"`
	Predictions []any `json:"predictions,omitempty"`
	Data        struct {
		Image string `json:"image,omitempty"`
	} `json:"data,omitempty"`
	Meta struct {
	} `json:"meta,omitempty"`
	CreatedAt              time.Time `json:"created_at,omitempty"`
	UpdatedAt              time.Time `json:"updated_at,omitempty"`
	InnerID                int       `json:"inner_id,omitempty"`
	TotalAnnotations       int       `json:"total_annotations,omitempty"`
	CancelledAnnotations   int       `json:"cancelled_annotations,omitempty"`
	TotalPredictions       int       `json:"total_predictions,omitempty"`
	CommentCount           int       `json:"comment_count,omitempty"`
	UnresolvedCommentCount int       `json:"unresolved_comment_count,omitempty"`
	LastCommentUpdatedAt   any       `json:"last_comment_updated_at,omitempty"`
	Project                int       `json:"project,omitempty"`
	UpdatedBy              int       `json:"updated_by,omitempty"`
	CommentAuthors         []any     `json:"comment_authors,omitempty"`
}
