package examples

import "fmt"

/*
  "current_user_url": "https://api.github.com/user",                                                                                                                                                                                            "current_user_authorizations_html_url": "https://github.com/settings/connections/applications{/client_id}",
  "authorizations_url": "https://api.github.com/authorizations",
	"repository_url": "https://api.github.com/repos/{owner}/{repo}",
*/

type EndPoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndPoints() (*EndPoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("StatusCode: %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body: %s\n", response.String()))

	var endpoints EndPoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("RepositoryUrl: %s", endpoints.RepositoryUrl))

	return &endpoints, nil
}
