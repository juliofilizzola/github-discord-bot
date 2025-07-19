package model

var registeredModels []interface{}

func RegisterModel(model interface{}) {
	registeredModels = append(registeredModels, model)
}

func GetRegisteredModels() []interface{} {
	return registeredModels
}
