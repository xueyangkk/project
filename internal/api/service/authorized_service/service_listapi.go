package authorized_service

import (
	"exams-api/internal/api/repository/db_repo"
	"exams-api/internal/api/repository/db_repo/authorized_api_repo"
	"exams-api/internal/pkg/core"
)

type SearchAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
}

func (s *service) ListAPI(ctx core.Context, searchAPIData *SearchAPIData) (listData []*authorized_api_repo.AuthorizedApi, err error) {

	qb := authorized_api_repo.NewQueryBuilder()
	qb = qb.WhereIsDeleted(db_repo.EqualPredicate, -1)

	if searchAPIData.BusinessKey != "" {
		qb.WhereBusinessKey(db_repo.EqualPredicate, searchAPIData.BusinessKey)
	}

	listData, err = qb.
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
