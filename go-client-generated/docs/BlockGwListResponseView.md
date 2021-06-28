# BlockGwListResponseView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | 请求失败时的原因，正常时为空 | [default to null]
**RetCode** | **string** | 请求的返回码，正常时为0，失败时为对应错误码，如02.ff.ff.ffff | [default to null]
**Total** | **int32** | 资源符合条件的总条目数 | [optional] [default to null]
**Gateways** | [**[]BlockGatewayDetailProperty**](BlockGatewayDetailProperty.md) | 查询到的网关列表 | [optional] [default to null]
**Count** | **int32** | 当前返回符合条件的的条目数 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


