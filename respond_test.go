package framework

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/xubiosueldos/conexionBD/structGormModel"
)

func Test_RespondError(t *testing.T) {
	t.Run("Check RespondError", func(context *testing.T) {
		request, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			context.Fatal(err)
		}
		expectedBody := `{"codigo":"400","error":"mensaje de error"}`
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			RespondError(writer, 400, "mensaje de error")
		})

		handler.ServeHTTP(recorder, request)
		if recorder.Body.String() != expectedBody {
			context.Errorf("bad response body, wanted %v got %v",
				recorder.Body, expectedBody)
		}
	})
}

func Test_RespondJSON(t *testing.T) {
	t.Run("Check RespondJSON", func(context *testing.T) {
		var body structGormModel.GormModel
		body.ID = 1
		body.CreatedAt = time.Date(2000, 01, 01, 01, 01, 01, 01, time.Local)
		body.UpdatedAt = time.Date(2000, 01, 01, 01, 01, 01, 01, time.Local)

		expectedBody := `{"ID":1,"CreatedAt":"2000-01-01T01:01:01.000000001-03:00","UpdatedAt":"2000-01-01T01:01:01.000000001-03:00","DeletedAt":null}`

		handler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			RespondJSON(writer, 200, body)
		})
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			context.Fatal(err)
		}

		handler.ServeHTTP(recorder, request)
		if recorder.Body.String() != expectedBody {
			context.Errorf("bad response body, wanted %v got %v",
				expectedBody, recorder.Body)
		}
	})
}

/*func FakeHandler(writer http.ResponseWriter, request *http.Request) {
	RespondError(writer, 400, "mensaje de error")
	//writer.WriteHeader(http.StatusOK)
}*/
