# \PoolsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1PoolsGet**](PoolsApi.md#ApiV1PoolsGet) | **Get** /api/v1/pools | 存储池列表


# **ApiV1PoolsGet**
> ListPoolResponseView ApiV1PoolsGet(ctx, optional)
存储池列表

存储池列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PoolsApiApiV1PoolsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PoolsApiApiV1PoolsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **optional.Int32**| 数据分页时，当前页允许返回的条目数 | [default to 0]
 **offset** | **optional.Int32**| 分页查询的偏移量 | [default to 0]
 **sortBy** | **optional.String**| 若返回的列表数据需要排序，则以该字段来排序 | 
 **orderBy** | **optional.String**| 排序的方式，升序或者降序 | 
 **name** | **optional.String**| 根据存储池名称进行过滤 | 
 **application** | **optional.String**| 根据存储池用途进行过滤 | 
 **failureDomainName** | **optional.String**| 根据物理池名称进行过滤 | 
 **safeLevel** | **optional.String**| 根据存储池的安全级别进行过滤 | 
 **healthStatus** | **optional.String**| 根据资源的健康状态进行过滤 | 
 **redundancyPloy** | **optional.String**| 根据冗余策略进行过滤 | 
 **failureDomainId** | **optional.String**| 根据物理吃uuid进行过滤 | 

### Return type

[**ListPoolResponseView**](ListPoolResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

