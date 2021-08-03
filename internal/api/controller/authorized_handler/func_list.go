package authorized_handler

import (
	"net/http"

	"exams-api/internal/api/code"
	"exams-api/internal/api/service/authorized_service"
	"exams-api/internal/pkg/core"
	"exams-api/pkg/errno"
	"exams-api/pkg/time_parse"

	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type listRequest struct {
	Page              int    `form:"page"`               // 第几页
	PageSize          int    `form:"page_size"`          // 每页显示条数
	BusinessKey       string `form:"business_key"`       // 调用方key
	BusinessSecret    string `form:"business_secret"`    // 调用方secret
	BusinessDeveloper string `form:"business_developer"` // 调用方对接人
	Remark            string `form:"remark"`             // 备注
}

type listData struct {
	Id                int    `json:"id"`                 // ID
	HashID            string `json:"hashid"`             // hashid
	BusinessKey       string `json:"business_key"`       // 调用方key
	BusinessSecret    string `json:"business_secret"`    // 调用方secret
	BusinessDeveloper string `json:"business_developer"` // 调用方对接人
	Remark            string `json:"remark"`             // 备注
	IsUsed            int    `json:"is_used"`            // 是否启用 1:是 -1:否
	CreatedAt         string `json:"created_at"`         // 创建时间
	CreatedUser       string `json:"created_user"`       // 创建人
	UpdatedAt         string `json:"updated_at"`         // 更新时间
	UpdatedUser       string `json:"updated_user"`       // 更新人
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PrePageCount int `json:"pre_page_count"`
	} `json:"pagination"`
}

// List 调用方列表
// @Summary 调用方列表
// @Description 调用方列表
// @Tags API.authorized
// @Accept json
// @Produce json
// @Param page query int false "第几页"
// @Param page_size query string false "每页显示条数"
// @Param business_key query string false "调用方key"
// @Param business_secret query string false "调用方secret"
// @Param business_developer query string false "调用方对接人"
// @Param remark path string false "备注"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}

		searchData := new(authorized_service.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.BusinessKey = req.BusinessKey
		searchData.BusinessSecret = req.BusinessSecret
		searchData.Remark = req.Remark

		resListData, err := h.authorizedService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedListError,
				code.Text(code.AuthorizedListError)).WithErr(err),
			)
			return
		}

		resCountData, err := h.authorizedService.PageListCount(c, searchData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AuthorizedListError,
				code.Text(code.AuthorizedListError)).WithErr(err),
			)
			return
		}
		res.Pagination.Total = cast.ToInt(resCountData)
		res.Pagination.PrePageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				h.logger.Info("hashids err", zap.Error(err))
			}

			data := listData{
				Id:                cast.ToInt(v.Id),
				HashID:            hashId,
				BusinessKey:       v.BusinessKey,
				BusinessSecret:    v.BusinessSecret,
				BusinessDeveloper: v.BusinessDeveloper,
				Remark:            v.Remark,
				IsUsed:            cast.ToInt(v.IsUsed),
				CreatedAt:         v.CreatedAt.Format(time_parse.CSTLayout),
				CreatedUser:       v.CreatedUser,
				UpdatedAt:         v.UpdatedAt.Format(time_parse.CSTLayout),
				UpdatedUser:       v.UpdatedUser,
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
