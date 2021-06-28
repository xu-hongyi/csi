# \GatewayApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1BlockIscsiGatewaysClientsGet**](GatewayApi.md#ApiV1BlockIscsiGatewaysClientsGet) | **Get** /api/v1/block/iscsi/gateways/clients | 网关可添加的客户端列表
[**ApiV1BlockIscsiGatewaysGatewayIdChapPut**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdChapPut) | **Put** /api/v1/block/iscsi/gateways/{gateway_id}/chap | 更新块网关CHAP
[**ApiV1BlockIscsiGatewaysGatewayIdClientsGet**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdClientsGet) | **Get** /api/v1/block/iscsi/gateways/{gateway_id}/clients | 网关已关联的客户端列表
[**ApiV1BlockIscsiGatewaysGatewayIdClientsPut**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdClientsPut) | **Put** /api/v1/block/iscsi/gateways/{gateway_id}/clients | 块网关增加/删除客户端
[**ApiV1BlockIscsiGatewaysGatewayIdDelete**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdDelete) | **Delete** /api/v1/block/iscsi/gateways/{gateway_id} | 删除iSCSI网关
[**ApiV1BlockIscsiGatewaysGatewayIdDescPut**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdDescPut) | **Put** /api/v1/block/iscsi/gateways/{gateway_id}/desc | 更新块网关描述
[**ApiV1BlockIscsiGatewaysGatewayIdGet**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdGet) | **Get** /api/v1/block/iscsi/gateways/{gateway_id} | 根据ID查询网关的详情信息
[**ApiV1BlockIscsiGatewaysGatewayIdLunsGet**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdLunsGet) | **Get** /api/v1/block/iscsi/gateways/{gateway_id}/luns | 网关已关联的卷列表
[**ApiV1BlockIscsiGatewaysGatewayIdNamePut**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdNamePut) | **Put** /api/v1/block/iscsi/gateways/{gateway_id}/name | 更新块网关名称
[**ApiV1BlockIscsiGatewaysGatewayIdNodesGet**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdNodesGet) | **Get** /api/v1/block/iscsi/gateways/{gateway_id}/nodes | 网关已添加的节点列表
[**ApiV1BlockIscsiGatewaysGatewayIdNodesPut**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdNodesPut) | **Put** /api/v1/block/iscsi/gateways/{gateway_id}/nodes | 块网关增加/删除服务器节点
[**ApiV1BlockIscsiGatewaysGatewayIdPut**](GatewayApi.md#ApiV1BlockIscsiGatewaysGatewayIdPut) | **Put** /api/v1/block/iscsi/gateways/{gateway_id} | 更新块网关基本信息
[**ApiV1BlockIscsiGatewaysGet**](GatewayApi.md#ApiV1BlockIscsiGatewaysGet) | **Get** /api/v1/block/iscsi/gateways | 查询块网关列表
[**ApiV1BlockIscsiGatewaysNodesGet**](GatewayApi.md#ApiV1BlockIscsiGatewaysNodesGet) | **Get** /api/v1/block/iscsi/gateways/nodes | 网关可添加的网关节点列表
[**ApiV1BlockIscsiGatewaysPost**](GatewayApi.md#ApiV1BlockIscsiGatewaysPost) | **Post** /api/v1/block/iscsi/gateways | 创建iSCSI网关


# **ApiV1BlockIscsiGatewaysClientsGet**
> AvailClientsResponseView ApiV1BlockIscsiGatewaysClientsGet(ctx, optional)
网关可添加的客户端列表

网关可添加的客户端列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GatewayApiApiV1BlockIscsiGatewaysClientsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayApiApiV1BlockIscsiGatewaysClientsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| 网关的ID，若指定网关ID，则获取网关可关联的客户端列表,若不指定ID，则获取所有状态正常的客户端 | 
 **name** | **optional.String**|  | [default to 根据客户端的名称进行过滤, 支持模糊过滤]
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 

### Return type

[**AvailClientsResponseView**](AvailClientsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdChapPut**
> UpdateBlockGatewayResponseView ApiV1BlockIscsiGatewaysGatewayIdChapPut(ctx, body, gatewayId)
更新块网关CHAP

更新块网关CHAP

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateGatewayChapRequestView**](UpdateGatewayChapRequestView.md)|  | 
  **gatewayId** | **string**| 待更新块网关的id | 

### Return type

[**UpdateBlockGatewayResponseView**](UpdateBlockGatewayResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdClientsGet**
> GatewayClientsResponseView ApiV1BlockIscsiGatewaysGatewayIdClientsGet(ctx, gatewayId, optional)
网关已关联的客户端列表

网关已关联的客户端列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**| 待查询块网关的id | 
 **optional** | ***GatewayApiApiV1BlockIscsiGatewaysGatewayIdClientsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayApiApiV1BlockIscsiGatewaysGatewayIdClientsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| 根据名称进行过滤 | 
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 
 **iscsiGwId** | **optional.String**| 通过网关的ID进行过滤 | 
 **isChap** | **optional.Bool**| 通过客户端chap是否开启进行过滤 | 

### Return type

[**GatewayClientsResponseView**](GatewayClientsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdClientsPut**
> UpdateBlockGatewayClientsResponseView ApiV1BlockIscsiGatewaysGatewayIdClientsPut(ctx, body, gatewayId)
块网关增加/删除客户端

块网关增加/删除客户端

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBlockGatewayClientsRequestView**](UpdateBlockGatewayClientsRequestView.md)|  | 
  **gatewayId** | **string**| 待更新块网关的id | 

### Return type

[**UpdateBlockGatewayClientsResponseView**](UpdateBlockGatewayClientsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdDelete**
> AsyncTaskResponseView ApiV1BlockIscsiGatewaysGatewayIdDelete(ctx, gatewayId, optional)
删除iSCSI网关

删除iSCSI网关

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**| iSCSI网关id | 
 **optional** | ***GatewayApiApiV1BlockIscsiGatewaysGatewayIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayApiApiV1BlockIscsiGatewaysGatewayIdDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.String**| 强制删除标记(缺省为false) | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdDescPut**
> UpdateBlockGatewayResponseView ApiV1BlockIscsiGatewaysGatewayIdDescPut(ctx, body, gatewayId)
更新块网关描述

更新块网关描述

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateGatewayDescRequestView**](UpdateGatewayDescRequestView.md)|  | 
  **gatewayId** | **string**| 待更新块网关的id | 

### Return type

[**UpdateBlockGatewayResponseView**](UpdateBlockGatewayResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdGet**
> BlockGwDetailResponseView ApiV1BlockIscsiGatewaysGatewayIdGet(ctx, gatewayId)
根据ID查询网关的详情信息

根据ID查询网关的详情信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**| iSCSI网关id | 

### Return type

[**BlockGwDetailResponseView**](BlockGWDetailResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdLunsGet**
> GatewayLunsResponseView ApiV1BlockIscsiGatewaysGatewayIdLunsGet(ctx, gatewayId, optional)
网关已关联的卷列表

网关已关联的卷列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**| 待查询块网关的id | 
 **optional** | ***GatewayApiApiV1BlockIscsiGatewaysGatewayIdLunsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayApiApiV1BlockIscsiGatewaysGatewayIdLunsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| 根据卷的名称进行过滤，支持模糊过滤 | 
 **volumeName** | **optional.String**| 根据卷底层名称进行过滤，支持模糊过滤 | 
 **description** | **optional.String**| 根据卷描述进行过滤，支持模糊过滤 | 
 **poolName** | **optional.String**| 根据存储池名称进行过滤 | 
 **volumeType** | **optional.String**| 卷类型过滤 | 
 **capHealthStatus** | **optional.String**| 根据资源的容量健康状态进行过滤 | 
 **perfHealthStatus** | **optional.String**| 根据资源性能健康状态进行过滤 | 
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 
 **syncStatus** | **optional.String**| 通过同步状态进行过滤 | 
 **createFrom** | **optional.String**| 根据卷的创建方式进行过滤 | 
 **useStatus** | **optional.String**| 卷使用状态进行过滤 | 
 **isReadonly** | **optional.Bool**| 卷读写方式进行过滤 | 

### Return type

[**GatewayLunsResponseView**](GatewayLunsResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdNamePut**
> UpdateBlockGatewayResponseView ApiV1BlockIscsiGatewaysGatewayIdNamePut(ctx, body, gatewayId)
更新块网关名称

更新块网关名称

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateGatewayNameRequestView**](UpdateGatewayNameRequestView.md)|  | 
  **gatewayId** | **string**| 待更新块网关的id | 

### Return type

[**UpdateBlockGatewayResponseView**](UpdateBlockGatewayResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdNodesGet**
> GatewayNodesResponseView ApiV1BlockIscsiGatewaysGatewayIdNodesGet(ctx, gatewayId, optional)
网关已添加的节点列表

网关已添加的节点列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **gatewayId** | **string**| 待查询块网关的id | 
 **optional** | ***GatewayApiApiV1BlockIscsiGatewaysGatewayIdNodesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayApiApiV1BlockIscsiGatewaysGatewayIdNodesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| 根据名称进行过滤 | 
 **nodeId** | **optional.String**| 根据节点ID进行过滤查询 | 
 **serverId** | **optional.String**| 根据服务器ID进行过滤查询 | 
 **iscsiGwId** | **optional.String**| 根据网关ID进行过滤查询 | 
 **managerIp** | **optional.String**| 根据节点的管理IP进行过滤查询 | 
 **gatewayIp** | **optional.String**| 根据节点的网关IP进行过滤查询 | 
 **healthStatus** | **optional.String**| 根据资源的健康状态进行过滤 | 

### Return type

[**GatewayNodesResponseView**](GatewayNodesResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdNodesPut**
> UpdateBlockGatewayNodesResponseView ApiV1BlockIscsiGatewaysGatewayIdNodesPut(ctx, body, gatewayId)
块网关增加/删除服务器节点

块网关增加/删除服务器节点

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBlockGatewayNodesRequestView**](UpdateBlockGatewayNodesRequestView.md)|  | 
  **gatewayId** | **string**| 待更新块网关的id | 

### Return type

[**UpdateBlockGatewayNodesResponseView**](UpdateBlockGatewayNodesResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGatewayIdPut**
> UpdateBlockGatewayResponseView ApiV1BlockIscsiGatewaysGatewayIdPut(ctx, body, gatewayId)
更新块网关基本信息

更新块网关基本信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBlockGatewayRequestView**](UpdateBlockGatewayRequestView.md)|  | 
  **gatewayId** | **string**| 待更新块网关的id | 

### Return type

[**UpdateBlockGatewayResponseView**](UpdateBlockGatewayResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysGet**
> BlockGwListResponseView ApiV1BlockIscsiGatewaysGet(ctx, optional)
查询块网关列表

查询块网关列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GatewayApiApiV1BlockIscsiGatewaysGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayApiApiV1BlockIscsiGatewaysGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| 根据ID进行过滤 | 
 **name** | **optional.String**| 根据名称进行过滤 | 
 **iqn** | **optional.String**| 根据网关的IQN进行过滤 | 
 **port** | **optional.Int32**| 根据网关的端口号进行过滤 | 
 **runStatus** | **optional.String**| 根据资源的运行状态进行过滤 | 
 **healthStatus** | **optional.String**| 根据资源的健康状态进行过滤 | 
 **isEnabled** | **optional.Bool**| 根据网关的启用状态进行过滤 | 
 **isChap** | **optional.Bool**| 根据网关的CHAP启用状态进行过滤 | 
 **chapUsername** | **optional.String**| 根据网关的CHAP用户名进行过滤 | 
 **description** | **optional.String**| 根据网关的描述信息进行过滤 | 

### Return type

[**BlockGwListResponseView**](BlockGWListResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysNodesGet**
> AvailNodesResponseView ApiV1BlockIscsiGatewaysNodesGet(ctx, optional)
网关可添加的网关节点列表

网关可添加的网关节点列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GatewayApiApiV1BlockIscsiGatewaysNodesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GatewayApiApiV1BlockIscsiGatewaysNodesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**| 网关的ID，若指定网关ID，则获取网关可添加的网关节点列表,若不指定ID，则获取所有状态正常的块网关节点 | 
 **name** | **optional.String**| 根据网关节点的名称进行过滤，支持模糊过滤 | 
 **managerIp** | **optional.String**| 根据网关节点的管理IP进行过滤，支持模糊过滤 | 
 **gatewayIp** | **optional.String**| 根据网关节点的网关IP进行过滤，支持模糊过滤 | 

### Return type

[**AvailNodesResponseView**](AvailNodesResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV1BlockIscsiGatewaysPost**
> AsyncTaskResponseView ApiV1BlockIscsiGatewaysPost(ctx, body)
创建iSCSI网关

创建iSCSI网关

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateIscsiGwRequestView**](CreateIscsiGwRequestView.md)|  | 

### Return type

[**AsyncTaskResponseView**](AsyncTaskResponseView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

