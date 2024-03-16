# TaskFilterOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**View** | **int32** | Apply filters from the view ID (a tab from the Data Manager) | [optional] [default to null]
**Skipped** | **string** | &#x60;only&#x60; - include all tasks with skipped annotations&lt;br&gt;&#x60;exclude&#x60; - exclude all tasks with skipped annotations | [optional] [default to null]
**Finished** | **string** | &#x60;only&#x60; - include all finished tasks (is_labeled &#x3D; true)&lt;br&gt;&#x60;exclude&#x60; - exclude all finished tasks | [optional] [default to null]
**Annotated** | **string** | &#x60;only&#x60; - include all tasks with at least one not skipped annotation&lt;br&gt;&#x60;exclude&#x60; - exclude all tasks with at least one not skipped annotation | [optional] [default to null]
**OnlyWithAnnotations** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


