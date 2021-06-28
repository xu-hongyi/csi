# CreateIscsiClientRequestView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChapUsername** | **string** | iSCSI网关chap用户名(约束：8-64个字符，支持大小写字母数字、下划线、连字符、’@‘、冒号、句点，用户名和密码必须同时存在) | [optional] [default to null]
**Hosts** | **[]string** | 客户端包含的客户端主机列表 | [default to null]
**IsChap** | **bool** | 客户端是否开启CHAP | [default to null]
**Name** | **string** | 客户端的名称 | [default to null]
**ChapPassword** | **string** | iSCSI网关chap密码(约束:12-16个字符，支持大小写字母、数字、下划线、连字符、‘/’、’@‘) | [optional] [default to null]
**Description** | **string** | iSCSI客户端描述信息（支持utf-8字符,长度256） | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


