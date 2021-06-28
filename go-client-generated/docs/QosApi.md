# \QosApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BlockQosGet**](QosApi.md#ApiV1BlockQosGet) | **Get** /api/v1/block/qos | 查询Qos列表
[**ApiV1BlockQosPost**](QosApi.md#ApiV1BlockQosPost) | **Post** /api/v1/block/qos | 创建Volume的QOS策略
[**ApiV1BlockQosQosPolicyIdDelete**](QosApi.md#ApiV1BlockQosQosPolicyIdDelete) | **Delete** /api/v1/block/qos/{qos_policy_id} | 删除Volume的QOS策略
[**ApiV1BlockQosQosPolicyIdGet**](QosApi.md#ApiV1BlockQosQosPolicyIdGet) | **Get** /api/v1/block/qos/{qos_policy_id} | 根据ID查询Qos的详情信息
[**ApiV1BlockQosQosPolicyIdPut**](QosApi.md#ApiV1BlockQosQosPolicyIdPut) | **Put** /api/v1/block/qos/{qos_policy_id} | 修改Volume的QOS策略


# **ApiV1BlockQosGet**
> ListQosResponseView ApiV1BlockQosGet(ctx, optional)
查询Qos列表

查询Qos列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QosApiApiV1BlockQosGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QosApiApiV1BlockQosGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **optional.Int32**| 数据分页时，当前页允许返回的条目数 | [default to 0]
 **offset** | **optional.Int32**| 分页查询的偏移量 | [default to 0]
 **sortBy** | **optional.String**| 若返回的列表数据需要排序，则以该字段来排序 | 
 **orderBy** | **optional.String**| 排序的方式，升序或者降序 | 
 **name** | **optional.String**| 根据名称进行过滤 | 
 **id** | **optional.String**| 根据ID进行过滤 | 
 **iopsLimit** | **optional.Int32**| 通过最大IOPS进行过滤 | 
 **iopsBurst** | **optional.Int32**| 通过突发IOPS进行过滤 | 
 **bpsLimit** | **optional.Int32**| 通过最大带宽进行过滤 | 
 **bpsBurst** | **optional.Int32**| 通过突发带宽进行过滤 | 

### Return type

[**ListQosResponseView**](ListQosResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockQosPost**
> CreateQosResponseView ApiV1BlockQosPost(ctx, body)
创建Volume的QOS策略

创建Volume的QOS策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateQosRequestView**](CreateQosRequestView.md)|  | 

### Return type

[**CreateQosResponseView**](CreateQosResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockQosQosPolicyIdDelete**
> DeleteQosResponseView ApiV1BlockQosQosPolicyIdDelete(ctx, qosPolicyId)
删除Volume的QOS策略

删除Volume的QOS策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **qosPolicyId** | **string**| qos策略的uuid | 

### Return type

[**DeleteQosResponseView**](DeleteQosResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockQosQosPolicyIdGet**
> QosDetailResponseView ApiV1BlockQosQosPolicyIdGet(ctx, qosPolicyId)
根据ID查询Qos的详情信息

根据ID查询Qos的详情信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **qosPolicyId** | **string**| Qos的uuid | 

### Return type

[**QosDetailResponseView**](QosDetailResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockQosQosPolicyIdPut**
> UpdateQosResponseView ApiV1BlockQosQosPolicyIdPut(ctx, body, qosPolicyId)
修改Volume的QOS策略

修改Volume的QOS策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateQosRequestView**](UpdateQosRequestView.md)|  | 
  **qosPolicyId** | **string**| qos策略的uuid | 

### Return type

[**UpdateQosResponseView**](UpdateQosResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

