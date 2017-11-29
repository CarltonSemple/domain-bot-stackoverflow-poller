package discovery

// DiscoveryDocAdapter ...
type DiscoveryDocAdapter struct {
	EnvironmentID     string `json:"environment_id"`
	CollectionID      string `json:"collection_id"`
	DocumentID        string `json:"document_id"`
	Content           string `json:"content"`
	URL               string `json:"url"`
	RepositoryAccount string `json:"repositoryAccount"`
	RepositoryName    string `json:"repositoryName"`
}

type updateDocContainer struct {
	EnvironmentID string       `json:"environment_id"`
	CollectionID  string       `json:"collection_id"`
	DocumentID    string       `json:"document_id"`
	File          fileData     `json:"file"`
	Metadata      fileMetadata `json:"metadata"`
}

type fileData struct {
	Value   string          `json:"value"`
	Options fileDataOptions `json:"options"`
}

type fileDataOptions struct {
	Filename    string `json:"filename"`
	SourceURL   string `json:"sourceUrl"`
	ContentType string `json:"contentType"`
}

type fileMetadata struct {
	OriginalURL       string `json:"originalUrl"`
	SrcURL            string `json:"srcUrl"`
	RepositoryAccount string `json:"repositoryAccount"`
	RepositoryName    string `json:"repositoryName"`
	UploadDate        string `json:"uploadDate"`
}
