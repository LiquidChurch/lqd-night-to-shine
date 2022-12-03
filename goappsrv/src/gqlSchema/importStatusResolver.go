package gqlSchema

import (
  "X/goappsrv/src/helper"
  "X/goappsrv/src/model"
)

type importStatusResolver struct {
  c helper.ContextDetail
  u *model.ImportStatus
}

func (r *importStatusResolver) Created() int32 {
  return r.u.Created
}

func (r *importStatusResolver) Modified() int32 {
  return r.u.Modified
}

func (r *importStatusResolver) Skipped() int32 {
  return r.u.Skipped
}

func (r *importStatusResolver) Total() int32 {
  return r.u.Total
}