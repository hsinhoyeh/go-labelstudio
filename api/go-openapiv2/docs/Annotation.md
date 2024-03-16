# Annotation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**CreatedUsername** | **string** | Username string | [optional] [default to null]
**CreatedAgo** | **string** | Time delta from creation time | [optional] [default to null]
**CompletedBy** | **int32** |  | [optional] [default to null]
**UniqueId** | **string** |  | [optional] [default to null]
**Result** | **interface{}** | The main value of annotator work - labeling result in JSON format | [optional] [default to null]
**WasCancelled** | **bool** | User skipped the task | [optional] [default to null]
**GroundTruth** | **bool** | This annotation is a Ground Truth (ground_truth) | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Last updated time | [optional] [default to null]
**DraftCreatedAt** | [**time.Time**](time.Time.md) | Draft creation time | [optional] [default to null]
**LeadTime** | **float32** | How much time it took to annotate the task | [optional] [default to null]
**ImportId** | **int32** | Original annotation ID that was at the import step or NULL if this annotation wasn&#39;t imported | [optional] [default to null]
**LastAction** | **string** | Action which was performed in the last annotation history item | [optional] [default to null]
**Task** | **int32** | Corresponding task for this annotation | [optional] [default to null]
**Project** | **int32** | Project ID for this annotation | [optional] [default to null]
**UpdatedBy** | **int32** | Last user who updated this annotation | [optional] [default to null]
**ParentPrediction** | **int32** | Points to the prediction from which this annotation was created | [optional] [default to null]
**ParentAnnotation** | **int32** | Points to the parent annotation from which this annotation was created | [optional] [default to null]
**LastCreatedBy** | **int32** | User who created the last annotation history item | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


