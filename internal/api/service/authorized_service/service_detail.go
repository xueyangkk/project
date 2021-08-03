package authorized_service

import (
	"exams-api/internal/api/repository/db_repo"
	"exams-api/internal/api/repository/db_repo/authorized_repo"
	"exams-api/internal/pkg/core"
)

func (s *service) Detail(ctx core.Context, id int32) (info *authorized_repo.Authorized, err error) {
	qb := authorized_repo.NewQueryBuilder()
	qb.WhereIsDeleted(db_repo.EqualPredicate, -1)
	qb.WhereId(db_repo.EqualPredicate, id)

	info, err = qb.First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	return
}
