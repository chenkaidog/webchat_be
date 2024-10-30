package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"webchat_be/biz/dao"
	"webchat_be/biz/model/consts"
	"webchat_be/biz/model/dto"
	"webchat_be/biz/model/errs"
)

// GetModels 获取模型列表
//
//	@Tags			model
//	@Summary		获取模型列表
//	@Description	获取模型列表
//	@Accept			json
//	@Produce		json
//	@Param			X-CSRF-TOKEN	header		string	true	"csrf token"
//	@Param			cookie			header		string	true	"cookie"
//	@Success		200				{object}	dto.CommonResp{data=dto.ModelQueryResp}
//	@Header			200				{string}	set-cookie	"cookie"
//	@Failure		400,500			{object}	dto.CommonResp
//	@Router			/api/v1/account/models [GET]
func GetModels(ctx context.Context, c *app.RequestContext) {
	accountId := ctx.Value(consts.SessionKeyAccountId).(string)

	modelInfoList, err := dao.NewModelDao().QueryByAccountId(ctx, accountId)
	if err != nil {
		dto.FailResp(c, errs.ServerError)
		return
	}

	var resp dto.ModelQueryResp
	for _, modelInfo := range modelInfoList {
		resp.Models = append(resp.Models, &dto.Model{
			ModelId:   modelInfo.ModelId,
			ModelName: modelInfo.DisplayName,
		})
	}
	dto.SuccessResp(c, resp)
}
