package gohttp

type httpClient struct{}

// HTTPClient ...
type HTTPClient interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
}

// Get ...
func (c *httpClient) Get() {
	panic("implement me!")
}

// Post ...
func (c *httpClient) Post() {
	panic("implement me!")
}

// Put ...
func (c *httpClient) Put() {
	panic("implement me!")
}

// Patch ...
func (c *httpClient) Patch() {
	panic("implement me!")
}

// Delete ...
func (c *httpClient) Delete() {
	panic("implement me!")
}

// New ...
func New() HTTPClient {
	client := &httpClient{}
	return client
}
