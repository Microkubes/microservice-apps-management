// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "apps-management": apps Resource Client
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-apps-management/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-apps-management
// --version=v1.2.0-dirty

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// DeleteAppAppsPath computes a request path to the deleteApp action of apps.
func DeleteAppAppsPath(appID string) string {
	param0 := appID

	return fmt.Sprintf("/apps/%s", param0)
}

// Delete an app
func (c *Client) DeleteAppApps(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteAppAppsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteAppAppsRequest create the request corresponding to the deleteApp action endpoint of the apps resource.
func (c *Client) NewDeleteAppAppsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetAppsPath computes a request path to the get action of apps.
func GetAppsPath(appID string) string {
	param0 := appID

	return fmt.Sprintf("/apps/%s", param0)
}

// Get app by id
func (c *Client) GetApps(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetAppsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetAppsRequest create the request corresponding to the get action endpoint of the apps resource.
func (c *Client) NewGetAppsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetMyAppsAppsPath computes a request path to the getMyApps action of apps.
func GetMyAppsAppsPath() string {

	return fmt.Sprintf("/apps/my")
}

// Get all user's apps
func (c *Client) GetMyAppsApps(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetMyAppsAppsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetMyAppsAppsRequest create the request corresponding to the getMyApps action endpoint of the apps resource.
func (c *Client) NewGetMyAppsAppsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// GetUserAppsAppsPath computes a request path to the getUserApps action of apps.
func GetUserAppsAppsPath(userID string) string {
	param0 := userID

	return fmt.Sprintf("/apps/users/%s/all", param0)
}

// Get app by id
func (c *Client) GetUserAppsApps(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserAppsAppsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserAppsAppsRequest create the request corresponding to the getUserApps action endpoint of the apps resource.
func (c *Client) NewGetUserAppsAppsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RegenerateClientSecretAppsPath computes a request path to the regenerateClientSecret action of apps.
func RegenerateClientSecretAppsPath(appID string) string {
	param0 := appID

	return fmt.Sprintf("/apps/%s/regenerate-secret", param0)
}

// Regenerate client secret
func (c *Client) RegenerateClientSecretApps(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewRegenerateClientSecretAppsRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegenerateClientSecretAppsRequest create the request corresponding to the regenerateClientSecret action endpoint of the apps resource.
func (c *Client) NewRegenerateClientSecretAppsRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RegisterAppAppsPath computes a request path to the registerApp action of apps.
func RegisterAppAppsPath() string {

	return fmt.Sprintf("/apps")
}

// Register new app
func (c *Client) RegisterAppApps(ctx context.Context, path string, payload *AppPayload, contentType string) (*http.Response, error) {
	req, err := c.NewRegisterAppAppsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegisterAppAppsRequest create the request corresponding to the registerApp action endpoint of the apps resource.
func (c *Client) NewRegisterAppAppsRequest(ctx context.Context, path string, payload *AppPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// UpdateAppAppsPath computes a request path to the updateApp action of apps.
func UpdateAppAppsPath(appID string) string {
	param0 := appID

	return fmt.Sprintf("/apps/%s", param0)
}

// Register new app
func (c *Client) UpdateAppApps(ctx context.Context, path string, payload *AppPayload, contentType string) (*http.Response, error) {
	req, err := c.NewUpdateAppAppsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateAppAppsRequest create the request corresponding to the updateApp action endpoint of the apps resource.
func (c *Client) NewUpdateAppAppsRequest(ctx context.Context, path string, payload *AppPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// VerifyAppAppsPath computes a request path to the verifyApp action of apps.
func VerifyAppAppsPath() string {

	return fmt.Sprintf("/apps/verify")
}

// Verify an application by its ID and secret
func (c *Client) VerifyAppApps(ctx context.Context, path string, payload *AppCredentialsPayload, contentType string) (*http.Response, error) {
	req, err := c.NewVerifyAppAppsRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewVerifyAppAppsRequest create the request corresponding to the verifyApp action endpoint of the apps resource.
func (c *Client) NewVerifyAppAppsRequest(ctx context.Context, path string, payload *AppCredentialsPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}
