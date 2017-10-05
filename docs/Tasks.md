# \Tasks

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupsToTask**](Tasks.md#AddGroupsToTask) | **Put** /processes/{process_id}/tasks/{task_id}/groups | 
[**AddTask**](Tasks.md#AddTask) | **Post** /processes/{process_id}/tasks | 
[**AddTaskConnector**](Tasks.md#AddTaskConnector) | **Post** /processes/{process_id}/tasks/{task_id}/connectors | 
[**DeleteTask**](Tasks.md#DeleteTask) | **Delete** /processes/{process_id}/tasks/{task_id} | 
[**DeleteTaskConnector**](Tasks.md#DeleteTaskConnector) | **Delete** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
[**FindTaskById**](Tasks.md#FindTaskById) | **Get** /processes/{process_id}/tasks/{task_id} | 
[**FindTaskConnectorById**](Tasks.md#FindTaskConnectorById) | **Get** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
[**FindTaskConnectors**](Tasks.md#FindTaskConnectors) | **Get** /processes/{process_id}/tasks/{task_id}/connectors | 
[**FindTaskInstanceById**](Tasks.md#FindTaskInstanceById) | **Get** /task_instances/{task_instance_id} | 
[**FindTaskInstances**](Tasks.md#FindTaskInstances) | **Get** /task_instances | 
[**FindTasks**](Tasks.md#FindTasks) | **Get** /processes/{process_id}/tasks | 
[**RemoveGroupsFromTask**](Tasks.md#RemoveGroupsFromTask) | **Delete** /processes/{process_id}/tasks/{task_id}/groups | 
[**SyncGroupsToTask**](Tasks.md#SyncGroupsToTask) | **Post** /processes/{process_id}/tasks/{task_id}/groups | 
[**UpdateTask**](Tasks.md#UpdateTask) | **Put** /processes/{process_id}/tasks/{task_id} | 
[**UpdateTaskConnector**](Tasks.md#UpdateTaskConnector) | **Put** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
[**UpdateTaskInstance**](Tasks.md#UpdateTaskInstance) | **Patch** /task_instances/{task_instance_id} | 


# **AddGroupsToTask**
> ResultSuccess AddGroupsToTask($processId, $taskId, $taskAddGroupsItem)



This method assigns group(s) to the chosen task


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **taskId** | **string**| ID of the task to be modified | 
 **taskAddGroupsItem** | [**TaskAddGroupsItem**](TaskAddGroupsItem.md)| JSON API with Group IDs to add | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTask**
> TaskItem AddTask($processId, $taskCreateItem)



This method creates a new task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to the task | 
 **taskCreateItem** | [**TaskCreateItem**](TaskCreateItem.md)| JSON API with the task object to add | 

### Return type

[**TaskItem**](TaskItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTaskConnector**
> TaskConnector1 AddTaskConnector($processId, $taskId, $taskConnectorCreateItem)



This method is for creating a new task connector


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **taskId** | **string**| ID of the task to fetch | 
 **taskConnectorCreateItem** | [**TaskConnectorCreateItem**](TaskConnectorCreateItem.md)| JSON API with the TaskConnector object to add | 

### Return type

[**TaskConnector1**](TaskConnector_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTask**
> ResultSuccess DeleteTask($processId, $taskId)



This method deletes a task using the task ID and the process ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **taskId** | **string**| ID of a task to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTaskConnector**
> ResultSuccess DeleteTaskConnector($processId, $taskId, $connectorId)



This method is for deleting a single task connector based on task ID, the process ID and the Connector ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process item to fetch | 
 **taskId** | **string**| ID of the task item to fetch | 
 **connectorId** | **string**| ID of TaskConnector to fetch | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTaskById**
> TaskItem FindTaskById($processId, $taskId)



This method retrieves a task using its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to return | 
 **taskId** | **string**| ID of the task to return | 

### Return type

[**TaskItem**](TaskItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTaskConnectorById**
> TaskConnector1 FindTaskConnectorById($processId, $taskId, $connectorId)



This method is for retrieving a task connector based on its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **taskId** | **string**| ID of the task to fetch | 
 **connectorId** | **string**| ID of TaskConnector to fetch | 

### Return type

[**TaskConnector1**](TaskConnector_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTaskConnectors**
> TaskConnectorsCollection FindTaskConnectors($processId, $taskId, $page, $perPage)



This method returns all task connectors related to the run process and task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **taskId** | **string**| ID of the task to fetch | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**TaskConnectorsCollection**](TaskConnectorsCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTaskInstanceById**
> InlineResponse200 FindTaskInstanceById($taskInstanceId, $page, $perPage)



This method retrieves a task instance based on its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskInstanceId** | **string**| ID of task instance to return | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTaskInstances**
> TaskInstanceCollection FindTaskInstances($page, $perPage)



This method retrieves all existing task instances.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**TaskInstanceCollection**](TaskInstanceCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTasks**
> TaskCollection FindTasks($processId, $page, $perPage)



This method is intended for returning a list of all tasks related to the process.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process relative to the task | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**TaskCollection**](TaskCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveGroupsFromTask**
> ResultSuccess RemoveGroupsFromTask($processId, $taskId, $taskRemoveGroupsItem)



This method removes groups from a task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **taskId** | **string**| Task ID | 
 **taskRemoveGroupsItem** | [**TaskRemoveGroupsItem**](TaskRemoveGroupsItem.md)| JSON API response with Group IDs to remove | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncGroupsToTask**
> ResultSuccess SyncGroupsToTask($processId, $taskId, $taskSyncGroupsItem)



This method synchronizes one or more groups with a task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **taskId** | **string**| ID of the task to modify | 
 **taskSyncGroupsItem** | [**TaskSyncGroupsItem**](TaskSyncGroupsItem.md)| JSON API response with group IDs to sync | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTask**
> TaskItem UpdateTask($processId, $taskId, $taskUpdateItem)



This method is for updating an existing task.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **taskId** | **string**| ID of the task to fetch | 
 **taskUpdateItem** | [**TaskUpdateItem**](TaskUpdateItem.md)| Task object to edit | 

### Return type

[**TaskItem**](TaskItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaskConnector**
> TaskConnector1 UpdateTaskConnector($processId, $taskId, $connectorId, $taskConnectorUpdateItem)



This method updates the existing task connector with new parameter values.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **taskId** | **string**| ID of the task to fetch | 
 **connectorId** | **string**| ID of the task connector to fetch | 
 **taskConnectorUpdateItem** | [**TaskConnectorUpdateItem**](TaskConnectorUpdateItem.md)| TaskConnector object to edit | 

### Return type

[**TaskConnector1**](TaskConnector_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTaskInstance**
> InlineResponse200 UpdateTaskInstance($taskInstanceId, $taskInstanceUpdateItem)



This method updates an existing task instance.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskInstanceId** | **string**| ID of the task instance to retrieve | 
 **taskInstanceUpdateItem** | [**TaskInstanceUpdateItem**](TaskInstanceUpdateItem.md)| Task instance object to update | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

