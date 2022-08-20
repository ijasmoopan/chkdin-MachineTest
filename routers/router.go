package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ijasmoopan/chkdin-MachineTest/user"
	"github.com/ijasmoopan/chkdin-MachineTest/repository"
) 

func Route() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	db := respository.ConnectDB()
	userdb := user.InterfaceHandler(db)

	router.Post("/signup", userdb.PostUser)
	router.Post("/login", userdb.Login)

	userRouter := router.Group(nil)

	userRouter.Use(userdb.IsUserAuthorized)
	userRouter.Get("/{email}", userdb.GetUser)
	userRouter.Patch("/{email}/patch", userdb.PatchUser)
	userRouter.Delete("/{email}/delete", userdb.DeleteUser)

	return router
}
