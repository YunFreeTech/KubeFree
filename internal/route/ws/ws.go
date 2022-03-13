package ws

//import (
//	shell "github.com/YunFreeTech/kubefree/internal/terminal"
//	"github.com/kataras/iris/v12"
//	"github.com/kataras/iris/v12/context"
//)
//
//func AddWebSocketRoute(app iris.Party) {
//	wsParty := app.Party("/ws")
//	h := shell.CreateAttachHandler("/ws/sockjs")
//	wsParty.Any("/sockjs/{p:path}", func(ctx *context.Context) {
//		h.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
//	})
//}
