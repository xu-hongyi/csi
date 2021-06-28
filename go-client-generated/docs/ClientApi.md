# \ClientApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BlockIscsiClientsClientIdChapPut**](ClientApi.md#ApiV1BlockIscsiClientsClientIdChapPut) | **Put** /api/v1/block/iscsi/clients/{client_id}/chap | 更新客户端CHAP配置
[**ApiV1BlockIscsiClientsClientIdDelete**](ClientApi.md#ApiV1BlockIscsiClientsClientIdDelete) | **Delete** /api/v1/block/iscsi/clients/{client_id} | 删除客户端
[**ApiV1BlockIscsiClientsClientIdDescPut**](ClientApi.md#ApiV1BlockIscsiClientsClientIdDescPut) | **Put** /api/v1/block/iscsi/clients/{client_id}/desc | 更新客户端的描述
[**ApiV1BlockIscsiClientsClientIdGatewaysGet**](ClientApi.md#ApiV1BlockIscsiClientsClientIdGatewaysGet) | **Get** /api/v1/block/iscsi/clients/{client_id}/gateways | 查询客户端关联的网关列表
[**ApiV1BlockIscsiClientsClientIdGet**](ClientApi.md#ApiV1BlockIscsiClientsClientIdGet) | **Get** /api/v1/block/iscsi/clients/{client_id} | 查询客户端详情
[**ApiV1BlockIscsiClientsClientIdHostsGet**](ClientApi.md#ApiV1BlockIscsiClientsClientIdHostsGet) | **Get** /api/v1/block/iscsi/clients/{client_id}/hosts | 查询客户端添加的客户端主机
[**ApiV1BlockIscsiClientsClientIdHostsPut**](ClientApi.md#ApiV1BlockIscsiClientsClientIdHostsPut) | **Put** /api/v1/block/iscsi/clients/{client_id}/hosts | 网关添加/删除客户端主机
[**ApiV1BlockIscsiClientsClientIdLunsGet**](ClientApi.md#ApiV1BlockIscsiClientsClientIdLunsGet) | **Get** /api/v1/block/iscsi/clients/{client_id}/luns | 查询客户端挂载的卷设备
[**ApiV1BlockIscsiClientsClientIdLunsPut**](ClientApi.md#ApiV1BlockIscsiClientsClientIdLunsPut) | **Put** /api/v1/block/iscsi/clients/{client_id}/luns | 网关挂载/卸载卷
[**ApiV1BlockIscsiClientsClientIdPut**](ClientApi.md#ApiV1BlockIscsiClientsClientIdPut) | **Put** /api/v1/block/iscsi/clients/{client_id} | 更新客户端
[**ApiV1BlockIscsiClientsGet**](ClientApi.md#ApiV1BlockIscsiClientsGet) | **Get** /api/v1/block/iscsi/clients | 获取客户端列表
[**ApiV1BlockIscsiClientsLunsGet**](ClientApi.md#ApiV1BlockIscsiClientsLunsGet) | **Get** /api/v1/block/iscsi/clients/luns | 查询客户端可挂载的卷设备
[**ApiV1BlockIscsiClientsPost**](ClientApi.md#ApiV1BlockIscsiClientsPost) | **Post** /api/v1/block/iscsi/clients | 创建iSCSI客户端


# **ApiV1BlockIscsiClientsClientIdChapPut**
> ClientUpdateResponseView ApiV1BlockIscsiClientsClientIdChapPut(ctx, body, clientId)
更新客户端CHAP配置

更新客户端CHAP配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateClientChapRequestView**](UpdateClientChapRequestView.md)|  | 
  **clientId** | **string**| 待更新客户端的id | 

### Return type

[**ClientUpdateResponseView**](ClientUpdateResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdDelete**
> ClientDeleteResponseView ApiV1BlockIscsiClientsClientIdDelete(ctx, clientId, optional)
删除客户端

删除客户端

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| 待删除客户端的id | 
 **optional** | ***ClientApiApiV1BlockIscsiClientsClientIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientApiApiV1BlockIscsiClientsClientIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| 是否强制删除 | 

### Return type

[**ClientDeleteResponseView**](ClientDeleteResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdDescPut**
> ClientUpdateResponseView ApiV1BlockIscsiClientsClientIdDescPut(ctx, body, clientId)
更新客户端的描述

更新客户端的描述

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateClientDescRequestView**](UpdateClientDescRequestView.md)|  | 
  **clientId** | **string**| 待更新客户端的id | 

### Return type

[**ClientUpdateResponseView**](ClientUpdateResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdGatewaysGet**
> ClientRelatedGatewayResponseView ApiV1BlockIscsiClientsClientIdGatewaysGet(ctx, clientId, optional)
查询客户端关联的网关列表

查询客户端关联的网关列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| 待查询客户端的id | 
 **optional** | ***ClientApiApiV1BlockIscsiClientsClientIdGatewaysGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientApiApiV1BlockIscsiClientsClientIdGatewaysGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| 根据网关的名称进行过滤，支持模糊过滤 | 
 **iqn** | **optional.String**| 根据网关的IQN进行过滤，支持模糊过滤 | 
 **port** | **optional.Int32**| 根据网关的端口号进行过滤 | 
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 
 **description** | **optional.String**| 根据描述进行过滤 | 
 **isChap** | **optional.Bool**| 根据网关是否开启CHAP进行过滤 | 
 **isEnabled** | **optional.Bool**| 根据网关是否启用进行过滤 | 

### Return type

[**ClientRelatedGatewayResponseView**](ClientRelatedGatewayResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdGet**
> ClientDetailResponseView ApiV1BlockIscsiClientsClientIdGet(ctx, clientId)
查询客户端详情

查询客户端详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| 待查询客户端的id | 

### Return type

[**ClientDetailResponseView**](ClientDetailResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdHostsGet**
> ClientBelongsHostsResponseView ApiV1BlockIscsiClientsClientIdHostsGet(ctx, clientId, optional)
查询客户端添加的客户端主机

查询客户端添加的客户端主机

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| 待查询客户端的id | 
 **optional** | ***ClientApiApiV1BlockIscsiClientsClientIdHostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientApiApiV1BlockIscsiClientsClientIdHostsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **node** | **optional.String**| 根据主机的IQN进行过滤 | 
 **useStatus** | **optional.String**| 根据主机的登录状态过滤 | 

### Return type

[**ClientBelongsHostsResponseView**](ClientBelongsHostsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdHostsPut**
> ClientUpdateHostsResponseView ApiV1BlockIscsiClientsClientIdHostsPut(ctx, body, clientId)
网关添加/删除客户端主机

网关添加/删除客户端主机

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientUpdateHostsRequestView**](ClientUpdateHostsRequestView.md)|  | 
  **clientId** | **string**| 待更新客户端的id | 

### Return type

[**ClientUpdateHostsResponseView**](ClientUpdateHostsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdLunsGet**
> ClientBelongsLunsResponseView ApiV1BlockIscsiClientsClientIdLunsGet(ctx, clientId, optional)
查询客户端挂载的卷设备

查询客户端挂载的卷设备

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| 待查询客户端的id | 
 **optional** | ***ClientApiApiV1BlockIscsiClientsClientIdLunsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientApiApiV1BlockIscsiClientsClientIdLunsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| 根据卷的名称进行过滤，支持模糊过滤 | 
 **volumeName** | **optional.String**| 根据卷底层名称进行过滤，支持模糊过滤 | 
 **description** | **optional.String**| 根据卷描述进行过滤，支持模糊过滤 | 
 **volumeType** | **optional.String**| 卷类型过滤 | 
 **poolName** | **optional.String**| 根据存储池名称进行过滤 | 
 **capHealthStatus** | **optional.String**| 根据资源的容量健康状态进行过滤 | 
 **perfHealthStatus** | **optional.String**| 根据资源性能健康状态进行过滤 | 
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 
 **syncStatus** | **optional.String**| 通过同步状态进行过滤 | 
 **createFrom** | **optional.String**| 根据卷的创建方式进行过滤 | 
 **useStatus** | **optional.String**| 卷使用状态进行过滤 | 
 **isReadonly** | **optional.Bool**| 卷读写方式进行过滤 | 

### Return type

[**ClientBelongsLunsResponseView**](ClientBelongsLunsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdLunsPut**
> ClientUpdateLunsResponseView ApiV1BlockIscsiClientsClientIdLunsPut(ctx, body, clientId)
网关挂载/卸载卷

网关挂载/卸载卷

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientUpdateLunsRequestView**](ClientUpdateLunsRequestView.md)|  | 
  **clientId** | **string**| 待更新客户端的id | 

### Return type

[**ClientUpdateLunsResponseView**](ClientUpdateLunsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsClientIdPut**
> ClientUpdateResponseView ApiV1BlockIscsiClientsClientIdPut(ctx, body, clientId)
更新客户端

更新客户端

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClientUpdateRequestView**](ClientUpdateRequestView.md)|  | 
  **clientId** | **string**| 待更新客户端的id | 

### Return type

[**ClientUpdateResponseView**](ClientUpdateResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsGet**
> ClientListResponseView ApiV1BlockIscsiClientsGet(ctx, optional)
获取客户端列表

获取客户端列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientApiApiV1BlockIscsiClientsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientApiApiV1BlockIscsiClientsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| 根据客户端客户端名称进行查询，支持模块搜索 | 
 **id** | **optional.String**| 根据客户端客户端ID进行查询 | 
 **description** | **optional.String**| 根据客户端客户端的描述信息系进行查询，支持模块搜索 | 
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 
 **isChap** | **optional.Bool**| 根据客户端客户端是否开启CHAP进行查询 | 

### Return type

[**ClientListResponseView**](ClientListResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsLunsGet**
> ClientAvailLunsResponseView ApiV1BlockIscsiClientsLunsGet(ctx, optional)
查询客户端可挂载的卷设备

查询客户端可挂载的卷设备

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ClientApiApiV1BlockIscsiClientsLunsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ClientApiApiV1BlockIscsiClientsLunsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| 根据卷的名称进行查询，支持模糊匹配 | 
 **poolName** | **optional.String**| 根据存储池名称进行查询，支持模糊匹配 | 
 **lunType** | **optional.String**| 根据卷类型进行过滤 | 

### Return type

[**ClientAvailLunsResponseView**](ClientAvailLunsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiClientsPost**
> CreateClientResponseView ApiV1BlockIscsiClientsPost(ctx, body)
创建iSCSI客户端

创建iSCSI客户端

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateIscsiClientRequestView**](CreateIscsiClientRequestView.md)|  | 

### Return type

[**CreateClientResponseView**](CreateClientResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

