package examples

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/qbantek/go-httpclient/gohttp"
)

func TestPost(t *testing.T) {
	repo := Repository{
		Name: "foo",
	}

	expected_body, err := json.Marshal(&repo)
	if err != nil {
		panic(err)
	}

	gohttp.FlushMocks()
	gohttp.AddMock(&gohttp.Mock{
		Method:             http.MethodPost,
		URL:                "https://api.github.com",
		RequestBody:        string(expected_body),
		ResponseStatusCode: 201,
		ResponseBody:       `{"ok": true}`,
	})

	response, err := PostRepository(&repo)
	if err != nil {
		t.Error("error was not expected:", err)
	}
	if want := `{"ok": true}`; response.String() != want {
		t.Errorf("want %s, got %s", want, response)
	}
}
