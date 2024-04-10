package factory

import (
	"reflect"
	"strings"
	"time"

	"github.com/sCuz12/celeritas/factory/faker"
)



type Factory struct {

}


type Model interface {
	Table() string
}

func (f *Factory) PopulateModel(m Model) Model {
	return f.fillModelData(m)
}

// creates multiple instances of the provided model type
func (f *Factory) PopulateModels(m Model, count int) []Model {
	models := make([]Model, count)
    modelType := reflect.TypeOf(m)

	for i := range models {
		// Create a new instance of the model for each iteration
        newModelPtr := reflect.New(modelType.Elem())
        newModel := newModelPtr.Interface().(Model)

        models[i] = f.fillModelData(newModel)
	}
    
	return models
}


func (f *Factory) fillModelData(m Model) Model {
	modelVal := reflect.ValueOf(m).Elem()
	modelType := modelVal.Type() // Get the type of the model

	for i := 0; i < modelVal.NumField(); i++ {
		field := modelVal.Field(i)
		fieldName := modelType.Field(i).Name // get field name

		f.fillerByName(fieldName, &field)
	}

	return m
}

// Fills the data by name (Ex : email )
func (f *Factory) fillerByName(fieldName string, field *reflect.Value) string {
	fieldName = strings.ToLower(fieldName)

	switch fieldName {
	case "email":
		{
            field.SetString(faker.Email())
		}
    case "name","firstname","first_name" :
        {
            field.SetString(faker.Name())
        }
    case "lastname", "last_name" : {
        field.SetString(faker.LastName())
    }

	default:
		{
            //generate data by type instead of field name
			f.fillerByType(field)
		}
	}
	return ""
}

func (f *Factory) fillerByType(field *reflect.Value) string {

	switch field.Kind() {

	case reflect.String:
		{
             field.SetString(faker.Random())
		}
    case reflect.Bool:
        field.SetBool(faker.RandomBool())

    case reflect.Struct:
        if field.Type() == reflect.TypeOf(time.Time{}) {
            field.Set(reflect.ValueOf(faker.RandomDate()))
        }

    case reflect.Int, reflect.Int64:
        field.SetInt(int64(faker.RandomInt(0, 100))) // Example range: 0 to 100

    case reflect.Float32, reflect.Float64:
        field.SetFloat(faker.RandomFloat(0.0, 100.0)) // Example range: 0.0 to 100.0
	}
    return ""
}