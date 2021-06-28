# VolumeDetailResponseView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataProtection** | **bool** | 卷是否开启数据保护 | [default to null]
**ParentSnapId** | **string** | 父快照的ID，仅用于克隆卷 | [default to null]
**PerfHealthStatusReason** | **string** | 资源的性能健康状态非健康的原因 | [default to null]
**CgId** | **string** | 卷所属一致性组的ID | [default to null]
**CreateTime** | **float32** | 资源的创建时间 | [default to null]
**PoolName** | **string** | 卷所在存储池的名称 | [default to null]
**ClientNum** | **int32** | 卷关联的终端数量 | [default to null]
**RetCode** | **string** | 请求的返回码，正常时为0，失败时为对应错误码，如02.ff.ff.ffff | [default to null]
**CreateFrom** | **string** | 资源的创建方式 | [default to null]
**CapHealthStatus** | **string** | 资源的容量健康状态 | [default to null]
**VerifyEnabled** | **bool** | 卷是否开启数据校验 | [default to null]
**RunStatus** | **string** | 资源的运行状态 | [default to null]
**PoolId** | **string** | 卷所在存储池的ID | [default to null]
**SyncStatus** | **string** | 资源的同步状态 | [default to null]
**PerfHealthStatus** | **string** | 资源的性能健康状态 | [default to null]
**CapHealthStatusReason** | **string** | 资源的容量健康状态非健康的原因 | [default to null]
**DataSize** | **float32** | 卷已使用数据量，单位为字节 | [default to null]
**Message** | **string** | 请求失败时的原因，正常时为空 | [default to null]
**UpdateTime** | **float32** | 资源的最后时间 | [default to null]
**OwnerPlatform** | **string** | 卷所属对接平台 | [default to null]
**QosStatus** | **bool** | 卷的QOS状态 | [default to null]
**SnapNum** | **int32** | 卷创建的快照数量 | [default to null]
**VolumeSize** | **float32** | 卷的大小，单位为字节 | [default to null]
**Name** | **string** | 资源的名称 | [default to null]
**VolumeType** | **string** | Volume的类型 | [default to null]
**IoPriority** | **string** | 卷的IO性能优先级 | [default to null]
**ParentSnapName** | **string** | 父快照的名称，仅用于克隆卷 | [default to null]
**CgName** | **string** | 卷所属一致性组的名称 | [default to null]
**Description** | **string** | 资源的描述 | [default to null]
**Id** | **string** | 资源的ID | [default to null]
**Qos** | [***QosPropertyView**](QOSPropertyView.md) | 卷的QOS配置 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


