package etc

import (
	"errors"
	"reflect"
)

const (
	TagStructTransfer string = "tst"
)

// TransferStruct makes transfer between DTO, DAO or Logic structures by tag 'tst'. Accepts pointers
func TransferStruct(src interface{}, dest interface{}) (modifiedDest interface{}, err error) {
	srcVal, destVal := reflect.ValueOf(src), reflect.ValueOf(dest)

	if srcVal.Kind() != reflect.Ptr || destVal.Kind() != reflect.Ptr {
		return nil, errors.New("can't transfer copies")
	}
	srcVal, destVal = srcVal.Elem(), destVal.Elem()
	srcType, destType := srcVal.Type(), destVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		fieldTag := srcType.Field(i).Tag.Get(TagStructTransfer)
		if fieldTag == "" {
			continue
		}
		// TODO: find out how to find struct by tag in another tag
		//destTag := destType.
	}
	print(destType)
	return nil, nil
}
