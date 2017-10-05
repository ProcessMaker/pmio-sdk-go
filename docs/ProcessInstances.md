# \ProcessInstances

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInstance**](ProcessInstances.md#AddInstance) | **Post** /processes/{process_id}/instances | 
[**DeleteInstance**](ProcessInstances.md#DeleteInstance) | **Delete** /processes/{process_id}/instances/{instance_id} | 
[**FindByFieldInsideDataModel**](ProcessInstances.md#FindByFieldInsideDataModel) | **Get** /processes/{process_id}/datamodels/search/{search_param} | 
[**FindDataModel**](ProcessInstances.md#FindDataModel) | **Get** /processes/{process_id}/instances/{instance_id}/datamodel | 
[**FindInstanceById**](ProcessInstances.md#FindInstanceById) | **Get** /processes/{process_id}/instances/{instance_id} | 
[**FindInstances**](ProcessInstances.md#FindInstances) | **Get** /processes/{process_id}/instances | 
[**FindTaskInstancesByInstanceAndTaskId**](ProcessInstances.md#FindTaskInstancesByInstanceAndTaskId) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances | 
[**FindTaskInstancesByInstanceAndTaskIdDelegated**](ProcessInstances.md#FindTaskInstancesByInstanceAndTaskIdDelegated) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances/delegated | 
[**FindTaskInstancesByInstanceAndTaskIdStarted**](ProcessInstances.md#FindTaskInstancesByInstanceAndTaskIdStarted) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances/started | 
[**FindTokens**](ProcessInstances.md#FindTokens) | **Get** /processes/{process_id}/instances/{instance_id}/tokens | 
[**UpdateInstance**](ProcessInstances.md#UpdateInstance) | **Put** /processes/{process_id}/instances/{instance_id} | 


# **AddInstance**
> InstanceItem AddInstance($processId, $instanceCreateItem)



This method creates a new instance.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to the instance | 
 **instanceCreateItem** | [**InstanceCreateItem**](InstanceCreateItem.md)| JSON API response with the instance object to add | 

### Return type

[**InstanceItem**](InstanceItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInstance**
> ResultSuccess DeleteInstance($processId, $instanceId)



This method deletes an instance using the instance ID and the process ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **instanceId** | **string**| ID of the instance to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindByFieldInsideDataModel**
> DataModelCollection FindByFieldInsideDataModel($processId, $searchParam, $page, $perPage)



This method returns the data model by field passed in get argument.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to return | 
 **searchParam** | **string**| Key and value of searched field in DataModel | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**DataModelCollection**](DataModelCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindDataModel**
> DataModelItem1 FindDataModel($processId, $instanceId, $page, $perPage)



This method returns the instance data model and lets the user work with it directly.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to return | 
 **instanceId** | **string**| ID of the instance to return | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**DataModelItem1**](DataModelItem_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindInstanceById**
> InstanceItem FindInstanceById($processId, $instanceId)



This method retrieves an instance using its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to return | 
 **instanceId** | **string**| ID of the instance to return | 

### Return type

[**InstanceItem**](InstanceItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindInstances**
> InstanceCollection FindInstances($processId, $page, $perPage)



This method retrieves instances related to the process using the process ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to the instances | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**InstanceCollection**](InstanceCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTaskInstancesByInstanceAndTaskId**
> TaskInstanceCollection FindTaskInstancesByInstanceAndTaskId($instanceId, $taskId)



This method retrieves task instances using the instance ID and the task ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **string**| ID of the instance | 
 **taskId** | **string**| ID of the task | 

### Return type

[**TaskInstanceCollection**](TaskInstanceCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTaskInstancesByInstanceAndTaskIdDelegated**
> TaskInstanceCollection FindTaskInstancesByInstanceAndTaskIdDelegated($instanceId, $taskId)



This method retrieves delegated task instances using the instance ID and the task ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **string**| ID of the instance | 
 **taskId** | **string**| ID of the task | 

### Return type

[**TaskInstanceCollection**](TaskInstanceCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTaskInstancesByInstanceAndTaskIdStarted**
> TaskInstanceCollection FindTaskInstancesByInstanceAndTaskIdStarted($instanceId, $taskId)



This method retrieves started task instances using the instance ID and the task ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **string**| ID of the instance | 
 **taskId** | **string**| ID of the task | 

### Return type

[**TaskInstanceCollection**](TaskInstanceCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTokens**
> TokenCollection FindTokens($processId, $instanceId, $page, $perPage)



This method retrieves tokens related to the process and instance using the process and instance IDs


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **instanceId** | **string**| Instance ID related to the process | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**TokenCollection**](TokenCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInstance**
> InstanceItem UpdateInstance($processId, $instanceId, $instanceUpdateItem)



This method updates  an existing instance.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to retrieve | 
 **instanceId** | **string**| ID of the instance to retrieve | 
 **instanceUpdateItem** | [**InstanceUpdateItem**](InstanceUpdateItem.md)| Instance object to edit | 

### Return type

[**InstanceItem**](InstanceItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

