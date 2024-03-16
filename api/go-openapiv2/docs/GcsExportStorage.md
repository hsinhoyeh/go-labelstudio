# GcsExportStorage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Synchronizable** | **bool** |  | [optional] [default to null]
**Bucket** | **string** | GCS bucket name | [optional] [default to null]
**Prefix** | **string** | GCS bucket prefix | [optional] [default to null]
**RegexFilter** | **string** | Cloud storage regex for filtering objects | [optional] [default to null]
**UseBlobUrls** | **bool** | Interpret objects as BLOBs and generate URLs | [optional] [default to null]
**GoogleApplicationCredentials** | **string** | The content of GOOGLE_APPLICATION_CREDENTIALS json file | [optional] [default to null]
**GoogleProjectId** | **string** | Google project ID | [optional] [default to null]
**LastSync** | [**time.Time**](time.Time.md) | Last sync finished time | [optional] [default to null]
**LastSyncCount** | **int32** | Count of tasks synced last time | [optional] [default to null]
**LastSyncJob** | **string** | Last sync job ID | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Traceback** | **string** | Traceback report for the last failed sync | [optional] [default to null]
**Meta** | **interface{}** | Meta and debug information about storage processes | [optional] [default to null]
**Title** | **string** | Cloud storage title | [optional] [default to null]
**Description** | **string** | Cloud storage description | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time | [optional] [default to null]
**CanDeleteObjects** | **bool** | Deletion from storage enabled | [optional] [default to null]
**Project** | **int32** | A unique integer value identifying this project. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


