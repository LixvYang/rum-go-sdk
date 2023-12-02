/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package user

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/lixvyang/rum-sdk-go/sdk/base"
	"github.com/lixvyang/rum-sdk-go/sdk/model"

	"github.com/antihax/optional"
)

func New(service *base.APIClient) *UserApiService {
	return &UserApiService{service}
}

// Linger please
var (
	_ _context.Context
)

type service struct {
	client *base.APIClient
}

// UserApiService UserApi service
type UserApiService service

// UserApiApiV1GroupAnnouncePostOpts Optional parameters for the method 'ApiV1GroupAnnouncePost'
type UserApiApiV1GroupAnnouncePostOpts struct {
	HandlersAnnounceParam optional.Interface
}

/*
ApiV1GroupAnnouncePost AnnounceUserPubkey
Announce User&#39;s encryption Pubkey to the group
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param optional nil or *UserApiApiV1GroupAnnouncePostOpts - Optional Parameters:
  - @param "HandlersAnnounceParam" (optional.Interface of HandlersAnnounceParam) -

@return model.HandlersAnnounceResult
*/
func (a *UserApiService) ApiV1GroupAnnouncePost(ctx _context.Context, localVarOptionals *UserApiApiV1GroupAnnouncePostOpts) (model.HandlersAnnounceResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  model.HandlersAnnounceResult
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/api/v1/group/announce"
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
	if localVarOptionals != nil && localVarOptionals.HandlersAnnounceParam.IsSet() {
		localVarOptionalHandlersAnnounceParam, localVarOptionalHandlersAnnounceParamok := localVarOptionals.HandlersAnnounceParam.Value().(model.HandlersAnnounceParam)
		if !localVarOptionalHandlersAnnounceParamok {
			return localVarReturnValue, nil, base.ReportError("handlersAnnounceParam should be HandlersAnnounceParam")
		}
		localVarPostBody = &localVarOptionalHandlersAnnounceParam
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
ApiV1GroupGroupIdAnnouncedProducersGet GetAnnouncedGroupProducer
Get the list of group producers
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId Group Id

@return []model.HandlersAnnouncedProducerListItem
*/
func (a *UserApiService) ApiV1GroupGroupIdAnnouncedProducersGet(ctx _context.Context, groupId string) ([]model.HandlersAnnouncedProducerListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []model.HandlersAnnouncedProducerListItem
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/api/v1/group/{group_id}/announced/producers"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", _neturl.QueryEscape(base.ParameterToString(groupId, "")), -1)

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
ApiV1GroupGroupIdAnnouncedUserSignPubkeyGet GetAnnouncedGroupUser
Get the one user announce status
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId Group Id
  - @param signPubkey User SignPubkey

@return model.HandlersAnnouncedUserListItem
*/
func (a *UserApiService) ApiV1GroupGroupIdAnnouncedUserSignPubkeyGet(ctx _context.Context, groupId string, signPubkey string) (model.HandlersAnnouncedUserListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  model.HandlersAnnouncedUserListItem
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/api/v1/group/{group_id}/announced/user/{sign_pubkey}"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", _neturl.QueryEscape(base.ParameterToString(groupId, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"sign_pubkey"+"}", _neturl.QueryEscape(base.ParameterToString(signPubkey, "")), -1)

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
ApiV1GroupGroupIdAnnouncedUsersGet GetAnnouncedGroupUsers
Get the list of private group users
  - @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param groupId Group Id

@return []model.HandlersAnnouncedUserListItem
*/
func (a *UserApiService) ApiV1GroupGroupIdAnnouncedUsersGet(ctx _context.Context, groupId string) ([]model.HandlersAnnouncedUserListItem, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []model.HandlersAnnouncedUserListItem
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/api/v1/group/{group_id}/announced/users"
	localVarPath = strings.Replace(localVarPath, "{"+"group_id"+"}", _neturl.QueryEscape(base.ParameterToString(groupId, "")), -1)

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
