package examples

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

	var endpoints EndPoints
	if err := response.UnmarshalJSONBody(&endpoints); err != nil {
		return nil, err
	}

	return &endpoints, nil
}
