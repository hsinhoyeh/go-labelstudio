package project

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	lshttp "github.com/hsinhoyeh/go-labelstudio/pkg/http"
)

func NewProjectService(client *lshttp.Client) *ProjectService {
	return &ProjectService{
		client: client,
	}
}

type ProjectService struct {
	client *lshttp.Client
}

func (p *ProjectService) CreateProject(ctx context.Context, title, description string, labelConfig LabelConfig) (*ProjectResource, error) {
	projectUrl, err := lshttp.JoinURL(p.client.HostURL(), "/api/projects")
	if err != nil {
		return nil, err
	}

	body := &CreateProjectRequestBody{
		Title:       title,
		Description: description,
		LabelConfig: labelConfig.String(),
	}

	req, err := http.NewRequestWithContext(ctx, "POST", projectUrl, body.ToJSONReader())
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Referer", projectUrl)

	rawResp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	resp := &ProjectResource{}
	if err := resp.Unmarshal(bytes.NewReader(rawResp)); err != nil {
		return nil, err
	}
	return resp, nil
}

type CreateProjectRequestBody struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	LabelConfig string `json:"label_config,omitempty"`
}

func (c *CreateProjectRequestBody) ToJSONReader() io.Reader {
	blob, _ := json.Marshal(c)
	return bytes.NewReader(blob)

}

type ProjectResource struct {
	ID                    int    `json:"id,omitempty"`
	Title                 string `json:"title,omitempty"`
	Description           string `json:"description,omitempty"`
	LabelConfig           string `json:"label_config,omitempty"`
	ExpertInstruction     string `json:"expert_instruction,omitempty"`
	ShowInstruction       bool   `json:"show_instruction,omitempty"`
	ShowSkipButton        bool   `json:"show_skip_button,omitempty"`
	EnableEmptyAnnotation bool   `json:"enable_empty_annotation,omitempty"`
	ShowAnnotationHistory bool   `json:"show_annotation_history,omitempty"`
	Organization          int    `json:"organization,omitempty"`
	Color                 string `json:"color,omitempty"`
	MaximumAnnotations    int    `json:"maximum_annotations,omitempty"`
	IsPublished           bool   `json:"is_published,omitempty"`
	ModelVersion          string `json:"model_version,omitempty"`
	IsDraft               bool   `json:"is_draft,omitempty"`
	CreatedBy             struct {
		ID        int    `json:"id,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Email     string `json:"email,omitempty"`
		Avatar    any    `json:"avatar,omitempty"`
	} `json:"created_by,omitempty"`
	CreatedAt                       time.Time `json:"created_at,omitempty"`
	MinAnnotationsToStartTraining   int       `json:"min_annotations_to_start_training,omitempty"`
	StartTrainingOnAnnotationUpdate bool      `json:"start_training_on_annotation_update,omitempty"`
	ShowCollabPredictions           bool      `json:"show_collab_predictions,omitempty"`
	NumTasksWithAnnotations         int       `json:"num_tasks_with_annotations,omitempty"`
	TaskNumber                      int       `json:"task_number,omitempty"`
	UsefulAnnotationNumber          int       `json:"useful_annotation_number,omitempty"`
	GroundTruthNumber               int       `json:"ground_truth_number,omitempty"`
	SkippedAnnotationsNumber        int       `json:"skipped_annotations_number,omitempty"`
	TotalAnnotationsNumber          int       `json:"total_annotations_number,omitempty"`
	TotalPredictionsNumber          int       `json:"total_predictions_number,omitempty"`
	Sampling                        string    `json:"sampling,omitempty"`
	ShowGroundTruthFirst            bool      `json:"show_ground_truth_first,omitempty"`
	ShowOverlapFirst                bool      `json:"show_overlap_first,omitempty"`
	OverlapCohortPercentage         int       `json:"overlap_cohort_percentage,omitempty"`
	TaskDataLogin                   any       `json:"task_data_login,omitempty"`
	TaskDataPassword                any       `json:"task_data_password,omitempty"`
	ControlWeights                  struct {
		Kp1 struct {
			Overall float64 `json:"overall,omitempty"`
			Type    string  `json:"type,omitempty"`
			Labels  struct {
				Face float64 `json:"Face,omitempty"`
				Nose float64 `json:"Nose,omitempty"`
			} `json:"labels,omitempty"`
		} `json:"kp-1,omitempty"`
		Choice struct {
			Type   string `json:"type,omitempty"`
			Labels struct {
				Weapons      float64 `json:"Weapons,omitempty"`
				Violence     float64 `json:"Violence,omitempty"`
				AdultContent float64 `json:"Adult content,omitempty"`
			} `json:"labels,omitempty"`
			Overall float64 `json:"overall,omitempty"`
		} `json:"choice,omitempty"`
		Label struct {
			Type   string `json:"type,omitempty"`
			Labels struct {
				EventA float64 `json:"Event A,omitempty"`
				EventB float64 `json:"Event B,omitempty"`
			} `json:"labels,omitempty"`
			Overall float64 `json:"overall,omitempty"`
		} `json:"label,omitempty"`
	} `json:"control_weights,omitempty"`
	ParsedLabelConfig struct {

		// Kp1 is for key point segmentation
		Kp1 struct {
			Type   string   `json:"type,omitempty"`
			ToName []string `json:"to_name,omitempty"`
			Inputs []struct {
				Type  string `json:"type,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"inputs,omitempty"`
			Labels      []string `json:"labels,omitempty"`
			LabelsAttrs struct {
				Face struct {
					Value      string `json:"value,omitempty"`
					Background string `json:"background,omitempty"`
				} `json:"Face,omitempty"`
				Nose struct {
					Value      string `json:"value,omitempty"`
					Background string `json:"background,omitempty"`
				} `json:"Nose,omitempty"`
			} `json:"labels_attrs,omitempty"`
		} `json:"kp-1,omitempty"`

		// Choice is for image classification
		Choice struct {
			Type   string `json:"type,omitempty"`
			Inputs []struct {
				Type  string `json:"type,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"inputs,omitempty"`
			Labels      []string `json:"labels,omitempty"`
			ToName      []string `json:"to_name,omitempty"`
			LabelsAttrs struct {
				Weapons struct {
					Value string `json:"value,omitempty"`
				} `json:"Weapons,omitempty"`
				Violence struct {
					Value string `json:"value,omitempty"`
				} `json:"Violence,omitempty"`
				AdultContent struct {
					Value string `json:"value,omitempty"`
				} `json:"Adult content,omitempty"`
			} `json:"labels_attrs,omitempty"`
		} `json:"choice,omitempty"`

		// Label is for audio labeling
		Label struct {
			Type   string `json:"type,omitempty"`
			Inputs []struct {
				Type  string `json:"type,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"inputs,omitempty"`
			Labels      []string `json:"labels,omitempty"`
			ToName      []string `json:"to_name,omitempty"`
			LabelsAttrs struct {
				EventA struct {
					Value      string `json:"value,omitempty"`
					Background string `json:"background,omitempty"`
				} `json:"Event A,omitempty"`
				EventB struct {
					Value      string `json:"value,omitempty"`
					Background string `json:"background,omitempty"`
				} `json:"Event B,omitempty"`
			} `json:"labels_attrs,omitempty"`
		} `json:"label,omitempty"`
	} `json:"parsed_label_config,omitempty"`
	EvaluatePredictionsAutomatically  bool   `json:"evaluate_predictions_automatically,omitempty"`
	ConfigHasControlTags              bool   `json:"config_has_control_tags,omitempty"`
	SkipQueue                         string `json:"skip_queue,omitempty"`
	RevealPreannotationsInteractively bool   `json:"reveal_preannotations_interactively,omitempty"`
	PinnedAt                          any    `json:"pinned_at,omitempty"`
	FinishedTaskNumber                int    `json:"finished_task_number,omitempty"`
	QueueTotal                        int    `json:"queue_total,omitempty"`
	QueueDone                         int    `json:"queue_done,omitempty"`
}

func (c *ProjectResource) Unmarshal(b io.Reader) error {
	blob, err := io.ReadAll(b)
	if err != nil {
		return err
	}

	return json.Unmarshal(blob, c)
}

func (p *ProjectService) GetProject(ctx context.Context, pid int) (*ProjectResource, error) {
	projectUrl, err := lshttp.JoinURL(p.client.HostURL(), fmt.Sprintf("/api/projects/%d", pid))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", projectUrl, nil)
	if err != nil {
		return nil, err
	}

	//req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Referer", projectUrl)

	rawResp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}

	resp := &ProjectResource{}
	if err := resp.Unmarshal(bytes.NewReader(rawResp)); err != nil {
		return nil, err
	}
	return resp, nil

}
