package application

type MicroserviceServer struct {
	MiddlewareApplication     MiddlewareApplication
	UserApplication           UserApplication
	AuthenticationApplication AuthenticationApplication
}

func NewMicroserviceServer(
	middlewareApplication MiddlewareApplication,
	userApplication UserApplication,
	authenticationApplication AuthenticationApplication,

) *MicroserviceServer {
	return &MicroserviceServer{
		MiddlewareApplication:     middlewareApplication,
		UserApplication:           userApplication,
		AuthenticationApplication: authenticationApplication,
	}
}
