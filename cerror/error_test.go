package cerror

import (
	"errors"
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	errorCode := "ErrorCode"
	errorMsg := "MyCustomMsg"
	errorExtra := "MyCustomMsg extra information"
	orgErrorMsg1 := "Origin Error 1"

	orgError1 := errors.New(orgErrorMsg1)

	myError := New(errorCode, errorMsg)

	myError = myError.AddExtraMsg(errorExtra).AddOrigError(orgError1)
	t.Log(myError)

	if myError.Code() != errorCode {
		t.Error("Error code doesn't match")

	}

	if myError.Message() != errorMsg {
		t.Error("Error message doesn't match")
	}

	if !strings.Contains(myError.ExtraMsg(), errorExtra) {
		t.Error("Error extra message doesn't match")
	}

	if !strings.Contains(myError.Error(), orgErrorMsg1) {
		t.Error("Error origin message doesn't match")
	}

}

func TestErrorDeep(t *testing.T) {
	errorCode := "ErrorCode"
	errorMsg := "MyCustomMsg"
	errorExtra := "MyCustomMsg extra information"
	orgErrorMsg1 := "Origin Error 1"
	orgErrorMsg2 := "Origin Error 2"
	orgErrorMsg3 := "Origin Error 3"

	orgError1 := errors.New(orgErrorMsg1)
	orgError2 := errors.New(orgErrorMsg2)
	orgError3 := errors.New(orgErrorMsg3)

	myError := New(errorCode, errorMsg)

	myError = myError.AddExtraMsg(errorExtra).AddOrigError(orgError1).AddOrigError(orgError2).AddOrigError(orgError3)
	t.Log(myError)

	if myError.Code() != errorCode {
		t.Error("Error code doesn't match")

	}

	if myError.Message() != errorMsg {
		t.Error("Error message doesn't match")
	}

	if !strings.Contains(myError.ExtraMsg(), errorExtra) {
		t.Error("Error extra message doesn't match")
	}

	if !strings.Contains(myError.Error(), orgErrorMsg1) {
		t.Error("Error origin message doesn't match")
	}

	if !strings.Contains(myError.Error(), orgErrorMsg2) {
		t.Error("Error origin message doesn't match")
	}
	if !strings.Contains(myError.Error(), orgErrorMsg3) {
		t.Error("Error origin message doesn't match")
	}
}
