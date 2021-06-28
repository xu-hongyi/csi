# \SnapsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BlockSnapsBatchDeletionPost**](SnapsApi.md#ApiV1BlockSnapsBatchDeletionPost) | **Post** /api/v1/block/snaps/batch/deletion | 批量删除快照
[**ApiV1BlockSnapsGet**](SnapsApi.md#ApiV1BlockSnapsGet) | **Get** /api/v1/block/snaps | 查询快照列表
[**ApiV1BlockSnapsPost**](SnapsApi.md#ApiV1BlockSnapsPost) | **Post** /api/v1/block/snaps | 创建快照
[**ApiV1BlockSnapsSnapshotIdClonePost**](SnapsApi.md#ApiV1BlockSnapsSnapshotIdClonePost) | **Post** /api/v1/block/snaps/{snapshot_id}/clone | 克隆快照
[**ApiV1BlockSnapsSnapshotIdDelete**](SnapsApi.md#ApiV1BlockSnapsSnapshotIdDelete) | **Delete** /api/v1/block/snaps/{snapshot_id} | 删除快照
[**ApiV1BlockSnapsSnapshotIdGet**](SnapsApi.md#ApiV1BlockSnapsSnapshotIdGet) | **Get** /api/v1/block/snaps/{snapshot_id} | 根据ID查询快照的详情信息
[**ApiV1BlockSnapsSnapshotIdPut**](SnapsApi.md#ApiV1BlockSnapsSnapshotIdPut) | **Put** /api/v1/block/snaps/{snapshot_id} | 修改快照的名字、描述信息


# **ApiV1BlockSnapsBatchDeletionPost**
> AsyncTaskResponseView ApiV1BlockSnapsBatchDeletionPost(ctx, body)
批量删除快照

批量删除快照

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteBatchSnapshotRequestView**](DeleteBatchSnapshotRequestView.md)|  | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockSnapsGet**
> ListSnapResponseView ApiV1BlockSnapsGet(ctx, optional)
查询快照列表

查询快照列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SnapsApiApiV1BlockSnapsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapsApiApiV1BlockSnapsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **optional.Int32**| 数据分页时，当前页允许返回的条目数 | [default to 0]
 **offset** | **optional.Int32**| 分页查询的偏移量 | [default to 0]
 **sortBy** | **optional.String**| 若返回的列表数据需要排序，则以该字段来排序 | 
 **orderBy** | **optional.String**| 排序的方式，升序或者降序 | 
 **name** | **optional.String**| 根据名称进行过滤 | 
 **id** | **optional.String**| 根据ID进行过滤 | 
 **poolName** | **optional.String**| 通过存储池名称进行过滤 | 
 **poolId** | **optional.String**| 通过存储池ID进行过滤 | 
 **volumeId** | **optional.String**| 通过卷的uuid进行过滤 | 
 **volumeName** | **optional.String**| 通过卷名称进行过滤，支持模糊匹配 | 
 **cgSnapId** | **optional.String**| 通过一致性组快照的uuid进行过滤 | 
 **cgSnapName** | **optional.String**| 通过一致性组快照的名称进行过滤，支持模糊匹配 | 
 **cloneVolNum** | **optional.Float32**| 通过克隆卷数量进行过滤 | 
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 
 **capHealthStatus** | **optional.String**| 根据资源的容量健康状态进行过滤 | 
 **syncStatus** | **optional.String**| 通过同步状态进行过滤 | 
 **createFrom** | **optional.String**| 通过创建途径进行过滤 | 
 **description** | **optional.String**| 通过描述信息进行过滤，支持模糊匹配 | 

### Return type

[**ListSnapResponseView**](ListSnapResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockSnapsPost**
> AsyncTaskResponseView ApiV1BlockSnapsPost(ctx, body)
创建快照

创建快照

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSnapshotRequestView**](CreateSnapshotRequestView.md)|  | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockSnapsSnapshotIdClonePost**
> AsyncTaskResponseView ApiV1BlockSnapsSnapshotIdClonePost(ctx, body, snapshotId)
克隆快照

克隆快照

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SnapCloneRequestView**](SnapCloneRequestView.md)|  | 
  **snapshotId** | **string**| 快照的uuid | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockSnapsSnapshotIdDelete**
> AsyncTaskResponseView ApiV1BlockSnapsSnapshotIdDelete(ctx, snapshotId)
删除快照

删除快照

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| 快照的uuid | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockSnapsSnapshotIdGet**
> SnapDetailResponseView ApiV1BlockSnapsSnapshotIdGet(ctx, snapshotId)
根据ID查询快照的详情信息

根据ID查询快照的详情信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **snapshotId** | **string**| 快照的uuid | 

### Return type

[**SnapDetailResponseView**](SnapDetailResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockSnapsSnapshotIdPut**
> RenameSnapResponseView ApiV1BlockSnapsSnapshotIdPut(ctx, body, snapshotId)
修改快照的名字、描述信息

修改快照的名字、描述信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RenameSnapRequestView**](RenameSnapRequestView.md)|  | 
  **snapshotId** | **string**| 快照的uuid | 

### Return type

[**RenameSnapResponseView**](RenameSnapResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

