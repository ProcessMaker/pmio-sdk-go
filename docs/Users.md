# \Users

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](Users.md#AddUser) | **Post** /users | 
[**DeleteUser**](Users.md#DeleteUser) | **Delete** /users/{id} | 
[**FindUserById**](Users.md#FindUserById) | **Get** /users/{id} | 
[**FindUsers**](Users.md#FindUsers) | **Get** /users | 
[**MyselfUser**](Users.md#MyselfUser) | **Get** /users/myself | 
[**UpdateUser**](Users.md#UpdateUser) | **Put** /users/{id} | 


# **AddUser**
> UserItem AddUser($userCreateItem)



This method creates a new user in the system. The client_id will appear in the results.  The `client_id` is required to obtain a `client_secret` and then you will be able to use it in an Oauth authorization key. Refer to [Oauth Client APIs](#tag/oauth)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userCreateItem** | [**UserCreateItem**](UserCreateItem.md)| JSON API with the User object to add | 

### Return type

[**UserItem**](UserItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUser**
> ResultSuccess DeleteUser($id)



This method deletes a user from the system.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of user to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUserById**
> UserItem FindUserById($id)



This method returns a user using its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the user to return | 

### Return type

[**UserItem**](UserItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindUsers**
> UserCollection FindUsers($page, $perPage)



This method returns all existing users in the system.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**UserCollection**](UserCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MyselfUser**
> UserItem MyselfUser($page, $perPage)



This method returns user information using a token.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**UserItem**](UserItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> UserItem UpdateUser($id, $userUpdateItem)



This method updates an existing user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of user to retrieve | 
 **userUpdateItem** | [**UserUpdateItem**](UserUpdateItem.md)| User object for update | 

### Return type

[**UserItem**](UserItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

