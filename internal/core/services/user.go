package services

type IUserService interface {
	List()
	ShowBy()
	Create()
	Update()
	Delete()
	Verify()
}

type UserService struct {
}
