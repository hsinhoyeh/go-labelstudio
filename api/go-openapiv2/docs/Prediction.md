# Prediction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**ModelVersion** | **string** |  | [optional] [default to null]
**CreatedAgo** | **string** | Delta time from creation time | [optional] [default to null]
**Result** | **interface{}** | Prediction result | [optional] [default to null]
**Score** | **float32** | Prediction score | [optional] [default to null]
**Cluster** | **int32** | Cluster for the current prediction | [optional] [default to null]
**Neighbors** | **interface{}** | Array of task IDs of the closest neighbors | [optional] [default to null]
**Mislabeling** | **float32** | Related task mislabeling score | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Task** | **int32** |  | [default to null]
**Project** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


