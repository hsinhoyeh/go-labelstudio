# MlBackend

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**State** | **string** |  | [optional] [default to null]
**IsInteractive** | **bool** | Used to interactively annotate tasks. If true, model returns one list with results | [optional] [default to null]
**Url** | **string** | URL for the machine learning model server | [default to null]
**ErrorMessage** | **string** | Error message in error state | [optional] [default to null]
**Title** | **string** | Name of the machine learning backend | [optional] [default to null]
**Description** | **string** | Description for the machine learning backend | [optional] [default to null]
**ModelVersion** | **string** | Current model version associated with this machine learning backend | [optional] [default to null]
**Timeout** | **float32** | Response model timeout | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**AutoUpdate** | **bool** | If false, model version is set by the user, if true - getting latest version from backend. | [optional] [default to null]
**Project** | **int32** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


