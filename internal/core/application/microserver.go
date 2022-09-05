package application

type MicroServer struct {
	UserApplication
}

func NewMicroServer(userApplication UserApplication) *MicroServer {
	return &MicroServer{UserApplication: userApplication}
}
