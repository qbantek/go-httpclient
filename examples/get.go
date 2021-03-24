package examples

import "fmt"

// EndPoints holds the values of the urls for some endpoints of the GitHub API.
type EndPoints struct {
	CurrentUserURL   string `json:"current_user_url"`
	AuthorizationURL string `json:"authorizations_url"`
	RepositoryURL    string `json:"repository_url"`
}

// GetEndPoints gets all endpoint urls from the GitHub API and returns the
// results as an EndPoint value.
func GetEndPoints() (*EndPoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("StatusCode: %d", response.StatusCode()))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status()))
	fmt.Println(fmt.Sprintf("Body: %s\n", response.String()))

	var endpoints EndPoints
	if err := response.UnmarshalJSONBody(&endpoints); err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("RepositoryUrl: %s", endpoints.RepositoryURL))

	return &endpoints, nil
}
