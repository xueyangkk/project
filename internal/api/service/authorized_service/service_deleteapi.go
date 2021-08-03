package authorized_service

import (
	"exams-api/internal/api/repository/db_repo"
	"exams-api/internal/api/repository/db_repo/authorized_api_repo"
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"

	"gorm.io/gorm"
)

func (s *service) DeleteAPI(ctx core.Context, id int32) (err error) {
	// 先查询 id 是否存在
	authorizedApiInfo, err := authorized_api_repo.NewQueryBuilder().
		WhereIsDeleted(db_repo.EqualPredicate, -1).
		WhereId(db_repo.EqualPredicate, id).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	model := authorized_api_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"is_deleted":   1,
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(cacheKeyPrefix+authorizedApiInfo.BusinessKey, cache.WithTrace(ctx.Trace()))
	return
}
