type Authenticator interface {
	GetAuthHeaders() map[string]string
}

type Client struct {
	auth       Authenticator
	BaseURL    string
	HTTPClient *http.Client
}