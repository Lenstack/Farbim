package application

type MicroServer struct {
	UserApplication           UserApplication
	AuthenticationApplication AuthenticationApplication
}

func NewMicroServer(userApplication UserApplication, authenticationApplication AuthenticationApplication) *MicroServer {
	return &MicroServer{UserApplication: userApplication, AuthenticationApplication: authenticationApplication}
}
