package testhelper

import (
	"OnionPractice/app/domain/model"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var ToDoCmpOptions = []cmp.Option{
	cmp.AllowUnexported(model.Todo{}),
	cmpopts.IgnoreFields(model.Todo{}, "id", "createdAt", "updatedAt"),
}
