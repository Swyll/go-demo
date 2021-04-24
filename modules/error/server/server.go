package server

import (
	"go-demo/modules/error/dao"
)

type myServer struct {
	DataEngine dao.Dao
}

type people struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type student struct {
	people
	Father *people
	Mother *people
}

func NewServer(engine dao.Dao) *myServer {
	return &myServer{
		DataEngine: engine,
	}
}

func (s *myServer) getOtherInfo(id string) {}

func (s *myServer) GetStudentInfo(id string) (*student, error) {
	stu := new(student)

	data, err := s.getBaseInfo(id)
	if err != nil {
		return nil, err
	}

	stu.people = *data
	s.getOtherInfo(id)

	return stu, nil
}

func (s *myServer) getBaseInfo(id string) (*people, error) {
	data, err := s.DataEngine.GetData(id)
	if err != nil {
		return nil, err
	}

	return &people{
		Name: data.Name,
		Age:  data.Age,
	}, nil
}
