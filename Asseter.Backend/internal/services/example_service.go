package services

import "log"

type ExampleService struct {
} 

func NewExampleService() *ExampleService{ 
	return  &ExampleService{}
}

func (e *ExampleService) Log(message string) {
	log.Println(message)
}
