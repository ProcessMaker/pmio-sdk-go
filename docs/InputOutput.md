# \Inputoutput

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInputOutput**](Inputoutput.md#AddInputOutput) | **Post** /processes/{process_id}/tasks/{task_id}/inputoutput | 
[**DeleteInputOutput**](Inputoutput.md#DeleteInputOutput) | **Delete** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
[**FindInputOutputById**](Inputoutput.md#FindInputOutputById) | **Get** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
[**FindInputOutputs**](Inputoutput.md#FindInputOutputs) | **Get** /processes/{process_id}/tasks/{task_id}/inputoutput | 
[**UpdateInputOutput**](Inputoutput.md#UpdateInputOutput) | **Put** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 


# **AddInputOutput**
> InputOutputItem AddInputOutput($processId, $taskId, $inputOutputCreateItem)



This method creates a new Input/Output object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to Input/Output object | 
 **taskId** | **string**| Task instance ID related to Input/Output object | 
 **inputOutputCreateItem** | [**InputOutputCreateItem**](InputOutputCreateItem.md)| Create and add a new Input/Output object with JSON API | 

### Return type

[**InputOutputItem**](InputOutputItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInputOutput**
> ResultSuccess DeleteInputOutput($processId, $taskId, $inputoutputUid)



This method deletes the Input/Output based on the Input/Output ID, process ID and task ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to the Input/Output object | 
 **taskId** | **string**| Task instance ID related to Input/Output object | 
 **inputoutputUid** | **string**| Input/Output ID to fetch | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindInputOutputById**
> InputOutputItem FindInputOutputById($processId, $taskId, $inputoutputUid)



This method retrieves an Input/Output object using its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to the Input/Output object | 
 **taskId** | **string**| Task instance ID related to the Input/Output object | 
 **inputoutputUid** | **string**| ID of Input/Output to return | 

### Return type

[**InputOutputItem**](InputOutputItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindInputOutputs**
> InputOutputCollection FindInputOutputs($processId, $taskId, $page, $perPage)



This method retrieves all existing Input/Output objects in the related task instance.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to Input/Output object | 
 **taskId** | **string**| Task instance ID related to Input/Output object | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**InputOutputCollection**](InputOutputCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInputOutput**
> InputOutputItem UpdateInputOutput($processId, $taskId, $inputoutputUid, $inputOutputUpdateItem)



This method updates an existing Input/Output object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to the Input/Output object | 
 **taskId** | **string**| Task instance ID related to the Input/Output object | 
 **inputoutputUid** | **string**| ID of Input/Output to retrieve | 
 **inputOutputUpdateItem** | [**InputOutputUpdateItem**](InputOutputUpdateItem.md)| Input/Output object to edit | 

### Return type

[**InputOutputItem**](InputOutputItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

