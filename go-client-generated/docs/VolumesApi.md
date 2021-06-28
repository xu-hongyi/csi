# \VolumesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BlockVolumesBatchCreationPost**](VolumesApi.md#ApiV1BlockVolumesBatchCreationPost) | **Post** /api/v1/block/volumes/batch/creation | 批量创建Volume
[**ApiV1BlockVolumesBatchDeletionPost**](VolumesApi.md#ApiV1BlockVolumesBatchDeletionPost) | **Post** /api/v1/block/volumes/batch/deletion | 批量删除卷
[**ApiV1BlockVolumesGet**](VolumesApi.md#ApiV1BlockVolumesGet) | **Get** /api/v1/block/volumes | 查询Volume列表
[**ApiV1BlockVolumesPost**](VolumesApi.md#ApiV1BlockVolumesPost) | **Post** /api/v1/block/volumes | 创建Volume
[**ApiV1BlockVolumesVolumeIdDelete**](VolumesApi.md#ApiV1BlockVolumesVolumeIdDelete) | **Delete** /api/v1/block/volumes/{volume_id} | 删除卷
[**ApiV1BlockVolumesVolumeIdExpandPut**](VolumesApi.md#ApiV1BlockVolumesVolumeIdExpandPut) | **Put** /api/v1/block/volumes/{volume_id}/expand | 卷扩容
[**ApiV1BlockVolumesVolumeIdFlattenPut**](VolumesApi.md#ApiV1BlockVolumesVolumeIdFlattenPut) | **Put** /api/v1/block/volumes/{volume_id}/flatten | 将链接克隆卷从快照树上断开关系链
[**ApiV1BlockVolumesVolumeIdGet**](VolumesApi.md#ApiV1BlockVolumesVolumeIdGet) | **Get** /api/v1/block/volumes/{volume_id} | 查询Volume详情
[**ApiV1BlockVolumesVolumeIdIoPriorityPut**](VolumesApi.md#ApiV1BlockVolumesVolumeIdIoPriorityPut) | **Put** /api/v1/block/volumes/{volume_id}/io-priority | 更新卷的数据校验
[**ApiV1BlockVolumesVolumeIdNamePut**](VolumesApi.md#ApiV1BlockVolumesVolumeIdNamePut) | **Put** /api/v1/block/volumes/{volume_id}/name | 更新卷名称
[**ApiV1BlockVolumesVolumeIdPut**](VolumesApi.md#ApiV1BlockVolumesVolumeIdPut) | **Put** /api/v1/block/volumes/{volume_id} | 更新卷
[**ApiV1BlockVolumesVolumeIdQosPut**](VolumesApi.md#ApiV1BlockVolumesVolumeIdQosPut) | **Put** /api/v1/block/volumes/{volume_id}/qos | 卷设置QOS
[**ApiV1BlockVolumesVolumeIdShrinkPut**](VolumesApi.md#ApiV1BlockVolumesVolumeIdShrinkPut) | **Put** /api/v1/block/volumes/{volume_id}/shrink | 卷缩容
[**ApiV1BlockVolumesVolumeIdVerifyEnabledPut**](VolumesApi.md#ApiV1BlockVolumesVolumeIdVerifyEnabledPut) | **Put** /api/v1/block/volumes/{volume_id}/verify-enabled | 更新卷的数据校验


# **ApiV1BlockVolumesBatchCreationPost**
> AsyncTaskResponseView ApiV1BlockVolumesBatchCreationPost(ctx, body)
批量创建Volume

批量创建Volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchCreateVolumeRequestView**](BatchCreateVolumeRequestView.md)|  | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesBatchDeletionPost**
> VolumeBatchDeleteResponseView ApiV1BlockVolumesBatchDeletionPost(ctx, body)
批量删除卷

批量删除卷

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VolumeBatchDeleteRequestView**](VolumeBatchDeleteRequestView.md)|  | 

### Return type

[**VolumeBatchDeleteResponseView**](VolumeBatchDeleteResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesGet**
> VolumeListResponseView ApiV1BlockVolumesGet(ctx, optional)
查询Volume列表

查询Volume列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumesApiApiV1BlockVolumesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumesApiApiV1BlockVolumesGetOpts struct

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
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 
 **capHealthStatus** | **optional.String**| 根据资源的容量健康状态进行过滤 | 
 **perfHealthStatus** | **optional.String**| 根据资源性能健康状态进行过滤 | 
 **syncStatus** | **optional.String**| 通过同步状态进行过滤 | 
 **qosStatus** | **optional.Bool**| 通过qos状态进行过滤 | 
 **verifyEnabled** | **optional.Bool**| 通过数据校验进行过滤 | 
 **dataProtection** | **optional.Bool**| 通过数据保护进行过滤 | 
 **cgName** | **optional.String**| 通过一致性组进行过滤 | 
 **cgId** | **optional.String**| 通过一致性组ID进行过滤 | 
 **volumeType** | **optional.String**| 通过卷类型进行过滤 | 
 **ownerPlatform** | **optional.String**| 通过平台进行过滤 | 
 **ioPriority** | **optional.String**| 通过性能优先级进行过滤 | 
 **snapNum** | **optional.Int32**| 通过快照数进行过滤 | 
 **clientNum** | **optional.Int32**| 通过终端数进行过滤 | 

### Return type

[**VolumeListResponseView**](VolumeListResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesPost**
> AsyncTaskResponseView ApiV1BlockVolumesPost(ctx, body)
创建Volume

创建Volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVolumeRequestViews**](CreateVolumeRequestViews.md)|  | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdDelete**
> VolumeDeleteResponseView ApiV1BlockVolumesVolumeIdDelete(ctx, volumeId, optional)
删除卷

删除卷

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| 卷的UUID | 
 **optional** | ***VolumesApiApiV1BlockVolumesVolumeIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumesApiApiV1BlockVolumesVolumeIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| 是否强制执行 | 

### Return type

[**VolumeDeleteResponseView**](VolumeDeleteResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdExpandPut**
> VolumeExpandResponseView ApiV1BlockVolumesVolumeIdExpandPut(ctx, body, volumeId)
卷扩容

卷扩容

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VolumeExpandRequestView**](VolumeExpandRequestView.md)|  | 
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**VolumeExpandResponseView**](VolumeExpandResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdFlattenPut**
> VolumeFlattenResponseView ApiV1BlockVolumesVolumeIdFlattenPut(ctx, volumeId)
将链接克隆卷从快照树上断开关系链

将链接克隆卷从快照树上断开关系链

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**VolumeFlattenResponseView**](VolumeFlattenResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdGet**
> VolumeDetailResponseView ApiV1BlockVolumesVolumeIdGet(ctx, volumeId)
查询Volume详情

查询Volume详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**VolumeDetailResponseView**](VolumeDetailResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdIoPriorityPut**
> UpdateVolumeResponseView ApiV1BlockVolumesVolumeIdIoPriorityPut(ctx, body, volumeId)
更新卷的数据校验

更新卷的数据校验

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVolumeIoPriorityRequestView**](UpdateVolumeIoPriorityRequestView.md)|  | 
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**UpdateVolumeResponseView**](UpdateVolumeResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdNamePut**
> UpdateVolumeResponseView ApiV1BlockVolumesVolumeIdNamePut(ctx, body, volumeId)
更新卷名称

更新卷名称

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVolumeNameRequestView**](UpdateVolumeNameRequestView.md)|  | 
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**UpdateVolumeResponseView**](UpdateVolumeResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdPut**
> UpdateVolumeResponseView ApiV1BlockVolumesVolumeIdPut(ctx, body, volumeId)
更新卷

更新卷

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVolumeRequestView**](UpdateVolumeRequestView.md)|  | 
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**UpdateVolumeResponseView**](UpdateVolumeResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdQosPut**
> SetVolumeQosResponseView ApiV1BlockVolumesVolumeIdQosPut(ctx, body, volumeId)
卷设置QOS

卷设置QOS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SetVolumeQosRequestView**](SetVolumeQosRequestView.md)|  | 
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**SetVolumeQosResponseView**](SetVolumeQOSResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdShrinkPut**
> VolumeShrinkResponseView ApiV1BlockVolumesVolumeIdShrinkPut(ctx, body, volumeId)
卷缩容

卷缩容

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VolumeShrinkRequestView**](VolumeShrinkRequestView.md)|  | 
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**VolumeShrinkResponseView**](VolumeShrinkResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockVolumesVolumeIdVerifyEnabledPut**
> UpdateVolumeResponseView ApiV1BlockVolumesVolumeIdVerifyEnabledPut(ctx, body, volumeId)
更新卷的数据校验

更新卷的数据校验

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVolumeVerifyEnableRequestView**](UpdateVolumeVerifyEnableRequestView.md)|  | 
  **volumeId** | **string**| 卷的UUID | 

### Return type

[**UpdateVolumeResponseView**](UpdateVolumeResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

