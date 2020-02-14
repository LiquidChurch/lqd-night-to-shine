package gqlSchema

import (
  "X/goappsrv/src/helper"
)

type ImportStatus struct {
  Created   int32
  Modified  int32
  Skipped   int32
  Total     int32
}

type importStatusResolver struct {
  c helper.ContextDetail
  u *ImportStatus
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