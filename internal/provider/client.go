// client.go
package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ListParams struct {
	Page     int
	PageSize int
	Filter   map[string]string
}

type Client struct {
	ServerURL  string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(serverURL, apiKey string) *Client {
	return &Client{
		ServerURL:  serverURL,
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) doRequest(method, path string, body interface{}) (*http.Response, error) {
	var bodyReader *bytes.Buffer

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %v", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.ServerURL, path), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		defer resp.Body.Close()
		var apiError APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
			// If we can't decode the error response, return a generic error
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
		}
		return nil, &apiError
	}

	return resp, nil
}

// Stack operations
func (c *Client) CreateStack(stack StackUpdate) (*StackResponse, error) {
	resp, err := c.doRequest("POST", "/api/v1/stacks", stack)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result StackResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) GetStack(id string) (*StackResponse, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("/api/v1/stacks/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result StackResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) UpdateStack(id string, stack StackUpdate) (*StackResponse, error) {
	resp, err := c.doRequest("PUT", fmt.Sprintf("/api/v1/stacks/%s", id), stack)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result StackResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) DeleteStack(id string) error {
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/api/v1/stacks/%s", id), nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// Component operations
func (c *Client) CreateComponent(component ComponentBody) (*ComponentResponse, error) {
	resp, err := c.doRequest("POST", "/api/v1/components", component)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ComponentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) GetComponent(id string) (*ComponentResponse, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("/api/v1/components/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ComponentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) UpdateComponent(id string, component ComponentUpdate) (*ComponentResponse, error) {
	resp, err := c.doRequest("PUT", fmt.Sprintf("/api/v1/components/%s", id), component)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ComponentResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) DeleteComponent(id string) error {
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/api/v1/components/%s", id), nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

// client.go (add these methods)

func (c *Client) CreateServiceConnector(connector ServiceConnectorBody) (*ServiceConnectorResponse, error) {
	resp, err := c.doRequest("POST", "/api/v1/service_connectors", connector)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ServiceConnectorResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) GetServiceConnector(id string) (*ServiceConnectorResponse, error) {
	resp, err := c.doRequest("GET", fmt.Sprintf("/api/v1/service_connectors/%s", id), nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ServiceConnectorResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) UpdateServiceConnector(id string, connector ServiceConnectorUpdate) (*ServiceConnectorResponse, error) {
	resp, err := c.doRequest("PUT", fmt.Sprintf("/api/v1/service_connectors/%s", id), connector)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ServiceConnectorResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &result, nil
}

func (c *Client) DeleteServiceConnector(id string) error {
	resp, err := c.doRequest("DELETE", fmt.Sprintf("/api/v1/service_connectors/%s", id), nil)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func (c *Client) ListStacks(params *ListParams) (*Page[StackResponse], error) {
	if params == nil {
		params = &ListParams{
			Page:     1,
			PageSize: 100,
		}
	}

	url := fmt.Sprintf("%s/api/v1/stacks?page=%d&size=%d", c.ServerURL, params.Page, params.PageSize)

	// Add filters if any
	for k, v := range params.Filter {
		url = fmt.Sprintf("%s&%s=%s", url, k, v)
	}

	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var result Page[StackResponse]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}

// Add pagination support to all list methods
func (c *Client) ListStackComponents(params *ListParams) (*Page[ComponentResponse], error) {
	url := "/api/v1/components"
	if params != nil {
		url = fmt.Sprintf("%s?page=%d&size=%d", url, params.Page, params.PageSize)
		for k, v := range params.Filter {
			url = fmt.Sprintf("%s&%s=%s", url, k, v)
		}
	}

	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Page[ComponentResponse]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}

func (c *Client) ListServiceConnectors(params *ListParams) (*Page[ServiceConnectorResponse], error) {
	url := "/api/v1/service_connectors"
	if params != nil {
		url = fmt.Sprintf("%s?page=%d&size=%d", url, params.Page, params.PageSize)
		for k, v := range params.Filter {
			url = fmt.Sprintf("%s&%s=%s", url, k, v)
		}
	}

	resp, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Page[ServiceConnectorResponse]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &result, nil
}
