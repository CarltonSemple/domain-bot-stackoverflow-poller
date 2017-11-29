package discovery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testEnvironmentID = "e46d8bbb-b8c2-44ac-b55d-a270d50a871b"
	testCollectionID  = "c183449f-f0d0-429f-a68c-e73b1a09c045"
)

func TestUpdateDocument(t *testing.T) {
	var (
		testDocumentContent   = "test document content"
		testDocumentURL       = "https://stackoverflow.com/questions/38486848/kubernetes-jenkins-plugin-slaves-always-offline"
		testRepositoryAccount = "stackexchange"
		testRepositoryName    = "stackoverflow"
		testDiscoveryUsername = "5627b8d2-f066-49c9-a5cf-73cae3c22527"
		testDiscoveryPassword = "sIYsVFbPcUpJ"
	)
	testDocumentID := StackoverflowURLToDiscoveryID(testDocumentURL)
	discoveryDoc := DiscoveryDocAdapter{
		EnvironmentID:     testEnvironmentID,
		CollectionID:      testCollectionID,
		DocumentID:        testDocumentID,
		Content:           testDocumentContent,
		URL:               testDocumentURL,
		RepositoryAccount: testRepositoryAccount,
		RepositoryName:    testRepositoryName,
	}

	err := UpdateDocument(discoveryDoc, testDiscoveryUsername, testDiscoveryPassword)
	assert.NoError(t, err)
}
