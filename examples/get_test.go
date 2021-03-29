package examples

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/qbantek/go-httpclient/gohttp"
)

func TestMain(m *testing.M) {
	// Mock any further HTTP requests
	gohttp.StartMockServer()
	defer gohttp.StopMockServer()

	os.Exit(m.Run())
}

func TestGetEndPoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGitHub", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(&gohttp.Mock{
			Method: http.MethodGet,
			URL:    "https://api.github.com",
			Error:  errors.New("timeout getting GitHub endpoints"),
		})

		endpoints, err := GetEndPoints()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}
		if err == nil {
			t.Error("error was expected")
		} else {
			if want := "timeout getting GitHub endpoints"; err.Error() != want {
				t.Errorf("want %s, got %s", want, err.Error())
			}
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(&gohttp.Mock{
			Method:             http.MethodGet,
			URL:                "https://api.github.com",
			ResponseBody:       `{"current_user_url":123}`,
			ResponseStatusCode: http.StatusOK,
		})

		endpoints, err := GetEndPoints()

		if endpoints != nil {
			t.Error("no endpoints expected")
		}
		if err == nil {
			t.Error("error was expected")
		} else {
			want := "json: cannot unmarshal number into Go value of type string"
			if err.Error() != want {
				t.Errorf("want %s, got %s", want, err.Error())
			}
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(&gohttp.Mock{
			Method:             http.MethodGet,
			URL:                "https://api.github.com",
			ResponseBody:       `{"current_user_url":"https://api.github.com/users"}`,
			ResponseStatusCode: http.StatusOK,
		})

		endpoints, err := GetEndPoints()

		if err != nil {
			t.Error("error was not expected")
		}
		if endpoints == nil {
			t.Error("endpoints expected")
		}
		if endpoints.CurrentUserURL != "https://api.github.com/users" {
			t.Error("current user endpoint expected")
		}
	})
}
