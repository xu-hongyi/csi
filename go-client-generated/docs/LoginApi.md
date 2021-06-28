# \LoginApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1LoginPost**](LoginApi.md#ApiV1LoginPost) | **Post** /api/v1/login | 登录


# **ApiV1LoginPost**
> LoginResponseView ApiV1LoginPost(ctx, body)
登录

登录

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LoginRequestView**](LoginRequestView.md)|  | 

### Return type

[**LoginResponseView**](LoginResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

