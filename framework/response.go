package framework

type IResponse interface {
	//Json输出
	Json(obj interface{}) IResponse

	// Jsonp输出
	Jsonp(obj interface{}) IResponse

	// Xlm输出
	Xml(obj interface{}) IResponse

	// Html输出
	Html(template string, obj interface{}) IResponse

	// String
	Text(format string, val ...interface{}) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置200状态
	SetOkStatus() IResponse
}

func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (ctx *Context) Text(status int, obj string) error {
	return nil
}
