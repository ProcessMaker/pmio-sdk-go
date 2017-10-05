# \Gateways

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGateway**](Gateways.md#AddGateway) | **Post** /processes/{process_id}/gateways | 
[**DeleteGateway**](Gateways.md#DeleteGateway) | **Delete** /processes/{process_id}/gateways/{gateway_id} | 
[**FindGatewayById**](Gateways.md#FindGatewayById) | **Get** /processes/{process_id}/gateways/{gateway_id} | 
[**FindGateways**](Gateways.md#FindGateways) | **Get** /processes/{process_id}/gateways | 
[**UpdateGateway**](Gateways.md#UpdateGateway) | **Put** /processes/{process_id}/gateways/{gateway_id} | 


# **AddGateway**
> GatewayItem AddGateway($processId, $gatewayCreateItem)



This method creates a new gateway.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process related to the gateway | 
 **gatewayCreateItem** | [**GatewayCreateItem**](GatewayCreateItem.md)| JSON API response with the gateway object to add | 

### Return type

[**GatewayItem**](GatewayItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGateway**
> ResultSuccess DeleteGateway($processId, $gatewayId)



This method deletes a single item using the gateway ID and the process ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **gatewayId** | **string**| ID of the process to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindGatewayById**
> GatewayItem FindGatewayById($processId, $gatewayId)



This method retrieves a gateway based on its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to return | 
 **gatewayId** | **string**| ID of gateway to return | 

### Return type

[**GatewayItem**](GatewayItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindGateways**
> GatewayCollection FindGateways($processId, $page, $perPage)



This method retrieves all existing gateways.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process related to the gateway | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**GatewayCollection**](GatewayCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGateway**
> GatewayItem UpdateGateway($processId, $gatewayId, $gatewayUpdateItem)



This method updates an existing gateway.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to retrieve | 
 **gatewayId** | **string**| ID of the gateway to retrieve | 
 **gatewayUpdateItem** | [**GatewayUpdateItem**](GatewayUpdateItem.md)| Gateway object to edit | 

### Return type

[**GatewayItem**](GatewayItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

