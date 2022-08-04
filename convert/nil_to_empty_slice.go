package convert

import "reflect"

// ConvertNilToEmptySlice 将对象中的 nil 转换成 空slice
func ConvertNilToEmptySlice(object interface{}) interface{} {
	objectType := reflect.TypeOf(object)
	objectValue := reflect.ValueOf(object)
	if objectType == nil {
		return nil
	}
	switch objectType.Kind() {
	case reflect.Ptr:
		if !objectValue.Elem().CanInterface() {
			break
		}
		newObjectValue := reflect.New(objectType.Elem())
		newObjectValue.Elem().Set(reflect.ValueOf(ConvertNilToEmptySlice(objectValue.Elem().Interface())))
		return newObjectValue.Interface()
	case reflect.Struct:
		newObjectValue := reflect.New(objectType)
		for i := 0; i < objectType.NumField(); i++ {
			if !objectValue.Field(i).CanInterface() {
				continue
			}
			newObjectValue.Elem().Field(i).Set(reflect.ValueOf(ConvertNilToEmptySlice(objectValue.Field(i).Interface())))
		}
		return newObjectValue.Elem().Interface()
	case reflect.Map:
		newObjectValue := reflect.MakeMapWithSize(objectType, objectValue.Len())
		for _, key := range objectValue.MapKeys() {
			if !objectValue.MapIndex(key).CanInterface() {
				continue
			}
			newObjectValue.SetMapIndex(key, reflect.ValueOf(ConvertNilToEmptySlice(objectValue.MapIndex(key).Interface())))
		}
		return newObjectValue.Interface()
	case reflect.Array, reflect.Slice:
		if objectType.Elem().Kind() == reflect.Uint8 {
			break
		}
		newObjectValue := reflect.MakeSlice(objectType, objectValue.Len(), objectValue.Len())
		for i := 0; i < objectValue.Len(); i++ {
			if !objectValue.Index(i).CanInterface() {
				continue
			}
			newObjectValue.Index(i).Set(reflect.ValueOf(ConvertNilToEmptySlice(objectValue.Index(i).Interface())))
		}
		return newObjectValue.Interface()
	}
	return object
}
