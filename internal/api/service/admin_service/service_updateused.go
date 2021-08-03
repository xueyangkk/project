package admin_service

import (
	"exams-api/internal/api/repository/db_repo/admin_repo"
	"exams-api/internal/pkg/cache"
	"exams-api/internal/pkg/core"
	"exams-api/internal/pkg/password"
)

func (s *service) UpdateUsed(ctx core.Context, id int32, used int32) (err error) {
	model := admin_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"is_used":      used,
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(cacheKeyPrefix+password.GenerateLoginToken(id), cache.WithTrace(ctx.Trace()))
	return
}
