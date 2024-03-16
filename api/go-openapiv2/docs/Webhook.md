# Webhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Organization** | **int32** |  | [optional] [default to null]
**Project** | **int32** |  | [optional] [default to null]
**Url** | **string** | URL of webhook | [default to null]
**SendPayload** | **bool** | If value is False send only action | [optional] [default to null]
**SendForAllActions** | **bool** | If value is False - used only for actions from WebhookAction | [optional] [default to null]
**Headers** | **interface{}** | Key Value Json of headers | [optional] [default to null]
**IsActive** | **bool** | If value is False the webhook is disabled | [optional] [default to null]
**Actions** | **[]string** |  | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time | [optional] [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | Last update time | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


