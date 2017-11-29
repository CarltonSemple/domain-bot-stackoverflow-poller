package discovery

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

var (
	apiVersionDate = "2017-09-01"
)

// UpdateDocument ...
func UpdateDocument(document DiscoveryDocAdapter, discoveryUsername, discoveryPassword string) error {
	t := time.Now()
	jsonPayload := updateDocContainer{
		EnvironmentID: document.EnvironmentID,
		CollectionID:  document.CollectionID,
		DocumentID:    document.DocumentID,
		File: fileData{
			Value: document.Content,
			Options: fileDataOptions{
				Filename:    document.URL,
				SourceURL:   document.URL,
				ContentType: "text/html",
			},
		},
		Metadata: fileMetadata{
			OriginalURL:       document.URL,
			SrcURL:            document.URL,
			RepositoryAccount: document.RepositoryAccount,
			RepositoryName:    document.RepositoryName,
			UploadDate:        fmt.Sprintf(`%d-%02d-%02d`, t.Year(), t.Month(), t.Day()),
		},
	}

	//fileBytes, _ := json.Marshal(&)
	sb := "<html><head><title> </title></head>" + document.Content + "</html>"
	fileBytes := []byte(sb)
	metadataBytes, _ := json.Marshal(&jsonPayload.Metadata)

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", jsonPayload.File.Options.Filename)
	if err != nil {
		return errors.Wrapf(err, "createFormField: ")
	}
	if _, err := fw.Write(fileBytes); err != nil {
		return errors.Wrapf(err, "form.write: ")
	}
	if _, err = io.Copy(fw, bytes.NewReader(fileBytes)); err != nil {
		return errors.Wrapf(err, "io.copy: ")
	}

	fw, err = w.CreateFormField("metadata")
	if err != nil {
		fmt.Println("err: ", err)
		return errors.Wrapf(err, "createFormField: ")
	}
	if _, err := fw.Write(metadataBytes); err != nil {
		fmt.Println("errrrr: ", err)
		return errors.Wrapf(err, "form.write: ")
	}

	w.Close()

	url := `https://gateway.watsonplatform.net/discovery/api/v1/environments/` +
		document.EnvironmentID + `/collections/` +
		document.CollectionID + `/documents/` + document.DocumentID + `?version=` + apiVersionDate
	fmt.Println("URL: ", url)

	bytesBuffer := new(bytes.Buffer)
	json.NewEncoder(bytesBuffer).Encode(jsonPayload)

	req, err := http.NewRequest("POST", url, &b)

	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(discoveryUsername+`:`+discoveryPassword)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrapf(err, "http.Client.Do: ")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("response: %v\n%v\n%v", resp.Status, resp.Header, string(body))
	}
	return nil
}
