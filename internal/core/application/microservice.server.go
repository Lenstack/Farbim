package application

type MicroserviceServer struct {
	UserApplication           UserApplication
	AuthenticationApplication AuthenticationApplication
}

func NewMicroserviceServer(
	userApplication UserApplication,
	authenticationApplication AuthenticationApplication,
) *MicroserviceServer {
	return &MicroserviceServer{
		UserApplication:           userApplication,
		AuthenticationApplication: authenticationApplication,
	}
}
