package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-plugin-template/server/constants"
	"github.com/mattermost/mattermost-plugin-template/server/serializers"
)

type Client interface {
	// TODO: Below is for demo purposes only remove it later
	GetMockPosts() ([]*serializers.DummyPost, int, error)
}

type client struct {
	plugin     *Plugin
	httpClient *http.Client
}

// TODO: Below is for demo purposes only remove it later
// This is a sample API call to a mock server available at "https://jsonplaceholder.typicode.com/posts" which returns some dummy posts
func (c *client) GetMockPosts() ([]*serializers.DummyPost, int, error) {
	var posts []*serializers.DummyPost
	_, statusCode, err := c.CallJSON(c.plugin.getConfiguration().MockClientBaseURL, constants.PathGetMockPosts, http.MethodGet, nil, &posts, nil)
	if err != nil {
		return nil, statusCode, err
	}

	return posts, statusCode, nil
}

// Wrapper to make REST API requests with "application/json" type content
func (c *client) CallJSON(url, path, method string, in, out interface{}, formValues url.Values) (responseData []byte, statusCode int, err error) {
	contentType := "application/json"
	buf := &bytes.Buffer{}
	if err = json.NewEncoder(buf).Encode(in); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return c.Call(url, method, path, contentType, buf, out, formValues)
}

// Makes HTTP request to REST APIs
func (c *client) Call(basePath, method, path, contentType string, inBody io.Reader, out interface{}, formValues url.Values) (responseData []byte, statusCode int, err error) {
	errContext := fmt.Sprintf("Template client: Call failed: method:%s, path:%s", method, path)
	pathURL, err := url.Parse(path)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.WithMessage(err, errContext)
	}

	if pathURL.Scheme == "" || pathURL.Host == "" {
		var baseURL *url.URL
		baseURL, err = url.Parse(basePath)
		if err != nil {
			return nil, http.StatusInternalServerError, errors.WithMessage(err, errContext)
		}
		if path[0] != '/' {
			path = "/" + path
		}
		path = baseURL.String() + path
	}

	var req *http.Request
	if formValues != nil {
		req, err = http.NewRequest(method, path, strings.NewReader(formValues.Encode()))
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	} else {
		req, err = http.NewRequest(method, path, inBody)
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}
	}

	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if resp.Body == nil {
		return nil, resp.StatusCode, nil
	}
	defer resp.Body.Close()

	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if out != nil {
			if err = json.Unmarshal(responseData, out); err != nil {
				return responseData, http.StatusInternalServerError, err
			}
		}
		return responseData, resp.StatusCode, nil

	case http.StatusNoContent:
		return nil, resp.StatusCode, nil

	case http.StatusNotFound:
		return nil, resp.StatusCode, ErrNotFound
	}

	errResp := serializers.ErrorResponse{}
	if err = json.Unmarshal(responseData, &errResp); err != nil {
		return responseData, http.StatusInternalServerError, errors.WithMessagef(err, "status: %s", resp.Status)
	}
	return responseData, resp.StatusCode, fmt.Errorf("errorMessage %s", errResp.Message)
}

func InitClient(p *Plugin) Client {
	return &client{
		plugin:     p,
		httpClient: &http.Client{},
	}
}
