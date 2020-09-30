package shortener

// Request is an object representing a create or read request
type Request struct {
	Slug string `json:"slug"`
	URL  string `json:"url"`
}

// Response is an object representing a create or read response
type Response struct {
	Slug string `json:"slug"`
	URL  string `json:"url"`
}
