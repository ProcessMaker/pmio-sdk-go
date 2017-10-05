# \Oauth

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOauthClient**](Oauth.md#AddOauthClient) | **Post** /users/{user_id}/clients | 
[**FindOauthClientById**](Oauth.md#FindOauthClientById) | **Get** /users/{user_id}/clients/{client_id} | 
[**FindOauthClients**](Oauth.md#FindOauthClients) | **Get** /users/{user_id}/clients | 


# **AddOauthClient**
> OauthClientItem AddOauthClient($userId, $oauthClientCreateItem)



This method creates a new Oauth client for the user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| ID of the user related to the Oauth client | 
 **oauthClientCreateItem** | [**OauthClientCreateItem**](OauthClientCreateItem.md)| JSON API with the Oauth Client object to add | 

### Return type

[**OauthClientItem**](OauthClientItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOauthClientById**
> OauthClientItem FindOauthClientById($userId, $clientId)



This method retrieves an Oauth client for the User based on its ID.  The response contains the `client_secret` required to obtain the `access_token`.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| ID of user to retrieve | 
 **clientId** | **string**| ID of Oauth client to retrieve | 

### Return type

[**OauthClientItem**](OauthClientItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindOauthClients**
> OauthClientCollection FindOauthClients($userId, $page, $perPage)



This method retrieves all existing Oauth clients belonging to a user.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| User ID related to the Oauth clients | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**OauthClientCollection**](OauthClientCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

