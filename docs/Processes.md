# \Processes

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProcess**](Processes.md#AddProcess) | **Post** /processes | 
[**CallImport**](Processes.md#CallImport) | **Post** /processes/import/bpmn | 
[**DeleteProcess**](Processes.md#DeleteProcess) | **Delete** /processes/{id} | 
[**FindProcessById**](Processes.md#FindProcessById) | **Get** /processes/{id} | 
[**FindProcesses**](Processes.md#FindProcesses) | **Get** /processes | 
[**ImportBpmnFile**](Processes.md#ImportBpmnFile) | **Post** /processes/import | 
[**UpdateProcess**](Processes.md#UpdateProcess) | **Put** /processes/{id} | 


# **AddProcess**
> ProcessItem AddProcess($processCreateItem)



This method creates a new process.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processCreateItem** | [**ProcessCreateItem**](ProcessCreateItem.md)| JSON API response with the process object to add | 

### Return type

[**ProcessItem**](ProcessItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CallImport**
> ProcessCollection1 CallImport($importItem)



This method imports BPMN 2.0 files. A new process(es) is/are created and its object returned back when import is successful.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importItem** | [**ImportItem**](ImportItem.md)| JSON API with the BPMN file to import | 

### Return type

[**ProcessCollection1**](ProcessCollection_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProcess**
> ResultSuccess DeleteProcess($id)



This method deletes a process using the process ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| Process ID to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProcessById**
> ProcessItem FindProcessById($id)



This method retrieves a process using its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the process to return | 

### Return type

[**ProcessItem**](ProcessItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindProcesses**
> ProcessCollection FindProcesses($page, $perPage)



This method retrieves all existing processes.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**ProcessCollection**](ProcessCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportBpmnFile**
> ProcessCollection1 ImportBpmnFile($bpmnImportItem)



This method imports BPMN 2.0 files. A new process(es) is/are created and its object returned back when import is successful.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bpmnImportItem** | [**BpmnImportItem**](BpmnImportItem.md)| JSON API with the BPMN file to import | 

### Return type

[**ProcessCollection1**](ProcessCollection_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProcess**
> ProcessItem UpdateProcess($id, $processUpdateItem)



This method updates an existing process.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| ID of the process to retrieve | 
 **processUpdateItem** | [**ProcessUpdateItem**](ProcessUpdateItem.md)| Process object to edit | 

### Return type

[**ProcessItem**](ProcessItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

