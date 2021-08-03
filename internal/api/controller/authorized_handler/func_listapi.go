package authorized_handler

import (
	"net/http"

	"exams-api/internal/api/code"
	"exams-api/internal/api/service/authorized_service"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/errno"

	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type listAPIRequest struct {
	Id string `form:"id"` // hashID
}

type listAPIData struct {
	HashId      string `json:"hash_id"`      // hashID
	BusinessKey string `json:"business_key"` // 调用方key
	Method      string `json:"method"`       // 调用方secret
	API         string `json:"api"`          // 调用方对接人
}

type listAPIResponse struct {
	List []listAPIData `json:"list"`
}

// ListAPI 调用方接口地址列表
// @Summary 调用方接口地址列表
// @Description 调用方接口地址列表
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param business_key query string false "调用方key"
// @Success 200 {object} listAPIResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized_api [get]
func (h *handler) ListAPI() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listAPIRequest)
		res := new(listAPIResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		id := int32(ids[0])

		// 通过 id 查询出 business_key
		authorizedInfo, err := h.authorizedService.Detail(c, id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedDetailError,
				code.Text(code.AuthorizedDetailError)).WithErr(err),
			)
			return
		}

		searchAPIData := new(authorized_service.SearchAPIData)
		searchAPIData.BusinessKey = authorizedInfo.BusinessKey

		resListData, err := h.authorizedService.ListAPI(c, searchAPIData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedListAPIError,
				code.Text(code.AuthorizedListAPIError)).WithErr(err),
			)
			return
		}

		res.List = make([]listAPIData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				h.logger.Info("hashids err", zap.Error(err))
			}

			data := listAPIData{
				HashId:      hashId,
				BusinessKey: v.BusinessKey,
				Method:      v.Method,
				API:         v.Api,
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
