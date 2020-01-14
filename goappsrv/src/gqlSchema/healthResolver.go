package gqlSchema

import (
)


func (r *Resolver) Health() (*healthDetailResolver) {
  return &healthDetailResolver{"ok"}
}

type healthDetailResolver struct {
  status string
}

func (r *healthDetailResolver) Status() string {
  return r.status
}