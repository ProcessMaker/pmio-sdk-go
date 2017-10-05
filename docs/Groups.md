# \Groups

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroup**](Groups.md#AddGroup) | **Post** /groups | 
[**AddUsersToGroup**](Groups.md#AddUsersToGroup) | **Put** /groups/{id}/users | 
[**DeleteGroup**](Groups.md#DeleteGroup) | **Delete** /groups/{id} | 
[**FindGroupById**](Groups.md#FindGroupById) | **Get** /groups/{id} | 
[**FindGroups**](Groups.md#FindGroups) | **Get** /groups | 
[**RemoveUsersFromGroup**](Groups.md#RemoveUsersFromGroup) | **Delete** /groups/{id}/users | 
[**SyncUsersToGroup**](Groups.md#SyncUsersToGroup) | **Post** /groups/{id}/users | 
[**UpdateGroup**](Groups.md#UpdateGroup) | **Put** /groups/{id} | 


# **AddGroup**
> GroupItem AddGroup($groupCreateItem)



This method creates a new group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupCreateItem** | [**GroupCreateItem**](GroupCreateItem.md)| JSON API with the Group object to add | 

### Return type

[**GroupItem**](GroupItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddUsersToGroup**
> ResultSuccess AddUsersToGroup($id, $groupAddUsersItem)



This method adds one or more new users to a group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of group to be modified | 
 **groupAddUsersItem** | [**GroupAddUsersItem**](GroupAddUsersItem.md)| JSON API response with array of user IDs | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroup**
> ResultSuccess DeleteGroup($id)



This method deletes a group using the group ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of group to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindGroupById**
> GroupItem FindGroupById($id)



This method retrieves a group using its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of group to return | 

### Return type

[**GroupItem**](GroupItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindGroups**
> GroupCollection FindGroups($page, $perPage)



This method retrieves all existing groups.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**GroupCollection**](GroupCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveUsersFromGroup**
> ResultSuccess RemoveUsersFromGroup($id, $groupRemoveUsersItem)



This method removes one or more users from a group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of group to be modified | 
 **groupRemoveUsersItem** | [**GroupRemoveUsersItem**](GroupRemoveUsersItem.md)| JSON API response with Users IDs to remove | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncUsersToGroup**
> ResultSuccess SyncUsersToGroup($id, $groupSyncUsersItem)



This method synchronizes one or more users with a group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of group to be modified | 
 **groupSyncUsersItem** | [**GroupSyncUsersItem**](GroupSyncUsersItem.md)| JSON API with array of users IDs to sync | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGroup**
> GroupItem UpdateGroup($id, $groupUpdateItem)



This method updates an existing group.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of group to retrieve | 
 **groupUpdateItem** | [**GroupUpdateItem**](GroupUpdateItem.md)| Group object to edit | 

### Return type

[**GroupItem**](GroupItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

