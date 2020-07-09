package routes

import (
	"net/http"

	"../../controllers"
)

var loginRoutes = []Route{
	Route{
		URI:          "/lawyerdashboard/api/signin",
		Method:       http.MethodPost,
		Handler:      controllers.GetLawyer,
		AuthRequired: false,
	},
	Route{
		URI:          "/lawyerdashboard/api/signup",
		Method:       http.MethodPost,
		Handler:      controllers.AddLawyer,
		AuthRequired: false,
	},
	Route{
		URI:          "/clientdashboard/api/signin",
		Method:       http.MethodPost,
		Handler:      controllers.GetClient,
		AuthRequired: false,
	},
	Route{
		URI:          "/clientdashboard/api/signup",
		Method:       http.MethodPost,
		Handler:      controllers.AddClient,
		AuthRequired: false,
	},
}
