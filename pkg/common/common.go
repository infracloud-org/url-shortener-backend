package common

type ShortenRequest struct {
	OriginalURL string `json:"originalURL"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortURL,omitempty"`
	Error    string `json:"error,omitempty"`
}
