package handler

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_SayHandler_Should_Be_Hello_World_In_Thai(t *testing.T) {
	expected := `{"message":"สวัสดีชาวโลก"}`

	request := httptest.NewRequest("GET", "/hello", nil)
	writer := httptest.NewRecorder()

	repository := MockRepository{}

	resource := Resource{Repository: &repository}

	testRoute := gin.Default()
	testRoute.GET("/hello", resource.SayHandler)
	testRoute.ServeHTTP(writer, request)
	response := writer.Result()
	actual, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Errorf("ReadAll error with %s", err.Error())
	}

	if string(actual) != expected {
		t.Errorf("actual %s not equals to expected %s", string(actual), expected)
	}
}

func Test_translate_Hello_World_Should_Be_In_Thai(t *testing.T) {
	expected := Message{Message: "สวัสดีชาวโลก"}

	message := Message{}
	message.translate("Hello World")

	if message != expected {
		t.Errorf("actual %+v is not equals expected %+v", message, expected)
	}
}
