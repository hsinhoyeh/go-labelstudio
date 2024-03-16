# TaskCreationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskCount** | **int32** | Number of tasks added | [optional] [default to null]
**AnnotationCount** | **int32** | Number of annotations added | [optional] [default to null]
**PredictionsCount** | **int32** | Number of predictions added | [optional] [default to null]
**Duration** | **float32** | Time in seconds to create | [optional] [default to null]
**FileUploadIds** | **[]int32** | Database IDs of uploaded files | [optional] [default to null]
**CouldBeTasksList** | **bool** | Whether uploaded files can contain lists of tasks, like CSV/TSV files | [optional] [default to null]
**FoundFormats** | **[]string** | The list of found file formats | [optional] [default to null]
**DataColumns** | **[]string** | The list of found data columns | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


