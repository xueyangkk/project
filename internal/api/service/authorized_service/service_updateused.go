package authorized_service

import (
	"exams-api/internal/api/repository/db_repo"
	"exams-api/internal/api/repository/db_repo/authorized_repo"
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"

	"gorm.io/gorm"
)

func (s *service) UpdateUsed(ctx core.Context, id int32, used int32) (err error) {
	authorizedInfo, err := authorized_repo.NewQueryBuilder().
		WhereIsDeleted(db_repo.EqualPredicate, -1).
		WhereId(db_repo.EqualPredicate, id).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	model := authorized_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(cacheKeyPrefix+authorizedInfo.BusinessKey, cache.WithTrace(ctx.Trace()))
	return
}
