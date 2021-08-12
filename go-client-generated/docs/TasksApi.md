# \TasksApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1TasksTaskIdActionPost**](TasksApi.md#ApiV1TasksTaskIdActionPost) | **Post** /api/v1/tasks/{task_id}/{action} | 任务结束处理
[**ApiV1TasksTaskIdGet**](TasksApi.md#ApiV1TasksTaskIdGet) | **Get** /api/v1/tasks/{task_id} | 任务详情


# **ApiV1TasksTaskIdActionPost**
> AsyncTaskResponseView ApiV1TasksTaskIdActionPost(ctx, taskId, action)
任务结束处理

任务结束处理

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| 处理task任务id | 
  **action** | **string**| 处理task任务动作 | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1TasksTaskIdGet**
> TaskDetailResponseView ApiV1TasksTaskIdGet(ctx, taskId)
任务详情

任务详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| 任务的UUID | 

### Return type

[**TaskDetailResponseView**](TaskDetailResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

