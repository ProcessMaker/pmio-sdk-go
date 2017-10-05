# \Flows

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFlow**](Flows.md#AddFlow) | **Post** /processes/{process_id}/flows | 
[**DeleteFlow**](Flows.md#DeleteFlow) | **Delete** /processes/{process_id}/flows/{flow_id} | 
[**FindFlowById**](Flows.md#FindFlowById) | **Get** /processes/{process_id}/flows/{flow_id} | 
[**FindFlows**](Flows.md#FindFlows) | **Get** /processes/{process_id}/flows | 
[**UpdateFlow**](Flows.md#UpdateFlow) | **Put** /processes/{process_id}/flows/{flow_id} | 


# **AddFlow**
> FlowItem AddFlow($processId, $flowCreateItem)



This method creates a new Sequence Flow.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process related to the flow | 
 **flowCreateItem** | [**FlowCreateItem**](FlowCreateItem.md)| JSON API response with the Flow object to add | 

### Return type

[**FlowItem**](FlowItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFlow**
> ResultSuccess DeleteFlow($processId, $flowId)



This method deletes the Sequence Flow using the flow ID and the process ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **flowId** | **string**| ID of the flow to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindFlowById**
> FlowItem FindFlowById($processId, $flowId)



This method retrieves a flow based on its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to return | 
 **flowId** | **string**| ID of the flow to return | 

### Return type

[**FlowItem**](FlowItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindFlows**
> FlowCollection FindFlows($processId, $page, $perPage)



This method retrieves all existing flows.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process related to the flow | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**FlowCollection**](FlowCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFlow**
> FlowItem UpdateFlow($processId, $flowId, $flowUpdateItem)



This method updates an existing flow.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to retrieve | 
 **flowId** | **string**| ID of the flow to retrieve | 
 **flowUpdateItem** | [**FlowUpdateItem**](FlowUpdateItem.md)| Flow object to edit | 

### Return type

[**FlowItem**](FlowItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

