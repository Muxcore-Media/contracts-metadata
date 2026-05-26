package metadata

import "context"

type MetadataResult struct {
	Title       string
	Year        int
	Synopsis    string
	Genres      []string
	Cast        []string
	Directors   []string
	Rating      float64
	RatingCount int
	Runtime     int
	MPAARating  string
	Studio      string
	Country     []string
	ExternalIDs map[string]string
	Source      string
	Language    string
}

type MetadataProvider interface {
	Search(ctx context.Context, title string, year int, mediaType string) ([]MetadataResult, error)
	GetByID(ctx context.Context, externalID, source string) (MetadataResult, error)
}

const EventMetadataFetched = "metadata.fetched"

type MetadataFetchedPayload struct {
	MediaID    string `json:"media_id"`
	Title      string `json:"title"`
	Source     string `json:"source"`
	ExternalID string `json:"external_id"`
}
