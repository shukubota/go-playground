package test

type OtherModel struct {
	id   int
	name string
}

func NewOtherModel() (OtherModel, error) {
	newModel := OtherModel{
		id:   1,
		name: "hoge",
	}
	return newModel, nil
}

func (model OtherModel) GetName() string {
	return model.name
}
