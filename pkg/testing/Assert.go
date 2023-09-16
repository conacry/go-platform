package commonTesting

import (
	goErrors "errors"
	"fmt"
	"github.com/conacry/go-platform/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func AssertErrors(t *testing.T, err error, expectedErrors []error) {
	require.Error(t, err)

	var errs *errors.Errors
	ok := goErrors.As(err, &errs)
	require.True(t, ok)

	require.Equal(t, len(expectedErrors), errs.Size())
	for _, expectedError := range expectedErrors {
		assert.True(t, errs.Contains(expectedError), fmt.Sprintf("fail assert - missing expected error = %q", expectedError))
	}
}

func AssertErrorCodes(t *testing.T, err error, expectedErrorCodes []errors.ErrorCode) {
	require.Error(t, err)

	var errs *errors.Errors
	ok := goErrors.As(err, &errs)
	require.True(t, ok)

	actualCodes := map[string]interface{}{}
	for _, actualErr := range errs.ToArray() {
		actualCodes[actualErr.Code().String()] = &struct{}{}
	}

	require.EqualValues(t, len(expectedErrorCodes), errs.Size())
	for _, expectedCode := range expectedErrorCodes {
		_, ok := actualCodes[expectedCode.String()]
		assert.True(t, ok, fmt.Sprintf("fail assert - missing expected error code = %q", expectedCode))
	}
}

func AssertFieldsNotNil(t *testing.T, object interface{}) {
	reflectObjectValue := reflect.ValueOf(object)
	reflectObjectElem := getReflectElem(reflectObjectValue)

	if reflectObjectElem.Kind() != reflect.Struct {
		assert.Fail(t, "object is not struct")
		return
	}

	reflectStructType := reflectObjectElem.Type()
	for i := 0; i < reflectStructType.NumField(); i++ {
		fieldType := reflectStructType.Field(i)
		AssertFieldNotNil(t, object, fieldType.Name)
	}
}

func AssertFieldNotNil(t *testing.T, object interface{}, fieldName string) {
	isNil, err := isValueNil(object, fieldName)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	if isNil {
		errMsg := fmt.Sprintf("field value with name = %q is nil", fieldName)
		assert.Fail(t, errMsg)
	}
}

func AssertFieldIsNil(t *testing.T, object interface{}, fieldName string) {
	isNil, err := isValueNil(object, fieldName)
	if err != nil {
		assert.Fail(t, err.Error())
	}

	if isNil {
		return
	}

	errMsg := fmt.Sprintf("field value with name = %q isn't nil", fieldName)
	assert.Fail(t, errMsg)
}

func isValueNil(object interface{}, fieldName string) (bool, error) {
	reflectFieldValue := getFieldReflectValue(object, fieldName)
	switch reflectFieldValue.Kind() {
	case reflect.Map, reflect.Chan, reflect.Pointer, reflect.Slice:
		return reflectFieldValue.IsNil(), nil
	case reflect.Interface:
		result := reflectFieldValue.IsNil() || reflectFieldValue.Elem().IsNil()
		return result, nil
	default:
		errMsg := fmt.Sprintf("unsupported type %q of field %q to check for nil", reflectFieldValue.Kind(), fieldName)
		return false, goErrors.New(errMsg)
	}
}

func AssertFieldValue(
	t *testing.T,
	object interface{},
	fieldName string,
	expectedFieldValue interface{},
) {
	reflectFieldValue := getFieldReflectValue(object, fieldName)
	reflectExpectedFieldValue := reflect.ValueOf(expectedFieldValue)
	if !compareReflectValues(reflectFieldValue, reflectExpectedFieldValue) {
		errMsg := fmt.Sprintf("field value with name = '%s' not equal with expected value. "+
			"Expected value = '%#v'", fieldName, expectedFieldValue)
		assert.Fail(t, errMsg)
		return
	}
}

func compareReflectValues(actualReflectValue, expectedReflectValue reflect.Value) bool {
	reflectFieldKind := actualReflectValue.Kind()
	reflectExpectedFieldKind := expectedReflectValue.Kind()

	if reflectFieldKind != reflectExpectedFieldKind {
		if reflectFieldKind == reflect.Interface {
			isImplInterface := expectedReflectValue.Type().Implements(actualReflectValue.Type())
			if !isImplInterface {
				return false
			}
		} else {
			return false
		}
	}

	kind := actualReflectValue.Kind()
	switch kind {
	case reflect.String:
		return actualReflectValue.String() == expectedReflectValue.String()
	case reflect.Bool:
		return actualReflectValue.Bool() == expectedReflectValue.Bool()
	case reflect.Int:
		return actualReflectValue.Int() == expectedReflectValue.Int()
	case reflect.Uint:
		return actualReflectValue.Uint() == expectedReflectValue.Uint()
	case reflect.Ptr, reflect.Chan, reflect.Slice:
		return actualReflectValue.Pointer() == expectedReflectValue.Pointer()
	case reflect.Map:
		return isMapEqual(actualReflectValue, expectedReflectValue)
	case reflect.Interface:
		if expectedReflectValue.Kind() == reflect.Interface {
			return actualReflectValue.Elem().Pointer() == expectedReflectValue.Elem().Pointer()
		} else {
			return actualReflectValue.Elem().Pointer() == expectedReflectValue.Pointer()
		}
	case reflect.Struct:
		return isStructEqual(actualReflectValue, expectedReflectValue)
	}

	return false
}

func isMapEqual(actualReflectValue, expectedReflectValue reflect.Value) bool {
	expectedMapIter := expectedReflectValue.MapRange()
	actualMapIter := actualReflectValue.MapRange()

	expectedKeys := make([]reflect.Value, 0)
	expectedValues := make([]reflect.Value, 0)
	actualKeys := make([]reflect.Value, 0)
	actualValues := make([]reflect.Value, 0)
	for expectedMapIter.Next() {
		expectedKeys = append(expectedKeys, expectedMapIter.Key())
		expectedValues = append(expectedValues, expectedMapIter.Value())
	}
	for actualMapIter.Next() {
		actualKeys = append(actualKeys, actualMapIter.Key())
		actualValues = append(actualValues, actualMapIter.Value())
	}

	if len(expectedKeys) != len(actualKeys) && len(expectedValues) != len(actualValues) {
		return false
	}

	for i := range expectedKeys {
		expectedKey := expectedKeys[i]
		expectedValue := expectedValues[i]

		var validActualKey bool
		for _, actualKey := range actualKeys {
			validActualKey = compareReflectValues(actualKey, expectedKey)
			if validActualKey {
				break
			}
		}
		if !validActualKey {
			return false
		}

		var validActualValue bool
		for _, actualValue := range actualValues {
			validActualValue = compareReflectValues(actualValue, expectedValue)
			if validActualValue {
				break
			}
		}
		if !validActualValue {
			return false
		}
	}

	return true
}

func isStructEqual(actualReflectValue, expectedReflectValue reflect.Value) bool {
	actualReflectStructType := actualReflectValue.Type()

	for i := 0; i < actualReflectStructType.NumField(); i++ {
		actualFieldType := actualReflectStructType.Field(i)

		actualReflectFieldValue := actualReflectValue.FieldByName(actualFieldType.Name)
		expectedReflectFieldValue := expectedReflectValue.FieldByName(actualFieldType.Name)
		if !compareReflectValues(actualReflectFieldValue, expectedReflectFieldValue) {
			return false
		}
	}
	return true
}

func getFieldReflectValue(object interface{}, fieldName string) reflect.Value {
	reflectObjectValue := reflect.ValueOf(object)
	reflectObjectElem := getReflectElem(reflectObjectValue)

	if reflectObjectElem.Kind() != reflect.Struct {
		panic("object is not struct")
	}

	reflectStructType := reflectObjectElem.Type()
	for i := 0; i < reflectStructType.NumField(); i++ {
		fieldType := reflectStructType.Field(i)
		if fieldType.Name != fieldName {
			continue
		}

		return reflectObjectElem.FieldByName(fieldType.Name)
	}

	panic(fmt.Sprintf("Field by name = %q not found in struct", fieldName))
}

func getReflectElem(reflectValue reflect.Value) reflect.Value {
	if reflectValue.Kind() == reflect.Ptr {
		return reflectValue.Elem()
	} else {
		return reflectValue
	}
}
