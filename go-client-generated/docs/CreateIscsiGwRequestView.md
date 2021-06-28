# CreateIscsiGwRequestView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iqn** | **string** | 自定义iSCSI网关的IQN，必须符合IQN格式，可不指定该参数 | [optional] [default to null]
**ChapUsername** | **string** | iSCSI网关chap用户名(约束：8-64个字符，支持大小写字母数字、下划线、连字符、’@‘、冒号、句点，用户名和密码必须同时存在) | [optional] [default to null]
**GatewayType** | **string** | 块网关的类型，iSCSI或者SCSI | [default to null]
**Port** | **int32** | 网关的端口信息, 目前不支持自定义，只能配置为3260 | [default to null]
**IsChap** | **bool** | iSCSI网关是否启用chap | [default to null]
**Name** | **string** | iSCSI网关名称(约束：1-64个字符，支持小写字母、数字、下划线，只能以字母开头 | [default to null]
**ChapPassword** | **string** | iSCSI网关chap密码(约束:12-16个字符，支持大小写字母、数字、下划线、连字符、‘/’、’@‘) | [optional] [default to null]
**Nodes** | [**[]GatewayAddNodeProperty**](GatewayAddNodeProperty.md) | 网关包含的网关节点 | [default to null]
**Description** | **string** | iSCSI网关描述信息 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


