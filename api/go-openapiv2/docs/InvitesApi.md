# \InvitesApi

All URIs are relative to *http://api.labelstud.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiInviteRead**](InvitesApi.md#ApiInviteRead) | **Get** /api/invite | Get organization invite link
[**ApiInviteResetTokenCreate**](InvitesApi.md#ApiInviteResetTokenCreate) | **Post** /api/invite/reset-token | Reset organization token


# **ApiInviteRead**
> OrganizationInvite ApiInviteRead(ctx, )
Get organization invite link

Get a link to use to invite a new member to an organization in Label Studio Enterprise.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrganizationInvite**](OrganizationInvite.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiInviteResetTokenCreate**
> OrganizationInvite ApiInviteResetTokenCreate(ctx, )
Reset organization token

Reset the token used in the invitation link to invite someone to an organization.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OrganizationInvite**](OrganizationInvite.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

