# BatchCreateVolumeRequestView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeSize** | **int32** | 卷的分配容量即卷大小。单位为Bytes, 最小1GB大小 | [default to null]
**StartNo** | **int32** | 批量创建块存储卷时，卷的起始编号。最小为1 | [default to null]
**VerifyEnabled** | **bool** | 是否数据校验。 | [default to null]
**IoPriority** | **string** | IO性能优先级。只能是default或priority | [default to null]
**Qos** | [***QosPropertyView**](QOSPropertyView.md) | Volume的QOS配置 | [default to null]
**PoolId** | **string** | 卷所属存储池的uuid | [default to null]
**Prefix** | **string** | 卷名称前缀 | [default to null]
**Count** | **int32** | 批量创建的卷个数 | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


