# GatewayLunNumResponseView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | 请求失败时的原因，正常时为空 | [default to null]
**Count** | **int32** | 当前返回符合条件的的条目数 | [optional] [default to null]
**Offset** | **int32** | 期望查询个数 | [optional] [default to null]
**LunNum** | [**[]GatewayLunProperty**](GatewayLunProperty.md) | 网关关联的lun编号 | [optional] [default to null]
**Index** | **int32** | 列表返回的起始位置 | [optional] [default to null]
**RetCode** | **string** | 请求的返回码，正常时为0，失败时为对应错误码，如02.ff.ff.ffff | [default to null]
**Total** | **int32** | 资源符合条件的总条目数 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


