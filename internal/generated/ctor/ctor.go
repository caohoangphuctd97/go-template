package ctor

/* DO NOT EDIT. This file generated due to '@ctor' annotation*/

import (
	"github.com/caohoangphuctd97/go-test/internal/app/controllers"
	databases "github.com/caohoangphuctd97/go-test/internal/app/database"
	"github.com/caohoangphuctd97/go-test/internal/app/repo"
	routes "github.com/caohoangphuctd97/go-test/internal/app/routers"
	"github.com/caohoangphuctd97/go-test/pkg/typapp"
)

func init() {
	typapp.Provide("", databases.NewDatabases)
	typapp.Provide("", repo.NewBookRepo)
	typapp.Provide("", controllers.NewBookSvc)
	typapp.Provide("", routes.NewBookCntrl)
}
