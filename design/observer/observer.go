package observer

type RegObserver interface {
	handlerRegSuccess(userId uint64)
}

type Controller interface {
	Register(telephone, password string) uint64
	SetRegObserver(regs []RegObserver)
}

type controller struct {
	regs        []RegObserver
	userService *userService
}

func (s *controller) Register(telephone, password string) uint64 {
	// do something
	// 检验参数

	userId := s.userService.Register()
	for _, v := range s.regs {
		v.handlerRegSuccess(userId)
	}

	return userId

}

func (s *controller) SetRegObserver(regs []RegObserver) {
	s.regs = regs
}

func NewController() Controller {
	return &controller{userService: &userService{}}
}

type userService struct{}

func (s userService) Register() uint64 {
	var userId uint64 = 9527
	return userId
}

type RegNotificationObserver struct {
	RegObserver
}

func (s *RegNotificationObserver) handlerRegSuccess(userId uint64) {
	println("发送消息.......")
}
