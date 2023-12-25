package router

import (
	"context"
	"fmt"
	"reflect"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"

	"fuya-ark/internal/controller"
	"fuya-ark/utility"
)

type Handler struct {
	reqType reflect.Type
	handler HandleFunc
}

var i18n *gi18n.Manager

func init() {
	if i18n == nil {
		i18n = gi18n.New()
	}
}

func BindRequest(reqModel controller.RequestParam) *Handler {
	h := &Handler{}
	if reqModel == nil {
		return nil
	}
	h.reqType = reflect.TypeOf(reqModel)
	return h
}

type HandleFunc func(ctx context.Context, param controller.RequestParam) (resp interface{}, er gcode.Code)

func (h *Handler) Do(handler HandleFunc) func(r *ghttp.Request) {
	if handler == nil {
		return nil
	}
	h.handler = handler

	return func(r *ghttp.Request) { // 参数解析
		reqValue := reflect.New(h.reqType)
		reqParamPointer, ok := reqValue.Interface().(controller.RequestParam)
		if !ok {
			return
		}

		if err := r.Parse(reqParamPointer); err != nil {
			g.Log().Debug(r.GetCtx(), fmt.Sprintf("参数校验异常 err:%v", err))

			r.Response.WriteJson(ghttp.DefaultHandlerResponse{Code: utility.ParamError.Code(), Message: h.Translate(r.GetCtx(), utility.ParamError)})
			return
		}
		data, myCode := h.handler(r.GetCtx(), reqParamPointer)
		if myCode == nil {
			r.Response.WriteJson(ghttp.DefaultHandlerResponse{Code: utility.Succeed.Code(), Message: h.Translate(r.GetCtx(), utility.Succeed)})
			return
		}

		r.Response.WriteJson(ghttp.DefaultHandlerResponse{Code: myCode.Code(), Message: h.Translate(r.GetCtx(), myCode), Data: data})
	}
}

func (h *Handler) DoGame(handler HandleFunc) func(r *ghttp.Request) {
	if handler == nil {
		return nil
	}
	h.handler = handler
	return func(r *ghttp.Request) { // 参数解析
		reqValue := reflect.New(h.reqType)
		reqParamPointer, ok := reqValue.Interface().(controller.RequestParam)
		if !ok {
			return
		}

		if err := r.Parse(reqParamPointer); err != nil {
			g.Log().Debug(r.GetCtx(), fmt.Sprintf("参数校验异常 err:%v", err))
			r.Response.WriteJson(
				g.Map{"dm_error": utility.ParamError.Code(), "error_msg": h.Translate(r.GetCtx(), utility.ParamError), "data": ""})
			return
		}
		data, myCode := h.handler(r.GetCtx(), reqParamPointer)
		if myCode == nil {
			r.Response.WriteJson(
				g.Map{"dm_error": utility.Succeed.Code(), "error_msg": h.Translate(r.GetCtx(), utility.Succeed), "data": ""})
			return
		}
		r.Response.WriteJson(g.Map{"dm_error": myCode.Code(), "error_msg": h.Translate(r.GetCtx(), myCode), "data": data})
	}
}

func (h *Handler) Translate(ctx context.Context, code gcode.Code) string {
	//reqData := common.Context.GetRequestData(ctx)
	//var header *model.ReqHeader
	//if reqData != nil {
	//	header = reqData.Header
	//}
	//language := common.LanguageEN //默认为英文
	//if header != nil {
	//	clientLanguage := header.ClientLanguage
	//	if clientLanguage == common.LanguageEN || clientLanguage == common.LanguageID || clientLanguage == common.LanguageMS {
	//		language = header.ClientLanguage
	//	}
	//}
	//translateStr := i18n.Translate(gi18n.WithLanguage(ctx, language), cast.ToString(code.Code()))
	//if cast.ToString(code.Code()) == translateStr {
	//	return code.Message()
	//}
	return ""
}
