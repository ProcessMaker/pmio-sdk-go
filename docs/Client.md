# \Client

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEvent**](Client.md#AddEvent) | **Post** /processes/{process_id}/events | 
[**AddEventConnector**](Client.md#AddEventConnector) | **Post** /processes/{process_id}/events/{event_id}/connectors | 
[**AddFlow**](Client.md#AddFlow) | **Post** /processes/{process_id}/flows | 
[**AddGateway**](Client.md#AddGateway) | **Post** /processes/{process_id}/gateways | 
[**AddGroup**](Client.md#AddGroup) | **Post** /groups | 
[**AddGroupsToTask**](Client.md#AddGroupsToTask) | **Put** /processes/{process_id}/tasks/{task_id}/groups | 
[**AddInputOutput**](Client.md#AddInputOutput) | **Post** /processes/{process_id}/tasks/{task_id}/inputoutput | 
[**AddInstance**](Client.md#AddInstance) | **Post** /processes/{process_id}/instances | 
[**AddOauthClient**](Client.md#AddOauthClient) | **Post** /users/{user_id}/clients | 
[**AddProcess**](Client.md#AddProcess) | **Post** /processes | 
[**AddTask**](Client.md#AddTask) | **Post** /processes/{process_id}/tasks | 
[**AddTaskConnector**](Client.md#AddTaskConnector) | **Post** /processes/{process_id}/tasks/{task_id}/connectors | 
[**AddUser**](Client.md#AddUser) | **Post** /users | 
[**AddUsersToGroup**](Client.md#AddUsersToGroup) | **Put** /groups/{id}/users | 
[**CallImport**](Client.md#CallImport) | **Post** /processes/import/bpmn | 
[**DeleteEvent**](Client.md#DeleteEvent) | **Delete** /processes/{process_id}/events/{event_id} | 
[**DeleteEventConnector**](Client.md#DeleteEventConnector) | **Delete** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
[**DeleteFlow**](Client.md#DeleteFlow) | **Delete** /processes/{process_id}/flows/{flow_id} | 
[**DeleteGateway**](Client.md#DeleteGateway) | **Delete** /processes/{process_id}/gateways/{gateway_id} | 
[**DeleteGroup**](Client.md#DeleteGroup) | **Delete** /groups/{id} | 
[**DeleteInputOutput**](Client.md#DeleteInputOutput) | **Delete** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
[**DeleteInstance**](Client.md#DeleteInstance) | **Delete** /processes/{process_id}/instances/{instance_id} | 
[**DeleteOauthClient**](Client.md#DeleteOauthClient) | **Delete** /users/{user_id}/clients/{client_id} | 
[**DeleteProcess**](Client.md#DeleteProcess) | **Delete** /processes/{id} | 
[**DeleteTask**](Client.md#DeleteTask) | **Delete** /processes/{process_id}/tasks/{task_id} | 
[**DeleteTaskConnector**](Client.md#DeleteTaskConnector) | **Delete** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
[**DeleteUser**](Client.md#DeleteUser) | **Delete** /users/{id} | 
[**EventTrigger**](Client.md#EventTrigger) | **Post** /processes/{process_id}/events/{event_id}/trigger | 
[**EventWebhook**](Client.md#EventWebhook) | **Post** /processes/{process_id}/events/{event_id}/webhook | 
[**FindByFieldInsideDataModel**](Client.md#FindByFieldInsideDataModel) | **Get** /processes/{process_id}/datamodels/search/{search_param} | 
[**FindDataModel**](Client.md#FindDataModel) | **Get** /processes/{process_id}/instances/{instance_id}/datamodel | 
[**FindEventById**](Client.md#FindEventById) | **Get** /processes/{process_id}/events/{event_id} | 
[**FindEventConnectorById**](Client.md#FindEventConnectorById) | **Get** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
[**FindEventConnectors**](Client.md#FindEventConnectors) | **Get** /processes/{process_id}/events/{event_id}/connectors | 
[**FindEvents**](Client.md#FindEvents) | **Get** /processes/{process_id}/events | 
[**FindFlowById**](Client.md#FindFlowById) | **Get** /processes/{process_id}/flows/{flow_id} | 
[**FindFlows**](Client.md#FindFlows) | **Get** /processes/{process_id}/flows | 
[**FindGatewayById**](Client.md#FindGatewayById) | **Get** /processes/{process_id}/gateways/{gateway_id} | 
[**FindGateways**](Client.md#FindGateways) | **Get** /processes/{process_id}/gateways | 
[**FindGroupById**](Client.md#FindGroupById) | **Get** /groups/{id} | 
[**FindGroups**](Client.md#FindGroups) | **Get** /groups | 
[**FindInputOutputById**](Client.md#FindInputOutputById) | **Get** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
[**FindInputOutputs**](Client.md#FindInputOutputs) | **Get** /processes/{process_id}/tasks/{task_id}/inputoutput | 
[**FindInstanceById**](Client.md#FindInstanceById) | **Get** /processes/{process_id}/instances/{instance_id} | 
[**FindInstances**](Client.md#FindInstances) | **Get** /processes/{process_id}/instances | 
[**FindOauthClientById**](Client.md#FindOauthClientById) | **Get** /users/{user_id}/clients/{client_id} | 
[**FindOauthClients**](Client.md#FindOauthClients) | **Get** /users/{user_id}/clients | 
[**FindProcessById**](Client.md#FindProcessById) | **Get** /processes/{id} | 
[**FindProcesses**](Client.md#FindProcesses) | **Get** /processes | 
[**FindTaskById**](Client.md#FindTaskById) | **Get** /processes/{process_id}/tasks/{task_id} | 
[**FindTaskConnectorById**](Client.md#FindTaskConnectorById) | **Get** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
[**FindTaskConnectors**](Client.md#FindTaskConnectors) | **Get** /processes/{process_id}/tasks/{task_id}/connectors | 
[**FindTaskInstanceById**](Client.md#FindTaskInstanceById) | **Get** /task_instances/{task_instance_id} | 
[**FindTaskInstances**](Client.md#FindTaskInstances) | **Get** /task_instances | 
[**FindTaskInstancesByInstanceAndTaskId**](Client.md#FindTaskInstancesByInstanceAndTaskId) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances | 
[**FindTaskInstancesByInstanceAndTaskIdDelegated**](Client.md#FindTaskInstancesByInstanceAndTaskIdDelegated) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances/delegated | 
[**FindTaskInstancesByInstanceAndTaskIdStarted**](Client.md#FindTaskInstancesByInstanceAndTaskIdStarted) | **Get** /instances/{instance_id}/tasks/{task_id}/task_instances/started | 
[**FindTasks**](Client.md#FindTasks) | **Get** /processes/{process_id}/tasks | 
[**FindTokens**](Client.md#FindTokens) | **Get** /processes/{process_id}/instances/{instance_id}/tokens | 
[**FindUserById**](Client.md#FindUserById) | **Get** /users/{id} | 
[**FindUsers**](Client.md#FindUsers) | **Get** /users | 
[**ImportBpmnFile**](Client.md#ImportBpmnFile) | **Post** /processes/import | 
[**MyselfUser**](Client.md#MyselfUser) | **Get** /users/myself | 
[**RemoveGroupsFromTask**](Client.md#RemoveGroupsFromTask) | **Delete** /processes/{process_id}/tasks/{task_id}/groups | 
[**RemoveUsersFromGroup**](Client.md#RemoveUsersFromGroup) | **Delete** /groups/{id}/users | 
[**SyncGroupsToTask**](Client.md#SyncGroupsToTask) | **Post** /processes/{process_id}/tasks/{task_id}/groups | 
[**SyncUsersToGroup**](Client.md#SyncUsersToGroup) | **Post** /groups/{id}/users | 
[**UpdateEvent**](Client.md#UpdateEvent) | **Put** /processes/{process_id}/events/{event_id} | 
[**UpdateEventConnector**](Client.md#UpdateEventConnector) | **Put** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
[**UpdateFlow**](Client.md#UpdateFlow) | **Put** /processes/{process_id}/flows/{flow_id} | 
[**UpdateGateway**](Client.md#UpdateGateway) | **Put** /processes/{process_id}/gateways/{gateway_id} | 
[**UpdateGroup**](Client.md#UpdateGroup) | **Put** /groups/{id} | 
[**UpdateInputOutput**](Client.md#UpdateInputOutput) | **Put** /processes/{process_id}/tasks/{task_id}/inputoutput/{inputoutput_uid} | 
[**UpdateInstance**](Client.md#UpdateInstance) | **Put** /processes/{process_id}/instances/{instance_id} | 
[**UpdateOauthClient**](Client.md#UpdateOauthClient) | **Put** /users/{user_id}/clients/{client_id} | 
[**UpdateProcess**](Client.md#UpdateProcess) | **Put** /processes/{id} | 
[**UpdateTask**](Client.md#UpdateTask) | **Put** /processes/{process_id}/tasks/{task_id} | 
[**UpdateTaskConnector**](Client.md#UpdateTaskConnector) | **Put** /processes/{process_id}/tasks/{task_id}/connectors/{connector_id} | 
[**UpdateTaskInstance**](Client.md#UpdateTaskInstance) | **Patch** /task_instances/{task_instance_id} | 
[**UpdateUser**](Client.md#UpdateUser) | **Put** /users/{id} | 


# **AddEvent**
> EventItem AddEvent($processId, $eventCreateItem)



This method creates the new event.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process related to the event | 
 **eventCreateItem** | [**EventCreateItem**](EventCreateItem.md)| JSON API response with the event object to add | 

### Return type

[**EventItem**](EventItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddEventConnector**
> EventConnector1 AddEventConnector($processId, $eventId, $eventConnectorCreateItem)



This method is intended for creating a new event connector.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **eventId** | **string**| ID of the event to fetch | 
 **eventConnectorCreateItem** | [**EventConnectorCreateItem**](EventConnectorCreateItem.md)| JSON API with the EventConnector object to add | 

### Return type

[**EventConnector1**](EventConnector_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **DeleteEvent**
> ResultSuccess DeleteEvent($processId, $eventId)



This method deletes an event using the event ID and process ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID | 
 **eventId** | **string**| ID of the event to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventConnector**
> ResultSuccess DeleteEventConnector($processId, $eventId, $connectorId)



This method is for deleting a single event connector based on event ID, process ID and Connector ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process item | 
 **eventId** | **string**| ID of item to fetch | 
 **connectorId** | **string**| ID of EventConnector to fetch | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

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

# **DeleteOauthClient**
> ResultSuccess DeleteOauthClient($userId, $clientId)



This method deletes an Oauth client using the Oauth client and user IDs.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| User ID | 
 **clientId** | **string**| ID of Oauth client to delete | 

### Return type

[**ResultSuccess**](ResultSuccess.md)

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

# **EventTrigger**
> DataModelItem1 EventTrigger($processId, $eventId, $triggerEventCreateItem)



This method starts/triggers an event.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to the event | 
 **eventId** | **string**| ID of the event to trigger | 
 **triggerEventCreateItem** | [**TriggerEventCreateItem**](TriggerEventCreateItem.md)| Json with some parameters | 

### Return type

[**DataModelItem1**](DataModelItem_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventWebhook**
> string EventWebhook($processId, $eventId, $anyVariable)



This webhook method triggers a given event object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| Process ID related to the event | 
 **eventId** | **string**| ID of the event to trigger | 
 **anyVariable** | **string**| Any POST or GET variable will be passed to the newly created DataModel | [optional] 

### Return type

**string**

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

# **FindEventById**
> EventItem FindEventById($processId, $eventId)



This method retrieves an event using its ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to return | 
 **eventId** | **string**| ID of the event to return | 

### Return type

[**EventItem**](EventItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindEventConnectorById**
> EventConnector1 FindEventConnectorById($processId, $eventId, $connectorId)



This method returns all event connectors related to the run process and event.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **eventId** | **string**| ID of Event to fetch | 
 **connectorId** | **string**| ID of EventConnector to fetch | 

### Return type

[**EventConnector1**](EventConnector_1.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindEventConnectors**
> EventConnectorsCollection FindEventConnectors($processId, $eventId, $page, $perPage)



This method returns all event connectors related to the run process and Event.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **eventId** | **string**| ID of the task to fetch | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**EventConnectorsCollection**](EventConnectorsCollection.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindEvents**
> EventCollection FindEvents($processId, $page, $perPage)



This method returns all events related to the running process.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process related to the event | 
 **page** | **int32**| Page number to fetch | [optional] [default to 1]
 **perPage** | **int32**| Amount of items per page | [optional] [default to 15]

### Return type

[**EventCollection**](EventCollection.md)

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

# **UpdateEvent**
> EventItem UpdateEvent($processId, $eventId, $eventUpdateItem)



This method updates an existing event.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to retrieve | 
 **eventId** | **string**| ID of the event to retrieve | 
 **eventUpdateItem** | [**EventUpdateItem**](EventUpdateItem.md)| Event object to edit | 

### Return type

[**EventItem**](EventItem.md)

### Authorization

[PasswordGrant](../README.md#PasswordGrant)

### HTTP request headers

 - **Content-Type**: application/vnd.api+json
 - **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventConnector**
> EventConnector1 UpdateEventConnector($processId, $eventId, $connectorId, $eventConnectorUpdateItem)



This method updates the existing event connector with new parameter values.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processId** | **string**| ID of the process to fetch | 
 **eventId** | **string**| ID of the event to fetch | 
 **connectorId** | **string**| ID of the event Connector to fetch | 
 **eventConnectorUpdateItem** | [**EventConnectorUpdateItem**](EventConnectorUpdateItem.md)| EventConnector object to edit | 

### Return type

[**EventConnector1**](EventConnector_1.md)

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

# **UpdateOauthClient**
> OauthClientItem UpdateOauthClient($userId, $clientId, $oauthClientUpdateItem)



This method updates an existing Oauth client.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string**| ID of user to retrieve | 
 **clientId** | **string**| ID of Oauth client to retrieve | 
 **oauthClientUpdateItem** | [**OauthClientUpdateItem**](OauthClientUpdateItem.md)| Oauth Client object to edit | 

### Return type

[**OauthClientItem**](OauthClientItem.md)

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

