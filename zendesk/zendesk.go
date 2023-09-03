package zendesk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/connorb645/appeak-go/domain"
	"github.com/connorb645/appeak-go/store"
)

type zendeskDocumentsResponse struct {
	Page	  int               `json:"page"`
	PageCount int               `json:"page_count"`
	PerPage   int               `json:"per_page"`
	PextPage  string            `json:"next_page"`
	Count     int               `json:"count"`
	Articles  []domain.Document `json:"articles"`
}

type zendesk struct {
	baseURL  string
	apiToken string
	username string
}

func NewZendesk(baseURL string, username string, apiToken string) store.HelpCenterProvider {
	return &zendesk{
		baseURL:  baseURL,
		apiToken: apiToken,
		username: username,
	}
}

func (z *zendesk) url(suffixPath string) string {
	return fmt.Sprintf("%s%s/%s", "https://", z.baseURL, suffixPath)
}

func (z *zendesk) authUsername() string {
	return fmt.Sprintf("%s/token", z.username)
}

func (z *zendesk) authPassword() string {
	return z.apiToken
}

func (z *zendesk) makeRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", z.url(path), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(z.authUsername(), z.authPassword())
	req.Header.Add("Accept", "application/json")
	return req, nil
}

func (z *zendesk) request(path string) ([]byte, error) {
	client := &http.Client{}
	req, err := z.makeRequest(path)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (z *zendesk) GetDocuments() ([]domain.Document, error) {
	body, err := z.request("api/v2/help_center/articles")
	if err != nil {
		return nil, err
	}

	var response zendeskDocumentsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Response: %v", response)

	return response.Articles, nil
}

func (z *zendesk) GetDocument(id string) (*domain.Document, error) {
	body, err := z.request(fmt.Sprintf("api/v2/help_center/articles/%s", id))
	if err != nil {
		return nil, err
	}

	var doc domain.Document
	err = json.Unmarshal(body, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}
