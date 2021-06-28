# CreateVolumeRequestViews

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeSize** | **int32** | 卷的分配容量即卷大小。单位为Bytes, 最小1GB大小 | [default to null]
**Name** | **string** | 卷名称 | [default to null]
**VerifyEnabled** | **bool** | 是否数据校验。 | [default to null]
**IoPriority** | **string** | IO性能优先级。只能是default或priority | [default to null]
**Qos** | [***QosPropertyView**](QOSPropertyView.md) | Volume的QOS配置 | [default to null]
**PoolId** | **string** | 卷所属存储池的uuid | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


