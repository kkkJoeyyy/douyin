// Code generated by hertz generator.

package Comment

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/simplecolding/douyin/hertz-server/biz/mw"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _commentMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{mw.JwtMiddleware.MiddlewareFunc()}
}

func _comment_ctionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcommentlistMw() []app.HandlerFunc {
	// your code...
	return nil
}
