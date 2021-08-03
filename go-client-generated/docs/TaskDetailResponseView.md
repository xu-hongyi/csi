# TaskDetailResponseView

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | 异步任务的编号 | [optional] [default to null]
**Status** | **string** | 任务结果 | [optional] [default to null]
**Description** | **string** | 任务信息 | [optional] [default to null]
**EndTime** | **float32** | 资源的最后时间 | [optional] [default to null]
**Message** | **string** | 请求失败时的原因，正常时为空 | [default to null]
**Private** | **string** |  | [optional] [default to null]
**Operand** | **string** | 任务操作对象 | [optional] [default to null]
**RetCode** | **string** | 请求的返回码，正常时为0，失败时为对应错误码，如02.ff.ff.ffff | [default to null]
**CurStep** | **float32** | 任务的当前步骤 | [optional] [default to null]
**Steps** | **[]string** | 任务的步骤 | [optional] [default to null]
**Name** | **string** | 资源的名称 | [default to null]
**Progress** | **float32** | 任务的进度 | [optional] [default to null]
**Id** | **string** | 资源的ID | [default to null]
**BeginTime** | **float32** | 资源的开始时间 | [optional] [default to null]
**RetMsg** | **string** | 任务结果信息 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


