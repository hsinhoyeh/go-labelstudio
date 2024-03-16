# S3ImportStorage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Type_** | **string** |  | [optional] [default to null]
**Synchronizable** | **bool** |  | [optional] [default to null]
**Presign** | **bool** |  | [optional] [default to null]
**LastSync** | [**time.Time**](time.Time.md) | Last sync finished time | [optional] [default to null]
**LastSyncCount** | **int32** | Count of tasks synced last time | [optional] [default to null]
**LastSyncJob** | **string** | Last sync job ID | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Traceback** | **string** | Traceback report for the last failed sync | [optional] [default to null]
**Meta** | **interface{}** | Meta and debug information about storage processes | [optional] [default to null]
**Title** | **string** | Cloud storage title | [optional] [default to null]
**Description** | **string** | Cloud storage description | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation time | [optional] [default to null]
**Bucket** | **string** | S3 bucket name | [optional] [default to null]
**Prefix** | **string** | S3 bucket prefix | [optional] [default to null]
**RegexFilter** | **string** | Cloud storage regex for filtering objects | [optional] [default to null]
**UseBlobUrls** | **bool** | Interpret objects as BLOBs and generate URLs | [optional] [default to null]
**AwsAccessKeyId** | **string** | AWS_ACCESS_KEY_ID | [optional] [default to null]
**AwsSecretAccessKey** | **string** | AWS_SECRET_ACCESS_KEY | [optional] [default to null]
**AwsSessionToken** | **string** | AWS_SESSION_TOKEN | [optional] [default to null]
**AwsSseKmsKeyId** | **string** | AWS SSE KMS Key ID | [optional] [default to null]
**RegionName** | **string** | AWS Region | [optional] [default to null]
**S3Endpoint** | **string** | S3 Endpoint | [optional] [default to null]
**PresignTtl** | **int32** | Presigned URLs TTL (in minutes) | [optional] [default to null]
**RecursiveScan** | **bool** | Perform recursive scan over the bucket content | [optional] [default to null]
**Project** | **int32** | A unique integer value identifying this project. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


