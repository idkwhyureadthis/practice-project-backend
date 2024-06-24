package service

type Response struct {
	Code int
	Data string
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) GetStatus() Response {
	return Response{
		Code: 200,
		Data: "service currently active",
	}
}
