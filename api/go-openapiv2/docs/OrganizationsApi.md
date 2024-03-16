# \OrganizationsApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiOrganizationsList**](OrganizationsApi.md#ApiOrganizationsList) | **Get** /api/organizations/ | List your organizations
[**ApiOrganizationsMembershipsDelete**](OrganizationsApi.md#ApiOrganizationsMembershipsDelete) | **Delete** /api/organizations/{id}/memberships/{user_pk}/ | Soft delete an organization member
[**ApiOrganizationsMembershipsList**](OrganizationsApi.md#ApiOrganizationsMembershipsList) | **Get** /api/organizations/{id}/memberships | Get organization members list
[**ApiOrganizationsPartialUpdate**](OrganizationsApi.md#ApiOrganizationsPartialUpdate) | **Patch** /api/organizations/{id} | Update organization settings
[**ApiOrganizationsRead**](OrganizationsApi.md#ApiOrganizationsRead) | **Get** /api/organizations/{id} | Get organization settings


# **ApiOrganizationsList**
> []OrganizationId ApiOrganizationsList(ctx, )
List your organizations

 Return a list of the organizations you've created or that you have access to. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OrganizationId**](OrganizationId.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrganizationsMembershipsDelete**
> ApiOrganizationsMembershipsDelete(ctx, id, pk, userPk)
Soft delete an organization member

Soft delete a member from the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **pk** | **int32**| A unique integer value identifying this organization. | 
  **userPk** | **int32**| A unique integer value identifying the user to be deleted from the organization. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrganizationsMembershipsList**
> InlineResponse2002 ApiOrganizationsMembershipsList(ctx, id, optional)
Get organization members list

Retrieve a list of the organization members and their IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this organization. | 
 **optional** | ***OrganizationsApiApiOrganizationsMembershipsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OrganizationsApiApiOrganizationsMembershipsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| A page number within the paginated result set. | 
 **pageSize** | **optional.Int32**| Number of results to return per page. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrganizationsPartialUpdate**
> Organization ApiOrganizationsPartialUpdate(ctx, id, data)
Update organization settings

Update the settings for a specific organization by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this organization. | 
  **data** | [**Organization**](Organization.md)|  | 

### Return type

[**Organization**](Organization.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiOrganizationsRead**
> Organization ApiOrganizationsRead(ctx, id)
Get organization settings

Retrieve the settings for a specific organization by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this organization. | 

### Return type

[**Organization**](Organization.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

