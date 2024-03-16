# DataManagerTaskSerializer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Predictions** | **string** |  | [optional] [default to null]
**Annotations** | [**[]Annotation**](Annotation.md) |  | [optional] [default to null]
**Drafts** | **string** |  | [optional] [default to null]
**Annotators** | **string** |  | [optional] [default to null]
**InnerId** | **int32** |  | [optional] [default to null]
**CancelledAnnotations** | **int32** |  | [optional] [default to null]
**TotalAnnotations** | **int32** |  | [optional] [default to null]
**TotalPredictions** | **int32** |  | [optional] [default to null]
**CompletedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**AnnotationsResults** | **string** |  | [optional] [default to null]
**PredictionsResults** | **string** |  | [optional] [default to null]
**PredictionsScore** | **float32** |  | [optional] [default to null]
**FileUpload** | **string** |  | [optional] [default to null]
**StorageFilename** | **string** |  | [optional] [default to null]
**AnnotationsIds** | **string** |  | [optional] [default to null]
**PredictionsModelVersions** | **string** |  | [optional] [default to null]
**AvgLeadTime** | **float32** |  | [optional] [default to null]
**DraftExists** | **bool** |  | [optional] [default to null]
**UpdatedBy** | **string** |  | [optional] [default to null]
**Data** | **interface{}** | User imported or uploaded data for a task. Data is formatted according to the project label config. You can find examples of data for your project on the Import page in the Label Studio Data Manager UI. | [default to null]
**Meta** | **interface{}** | Meta is user imported (uploaded) data and can be useful as input for an ML Backend for embeddings, advanced vectors, and other info. It is passed to ML during training/predicting steps. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Time a task was created | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Last time a task was updated | [optional] [default to null]
**IsLabeled** | **bool** | True if the number of annotations for this task is greater than or equal to the number of maximum_completions for the project | [optional] [default to null]
**Overlap** | **int32** | Number of distinct annotators that processed the current task | [optional] [default to null]
**CommentCount** | **int32** | Number of comments in the task including all annotations | [optional] [default to null]
**UnresolvedCommentCount** | **int32** | Number of unresolved comments in the task including all annotations | [optional] [default to null]
**LastCommentUpdatedAt** | [**time.Time**](time.Time.md) | When the last comment was updated | [optional] [default to null]
**Project** | **int32** | Project ID for this task | [optional] [default to null]
**CommentAuthors** | **[]int32** | Users who wrote comments | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


