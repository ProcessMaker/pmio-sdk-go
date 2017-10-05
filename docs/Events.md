# \Events

All URIs are relative to *https://CHANGEME.api.processmaker.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEvent**](Events.md#AddEvent) | **Post** /processes/{process_id}/events | 
[**AddEventConnector**](Events.md#AddEventConnector) | **Post** /processes/{process_id}/events/{event_id}/connectors | 
[**DeleteEvent**](Events.md#DeleteEvent) | **Delete** /processes/{process_id}/events/{event_id} | 
[**DeleteEventConnector**](Events.md#DeleteEventConnector) | **Delete** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
[**EventTrigger**](Events.md#EventTrigger) | **Post** /processes/{process_id}/events/{event_id}/trigger | 
[**EventWebhook**](Events.md#EventWebhook) | **Post** /processes/{process_id}/events/{event_id}/webhook | 
[**FindEventById**](Events.md#FindEventById) | **Get** /processes/{process_id}/events/{event_id} | 
[**FindEventConnectorById**](Events.md#FindEventConnectorById) | **Get** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 
[**FindEventConnectors**](Events.md#FindEventConnectors) | **Get** /processes/{process_id}/events/{event_id}/connectors | 
[**FindEvents**](Events.md#FindEvents) | **Get** /processes/{process_id}/events | 
[**UpdateEvent**](Events.md#UpdateEvent) | **Put** /processes/{process_id}/events/{event_id} | 
[**UpdateEventConnector**](Events.md#UpdateEventConnector) | **Put** /processes/{process_id}/events/{event_id}/connectors/{connector_id} | 


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

