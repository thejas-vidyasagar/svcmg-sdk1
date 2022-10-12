// Api classes for service-manager-for-msp's golang SDK
package api

import (
	"encoding/json"
	"github.com/thejas-vidyasagar/svcmg-sdk1/service-manager-for-msp-go-client/v16/client"
	import1 "github.com/thejas-vidyasagar/svcmg-sdk1/service-manager-for-msp-go-client/v16/models/msp/v1/svcmgr"
	"net/http"
	"net/url"
	"strings"
)

type ServiceManagerApi struct {
	ApiClient *client.ApiClient
}

func NewServiceManagerApi(apiClient *client.ApiClient) *ServiceManagerApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &ServiceManagerApi{
		ApiClient: apiClient,
	}
	return a
}

/*
*

	API to uninstall an application
	API to uninstall an application running on the Kuberentes platform. The user needs to provide appUUID for the application that needs to be uninstalled.

	parameters:-
	-> appUUID (string) (required) : UUID for the Application installed by Service Manager
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.DeleteAppApiResponse, error)
*/
func (api *ServiceManagerApi) DeleteApp(appUUID *string, args ...map[string]interface{}) (*import1.DeleteAppApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/applications/{appUUID}"

	// verify the required parameter 'appUUID' is set
	if nil == appUUID {
		return nil, client.ReportError("appUUID is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"appUUID"+"}", url.PathEscape(client.ParameterToString(*appUUID, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.DeleteAppApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	API to uninstall a service
	API to uninstall a service running on the Kuberentes platform. The user needs to provide serviceUUID for the service that needs to be uninstalled.

	parameters:-
	-> serviceUUID (string) (required) : UUID for the Service installed by Service Manager
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.DeleteServiceApiResponse, error)
*/
func (api *ServiceManagerApi) DeleteService(serviceUUID *string, args ...map[string]interface{}) (*import1.DeleteServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/services/{serviceUUID}"

	// verify the required parameter 'serviceUUID' is set
	if nil == serviceUUID {
		return nil, client.ReportError("serviceUUID is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"serviceUUID"+"}", url.PathEscape(client.ParameterToString(*serviceUUID, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.DeleteServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	API to get the status of all the applications
	API to get the status of all the applications along with its group of sub-services that are installed through the Service Manager.

	parameters:-
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.GetAllAppApiResponse, error)
*/
func (api *ServiceManagerApi) GetAllApp(args ...map[string]interface{}) (*import1.GetAllAppApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/applications"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetAllAppApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	API to get the status of all the services
	API to get the status of all the services that are installed through the Service Manager.

	parameters:-
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.GetAllServiceApiResponse, error)
*/
func (api *ServiceManagerApi) GetAllService(args ...map[string]interface{}) (*import1.GetAllServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/services"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetAllServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	API to get the status of all the tasks
	API to get the status of all the tasks along with its status of all the child tasks (if any) of the tasks. The user can to provide the following in the query param:   - uuid of the application or service. If uuid is given name and version will be ignored.   - name   - version   - type can be used as one of \"Application\" or \"Service\"   - limit   - offset

	parameters:-
	-> name (string) (optional) : Name of the application or service
	-> version (string) (optional) : Version of the application or service
	-> type_ (string) (optional) : Type is one of \"Application\" or \"Service\"
	-> limit (int) (optional) : The numbers of task records to return. (-1 for no limit)
	-> offset (int) (optional) : Offset from where to retrieve task records
	-> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page.
	-> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100.
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.GetAllTaskApiResponse, error)
*/
func (api *ServiceManagerApi) GetAllTask(name *string, version *string, type_ *string, limit *int, offset *int, page_ *int, limit_ *int, args ...map[string]interface{}) (*import1.GetAllTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/tasks"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if name != nil {
		queryParams.Add("name", client.ParameterToString(*name, ""))
	}
	if version != nil {
		queryParams.Add("version", client.ParameterToString(*version, ""))
	}
	if type_ != nil {
		queryParams.Add("type", client.ParameterToString(*type_, ""))
	}
	if limit != nil {
		queryParams.Add("limit", client.ParameterToString(*limit, ""))
	}
	if offset != nil {
		queryParams.Add("offset", client.ParameterToString(*offset, ""))
	}
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetAllTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	API to get the status of the requested application
	API to get the details of a specific application along with its group of services that are installed through the Service Manager. The user needs to provide with appUUID of the application in the path parameters in the API.

	parameters:-
	-> appUUID (string) (required) : UUID for the Application installed by Service Manager
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.GetAppApiResponse, error)
*/
func (api *ServiceManagerApi) GetApp(appUUID *string, args ...map[string]interface{}) (*import1.GetAppApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/applications/{appUUID}"

	// verify the required parameter 'appUUID' is set
	if nil == appUUID {
		return nil, client.ReportError("appUUID is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"appUUID"+"}", url.PathEscape(client.ParameterToString(*appUUID, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetAppApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	Api to get the the health of the service manager.
	Checks the health of the service manager.

	parameters:-
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.HealthApiResponse, error)
*/
func (api *ServiceManagerApi) GetHealth(args ...map[string]interface{}) (*import1.HealthApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/health"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.HealthApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	API to get the history of the applications and services.
	API to get the history of the applications and services. The user can to provide the following in the query param:   - uuid of the application or service. If uuid is given name and version     will be ignored.   - name   - version   - type can be used as one of \"Application\" or \"Service\"   - limit   - offset

	parameters:-
	-> name (string) (optional) : Name of the application or service
	-> version (string) (optional) : Version of the application or service
	-> uuid (string) (optional) : UUID of the application or service
	-> type_ (string) (optional) : Type is one of \"Application\" or \"Service\"
	-> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page.
	-> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100.
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.GetHistoryApiResponse, error)
*/
func (api *ServiceManagerApi) GetHistory(name *string, version *string, uuid *string, type_ *string, page_ *int, limit_ *int, args ...map[string]interface{}) (*import1.GetHistoryApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/history"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if name != nil {
		queryParams.Add("name", client.ParameterToString(*name, ""))
	}
	if version != nil {
		queryParams.Add("version", client.ParameterToString(*version, ""))
	}
	if uuid != nil {
		queryParams.Add("uuid", client.ParameterToString(*uuid, ""))
	}
	if type_ != nil {
		queryParams.Add("type", client.ParameterToString(*type_, ""))
	}
	if page_ != nil {
		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {
		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetHistoryApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	API to get the status of the requested service
	API to get the details of a specific service that are installed through the Service Manager. The user needs to provide with the serviceUUID of the service in the path parameters in the API.

	parameters:-
	-> serviceUUID (string) (required) : UUID for the Service installed by Service Manager
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.GetServiceApiResponse, error)
*/
func (api *ServiceManagerApi) GetService(serviceUUID *string, args ...map[string]interface{}) (*import1.GetServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/services/{serviceUUID}"

	// verify the required parameter 'serviceUUID' is set
	if nil == serviceUUID {
		return nil, client.ReportError("serviceUUID is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"serviceUUID"+"}", url.PathEscape(client.ParameterToString(*serviceUUID, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	API to get the task status
	API to get the details of a specific task of Service Manager. The user needs to provide the task uuid for the task. The API will fetch the task status of the task uuid provided as well as task status of all the child tasks (if any) of the task uuid provided by the user

	parameters:-
	-> taskUUID (string) (required) : Task UUID for the task
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.GetTaskApiResponse, error)
*/
func (api *ServiceManagerApi) GetTask(taskUUID *string, args ...map[string]interface{}) (*import1.GetTaskApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/tasks/{taskUUID}"

	// verify the required parameter 'taskUUID' is set
	if nil == taskUUID {
		return nil, client.ReportError("taskUUID is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"taskUUID"+"}", url.PathEscape(client.ParameterToString(*taskUUID, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.GetTaskApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	Deploys a new application on the kubernetes platform.
	Deploys a new application on the kubernetes platform.<br> The user needs to specify the following: <br> name: The name of the application that needs to be deployed. <br> version: The version of the application that needs to be deployed.

	parameters:-
	-> body (msp.v1.svcmgr.Install) (required)
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.PostAppApiResponse, error)
*/
func (api *ServiceManagerApi) PostApp(body *import1.Install, args ...map[string]interface{}) (*import1.PostAppApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/applications"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.PostAppApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	Deploys a new SMSP cluster. Different deployment models can be choosen based on deploymentType and availabilityType:   Combined VM Non HA - Minimum resource - No resiliency (recommended for service developers)   Combined VM HA - Optimized resource  with resiliency   Scale out VM HA - Fully resilient and used for shared wider deployments.   Scale out VM Non HA - No system process resiliency and used for wider deployments.
	Deploy SMSP cluster.

	parameters:-
	-> body (msp.v1.svcmgr.Cluster) (required)
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.PostClusterApiResponse, error)
*/
func (api *ServiceManagerApi) PostCluster(body *import1.Cluster, args ...map[string]interface{}) (*import1.PostClusterApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/clusters"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.PostClusterApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/*
*

	Deploys a new service on the kubernetes platform.
	Deploys a new service on the kubernetes platform.<br> The user needs to specify the following: <br> name: The name of the service that needs to be deployed. <br> version: The version of the service that needs to be deployed.

	parameters:-
	-> body (msp.v1.svcmgr.Install) (required)
	-> args (map[string]interface{}) (optional) : Additional Arguments

	returns: (*msp.v1.svcmgr.PostServiceApiResponse, error)
*/
func (api *ServiceManagerApi) PostService(body *import1.Install, args ...map[string]interface{}) (*import1.PostServiceApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/msp/v1.0.a1/svcmgr/services"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Header Params
	if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
		headerParams["If-Match"] = ifMatch
	}
	if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
		headerParams["If-None-Match"] = ifNoneMatch
	}
	authNames := []string{}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import1.PostServiceApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
