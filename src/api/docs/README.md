API 接口约定
================

## 线上域名

`http://anglea.ppapi.com`

## http返回字段格式

*正常情况：http code为200*

```json5
{
  // 返回数据
  "resp_data": {},
  // 正常情况下一定是true
  "succeeded": true
}
```

*异常情况：http code为非200*

```json5
{
  // 异常状态码
  "code": 500,
  // 异常信息
  "msg": "",
  // 一定是null
  "resp_data": null,
  // 一定是false
  "succeeded": false
}
```

*鉴权失败：请求头加上token:$token，验证失败返回http code 401*

```json
{
  "code": 401,
  "desc": "缺少token",
  "msg": "缺少token",
  "resp_data": null,
  "succeeded": false
}
```

## 一般异常处理。非200可以Toast Message处理

## http状态码

```go
package apix

const (
	OK                 = 200 // 没有错误
	InvalidArgument    = 400 // 客户端指定了无效的参数。 检查错误消息和错误详细信息以获取更多信息。
	FailedPrecondition = 400 // 请求不能在当前系统状态下执行，例如删除非空目录
	OutOfRange         = 400 // 客户端指定了无效的范围。
	Unauthenticated    = 401 // 由于遗失，无效或过期的OAuth令牌而导致请求未通过身份验证。
	PermissionDenied   = 403 // 客户端没有足够的权限。这可能是因为OAuth令牌没有正确的范围，客户端没有权限，或者客户端项目尚未启用API。
	NotFound           = 404 // 找不到指定的资源，或者该请求被未公开的原因（例如白名单）拒绝。
	Aborted            = 409 // 并发冲突，例如读-修改-写冲突。
	AlreadyExists      = 409 // 客户端尝试创建的资源已存在。
	ResourceExhausted  = 429 // 资源配额达到速率限制。 客户端应该查找google.rpc.QuotaFailure错误详细信息以获取更多信息。
	Cancelled          = 499 // 客户端取消请求
	DataLoss           = 500 // 不可恢复的数据丢失或数据损坏。 客户端应该向用户报告错误。
	Unknown            = 500 // 未知的服务器错误。 通常是服务器错误。
	Internal           = 500 // 内部服务错误。 通常是服务器错误。
	NotImplemented     = 501 // 服务器未实现该API方法。
	Unavailable        = 503 // 暂停服务。通常是服务器已经关闭。
	DeadlineExceeded   = 504 // 已超过请求期限。如果重复发生，请考虑降低请求的复杂性。
)
```

## 接口列表

[小程序接口文档](/docs/app.md)

[后台接口文档](/docs/admin.md)