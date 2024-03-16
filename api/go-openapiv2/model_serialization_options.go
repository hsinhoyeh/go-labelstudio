/*
 * Label Studio API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SerializationOptions struct {
	Drafts                 *SerializationOption `json:"drafts,omitempty"`
	Predictions            *SerializationOption `json:"predictions,omitempty"`
	AnnotationsCompletedBy *SerializationOption `json:"annotations__completed_by,omitempty"`
	// Interpolate video key frames
	InterpolateKeyFrames bool `json:"interpolate_key_frames,omitempty"`
}
