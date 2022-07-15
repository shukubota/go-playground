package test

import "fmt"

type OtherModel struct {
	id       int
	name     string
	OpenName string
}

func NewOtherModel() (OtherModel, error) {
	newModel := OtherModel{
		id:       1,
		name:     "hoge",
		OpenName: "openname",
	}
	return newModel, nil
}

func (model OtherModel) GetName() string {
	return model.name
}

func (model OtherModel) SetName(params string) {
	fmt.Println("setname")
	fmt.Println(&params)
	fmt.Println(params)
	model.name = params
	fmt.Println(model)
	fmt.Println("in setName ")
	// return model
}
