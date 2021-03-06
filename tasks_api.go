/* 
 * ProcessMaker API
 *
 * This ProcessMaker I/O API provides access to a BPMN 2.0 compliant workflow engine API that is designed to be used as a microservice to support enterprise cloud applications. The current Alpha 1.0 version supports most of the descriptive classes of the BPMN 2.0 specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: support@processmaker.io
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pmio

import (
	"net/url"
	"encoding/json"
	"fmt"
	"strings"
)

type Tasks struct {
	Configuration Configuration
}

func NewTasks() *Tasks {
	configuration := NewConfiguration()
	return &Tasks{
		Configuration: *configuration,
	}
}

func NewTasksWithBasePath(basePath string) *Tasks {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &Tasks{
		Configuration: *configuration,
	}
}

/**
 * 
 * This method assigns group(s) to the chosen task
 *
 * @param processId Process ID
 * @param taskId ID of the task to be modified
 * @param taskAddGroupsItem JSON API with Group IDs to add
 * @return *ResultSuccess
 */
func (a Tasks) AddGroupsToTask(processId string, taskId string, taskAddGroupsItem TaskAddGroupsItem) (*ResultSuccess, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}/groups"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &taskAddGroupsItem

	var successPayload = new(ResultSuccess)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method creates a new task.
 *
 * @param processId Process ID related to the task
 * @param taskCreateItem JSON API with the task object to add
 * @return *TaskItem
 */
func (a Tasks) AddTask(processId string, taskCreateItem TaskCreateItem) (*TaskItem, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &taskCreateItem

	var successPayload = new(TaskItem)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method is for creating a new task connector
 *
 * @param processId ID of the process to fetch
 * @param taskId ID of the task to fetch
 * @param taskConnectorCreateItem JSON API with the TaskConnector object to add
 * @return *TaskConnector1
 */
func (a Tasks) AddTaskConnector(processId string, taskId string, taskConnectorCreateItem TaskConnectorCreateItem) (*TaskConnector1, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}/connectors"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &taskConnectorCreateItem

	var successPayload = new(TaskConnector1)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method deletes a task using the task ID and the process ID.
 *
 * @param processId Process ID
 * @param taskId ID of a task to delete
 * @return *ResultSuccess
 */
func (a Tasks) DeleteTask(processId string, taskId string) (*ResultSuccess, *APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ResultSuccess)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method is for deleting a single task connector based on task ID, the process ID and the Connector ID.
 *
 * @param processId ID of the process item to fetch
 * @param taskId ID of the task item to fetch
 * @param connectorId ID of TaskConnector to fetch
 * @return *ResultSuccess
 */
func (a Tasks) DeleteTaskConnector(processId string, taskId string, connectorId string) (*ResultSuccess, *APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}/connectors/{connector_id}"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)
	path = strings.Replace(path, "{"+"connector_id"+"}", fmt.Sprintf("%v", connectorId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(ResultSuccess)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method retrieves a task using its ID.
 *
 * @param processId ID of the process to return
 * @param taskId ID of the task to return
 * @return *TaskItem
 */
func (a Tasks) FindTaskById(processId string, taskId string) (*TaskItem, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(TaskItem)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method is for retrieving a task connector based on its ID.
 *
 * @param processId ID of the process to fetch
 * @param taskId ID of the task to fetch
 * @param connectorId ID of TaskConnector to fetch
 * @return *TaskConnector1
 */
func (a Tasks) FindTaskConnectorById(processId string, taskId string, connectorId string) (*TaskConnector1, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}/connectors/{connector_id}"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)
	path = strings.Replace(path, "{"+"connector_id"+"}", fmt.Sprintf("%v", connectorId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(TaskConnector1)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method returns all task connectors related to the run process and task.
 *
 * @param processId ID of the process to fetch
 * @param taskId ID of the task to fetch
 * @param page Page number to fetch
 * @param perPage Amount of items per page
 * @return *TaskConnectorsCollection
 */
func (a Tasks) FindTaskConnectors(processId string, taskId string, page int32, perPage int32) (*TaskConnectorsCollection, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}/connectors"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
			queryParams.Add("per_page", a.Configuration.APIClient.ParameterToString(perPage, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(TaskConnectorsCollection)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method retrieves a task instance based on its ID.
 *
 * @param taskInstanceId ID of task instance to return
 * @param page Page number to fetch
 * @param perPage Amount of items per page
 * @return *InlineResponse200
 */
func (a Tasks) FindTaskInstanceById(taskInstanceId string, page int32, perPage int32) (*InlineResponse200, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/task_instances/{task_instance_id}"
	path = strings.Replace(path, "{"+"task_instance_id"+"}", fmt.Sprintf("%v", taskInstanceId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
			queryParams.Add("per_page", a.Configuration.APIClient.ParameterToString(perPage, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponse200)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method retrieves all existing task instances.
 *
 * @param page Page number to fetch
 * @param perPage Amount of items per page
 * @return *TaskInstanceCollection
 */
func (a Tasks) FindTaskInstances(page int32, perPage int32) (*TaskInstanceCollection, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/task_instances"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
			queryParams.Add("per_page", a.Configuration.APIClient.ParameterToString(perPage, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(TaskInstanceCollection)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method is intended for returning a list of all tasks related to the process.
 *
 * @param processId ID of the process relative to the task
 * @param page Page number to fetch
 * @param perPage Amount of items per page
 * @return *TaskCollection
 */
func (a Tasks) FindTasks(processId string, page int32, perPage int32) (*TaskCollection, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}
		queryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
			queryParams.Add("per_page", a.Configuration.APIClient.ParameterToString(perPage, ""))
	

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(TaskCollection)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method removes groups from a task.
 *
 * @param processId Process ID
 * @param taskId Task ID
 * @param taskRemoveGroupsItem JSON API response with Group IDs to remove
 * @return *ResultSuccess
 */
func (a Tasks) RemoveGroupsFromTask(processId string, taskId string, taskRemoveGroupsItem TaskRemoveGroupsItem) (*ResultSuccess, *APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}/groups"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &taskRemoveGroupsItem

	var successPayload = new(ResultSuccess)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method synchronizes one or more groups with a task.
 *
 * @param processId Process ID
 * @param taskId ID of the task to modify
 * @param taskSyncGroupsItem JSON API response with group IDs to sync
 * @return *ResultSuccess
 */
func (a Tasks) SyncGroupsToTask(processId string, taskId string, taskSyncGroupsItem TaskSyncGroupsItem) (*ResultSuccess, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}/groups"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &taskSyncGroupsItem

	var successPayload = new(ResultSuccess)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method is for updating an existing task.
 *
 * @param processId ID of the process to fetch
 * @param taskId ID of the task to fetch
 * @param taskUpdateItem Task object to edit
 * @return *TaskItem
 */
func (a Tasks) UpdateTask(processId string, taskId string, taskUpdateItem TaskUpdateItem) (*TaskItem, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &taskUpdateItem

	var successPayload = new(TaskItem)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method updates the existing task connector with new parameter values.
 *
 * @param processId ID of the process to fetch
 * @param taskId ID of the task to fetch
 * @param connectorId ID of the task connector to fetch
 * @param taskConnectorUpdateItem TaskConnector object to edit
 * @return *TaskConnector1
 */
func (a Tasks) UpdateTaskConnector(processId string, taskId string, connectorId string, taskConnectorUpdateItem TaskConnectorUpdateItem) (*TaskConnector1, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/processes/{process_id}/tasks/{task_id}/connectors/{connector_id}"
	path = strings.Replace(path, "{"+"process_id"+"}", fmt.Sprintf("%v", processId), -1)
	path = strings.Replace(path, "{"+"task_id"+"}", fmt.Sprintf("%v", taskId), -1)
	path = strings.Replace(path, "{"+"connector_id"+"}", fmt.Sprintf("%v", connectorId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &taskConnectorUpdateItem

	var successPayload = new(TaskConnector1)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * 
 * This method updates an existing task instance.
 *
 * @param taskInstanceId ID of the task instance to retrieve
 * @param taskInstanceUpdateItem Task instance object to update
 * @return *InlineResponse200
 */
func (a Tasks) UpdateTaskInstance(taskInstanceId string, taskInstanceUpdateItem TaskInstanceUpdateItem) (*InlineResponse200, *APIResponse, error) {

	var httpMethod = "Patch"
	// create path and map variables
	path := a.Configuration.BasePath + "/task_instances/{task_instance_id}"
	path = strings.Replace(path, "{"+"task_instance_id"+"}", fmt.Sprintf("%v", taskInstanceId), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication '(PasswordGrant)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		headerParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/vnd.api+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/vnd.api+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &taskInstanceUpdateItem

	var successPayload = new(InlineResponse200)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

