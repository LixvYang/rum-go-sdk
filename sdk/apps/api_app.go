/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package apps

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"reflect"
	"github.com/lixvyang/rum-sdk-go/sdk/base"
	"github.com/lixvyang/rum-sdk-go/sdk/model"
	"strings"

	"github.com/antihax/optional"
)

func New(service *base.APIClient) *AppsApiService {
	return &AppsApiService{service}
}

// Linger please
var (
	_ _context.Context
)

type service struct {
	client *base.APIClient
}

// AppsApiService AppsApi service
type AppsApiService service

// AppsApiAppApiV1GroupGroupIdContentGetOpts Optional parameters for the method 'AppApiV1GroupGroupIdContentGet'
type AppsApiAppApiV1GroupGroupIdContentGetOpts struct {
	IncludeStartTrx optional.String
	Num             optional.Int32
	Reverse         optional.String
	Senders         optional.Interface
	StartTrx        optional.String
}

/*
AppApiV1GroupGroupIdContentGet GetGroupContents
Get contents in a group
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId Group Id
  - @param groupId2
  - @param optional nil or *AppsApiAppApiV1GroupGroupIdContentGetOpts - Optional Parameters:
  - @param "IncludeStartTrx" (optional.String) -
  - @param "Num" (optional.Int32) -
  - @param "Reverse" (optional.String) -
  - @param "Senders" (optional.Interface of []string) -
  - @param "StartTrx" (optional.String) -

@return [][]PbTrx
*/
func (a *AppsApiService) AppApiV1GroupGroupIdContentGet(ctx _context.Context, groupId string, groupId2 string, localVarOptionals *AppsApiAppApiV1GroupGroupIdContentGetOpts) ([][]model.PbTrx, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  [][]model.PbTrx
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/app/api/v1/group/{group_id}/content"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", _neturl.QueryEscape(base.ParameterToString(groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("group_id", base.ParameterToString(groupId2, ""))
	if localVarOptionals != nil && localVarOptionals.IncludeStartTrx.IsSet() {
		localVarQueryParams.Add("include_start_trx", base.ParameterToString(localVarOptionals.IncludeStartTrx.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Num.IsSet() {
		localVarQueryParams.Add("num", base.ParameterToString(localVarOptionals.Num.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Reverse.IsSet() {
		localVarQueryParams.Add("reverse", base.ParameterToString(localVarOptionals.Reverse.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Senders.IsSet() {
		t := localVarOptionals.Senders.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("senders", base.ParameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("senders", base.ParameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.StartTrx.IsSet() {
		localVarQueryParams.Add("start_trx", base.ParameterToString(localVarOptionals.StartTrx.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := base.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := base.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// AppsApiAppApiV1TokenDeleteOpts Optional parameters for the method 'AppApiV1TokenDelete'
type AppsApiAppApiV1TokenDeleteOpts struct {
	GroupId optional.String
}

/*
AppApiV1TokenDelete RemoveToken
Remove a auth token
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param role
  - @param token
  - @param authorization current auth token
  - @param optional nil or *AppsApiAppApiV1TokenDeleteOpts - Optional Parameters:
  - @param "GroupId" (optional.String) -

@return UtilsSuccessResponse
*/
func (a *AppsApiService) AppApiV1TokenDelete(ctx _context.Context, role string, token string, authorization string, localVarOptionals *AppsApiAppApiV1TokenDeleteOpts) (model.UtilsSuccessResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  model.UtilsSuccessResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/app/api/v1/token"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.GroupId.IsSet() {
		localVarQueryParams.Add("group_id", base.ParameterToString(localVarOptionals.GroupId.Value(), ""))
	}
	localVarQueryParams.Add("role", base.ParameterToString(role, ""))
	localVarQueryParams.Add("token", base.ParameterToString(token, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := base.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := base.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Authorization"] = base.ParameterToString(authorization, "")
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
AppApiV1TokenListGet ListToken
List all auth token
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param authorization current auth token

@return OptionsJwt
*/
func (a *AppsApiService) AppApiV1TokenListGet(ctx _context.Context, authorization string) (model.OptionsJwt, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  model.OptionsJwt
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/app/api/v1/token/list"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := base.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := base.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Authorization"] = base.ParameterToString(authorization, "")
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// AppsApiAppApiV1TokenPostOpts Optional parameters for the method 'AppApiV1TokenPost'
type AppsApiAppApiV1TokenPostOpts struct {
	AppapiCreateJwtParams optional.Interface
}

/*
AppApiV1TokenPost CreateToken
Create a new auth token, only allow access from localhost
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *AppsApiAppApiV1TokenPostOpts - Optional Parameters:
  - @param "AppapiCreateJwtParams" (optional.Interface of AppapiCreateJwtParams) -

@return AppapiTokenItem
*/
func (a *AppsApiService) AppApiV1TokenPost(ctx _context.Context, localVarOptionals *AppsApiAppApiV1TokenPostOpts) (model.AppapiTokenItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  model.AppapiTokenItem
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/app/api/v1/token"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := base.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := base.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.AppapiCreateJwtParams.IsSet() {
		localVarOptionalAppapiCreateJwtParams, localVarOptionalAppapiCreateJwtParamsok := localVarOptionals.AppapiCreateJwtParams.Value().(model.AppapiCreateJwtParams)
		if !localVarOptionalAppapiCreateJwtParamsok {
			return localVarReturnValue, nil, base.ReportError("appapiCreateJwtParams should be AppapiCreateJwtParams")
		}
		localVarPostBody = &localVarOptionalAppapiCreateJwtParams
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

/*
AppApiV1TokenRefreshPost RefreshToken
Get a new auth token
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param authorization current auth token

@return AppapiTokenItem
*/
func (a *AppsApiService) AppApiV1TokenRefreshPost(ctx _context.Context, authorization string) (model.AppapiTokenItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  model.AppapiTokenItem
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/app/api/v1/token/refresh"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := base.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := base.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Authorization"] = base.ParameterToString(authorization, "")
	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// AppsApiAppApiV1TokenRevokePostOpts Optional parameters for the method 'AppApiV1TokenRevokePost'
type AppsApiAppApiV1TokenRevokePostOpts struct {
	AppapiRevokeJwtParams optional.Interface
}

/*
AppApiV1TokenRevokePost RevokeToken
Revoke a auth token
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param authorization current auth token
  - @param optional nil or *AppsApiAppApiV1TokenRevokePostOpts - Optional Parameters:
  - @param "AppapiRevokeJwtParams" (optional.Interface of AppapiRevokeJwtParams) -

@return UtilsSuccessResponse
*/
func (a *AppsApiService) AppApiV1TokenRevokePost(ctx _context.Context, authorization string, localVarOptionals *AppsApiAppApiV1TokenRevokePostOpts) (model.UtilsSuccessResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  model.UtilsSuccessResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/app/api/v1/token/revoke"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := base.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := base.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["Authorization"] = base.ParameterToString(authorization, "")
	// body params
	if localVarOptionals != nil && localVarOptionals.AppapiRevokeJwtParams.IsSet() {
		localVarOptionalAppapiRevokeJwtParams, localVarOptionalAppapiRevokeJwtParamsok := localVarOptionals.AppapiRevokeJwtParams.Value().(model.AppapiRevokeJwtParams)
		if !localVarOptionalAppapiRevokeJwtParamsok {
			return localVarReturnValue, nil, base.ReportError("appapiRevokeJwtParams should be AppapiRevokeJwtParams")
		}
		localVarPostBody = &localVarOptionalAppapiRevokeJwtParams
	}

	r, err := a.client.PrepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.CallAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := base.GenericOpenAPIError{
			Msg: localVarBody,
			Err: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}