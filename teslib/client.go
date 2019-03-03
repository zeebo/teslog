package teslib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/zeebo/errs"
	"github.com/zeebo/mon"
)

var Unavailable = errs.Class("unavailable")

type Creds struct {
	Email    string
	Password string
	Token    *Token
}

type Client struct {
	Client *http.Client
	Creds
}

func (c *Client) client() *http.Client {
	if c.Client != nil {
		return c.Client
	}
	return http.DefaultClient
}

func (c *Client) Post(path string, request, response interface{}) (err error) {
	defer mon.Start().Stop(&err)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(request); err != nil {
		return errs.Wrap(err)
	}
	req, err := http.NewRequest("POST", makeURL(path), bytes.NewReader(buf.Bytes()))
	if err != nil {
		return errs.Wrap(err)
	}
	return c.Do(req, response)
}

func (c *Client) Get(path string, response interface{}) (err error) {
	defer mon.Start().Stop(&err)

	req, err := http.NewRequest("GET", makeURL(path), nil)
	if err != nil {
		return errs.Wrap(err)
	}
	return c.Do(req, response)
}

func (c *Client) Do(req *http.Request, response interface{}) (err error) {
	defer mon.Start().Stop(&err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", UserAgent)
	if c.Token != nil {
		req.Header.Set("Authorization", "Bearer "+c.Token.AccessToken)
	}

	resp, err := c.client().Do(req)
	if err != nil {
		return errs.Wrap(err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusRequestTimeout:
		return Unavailable.New("%s", req.URL.String())
	case http.StatusOK:
	default:
		body, _ := ioutil.ReadAll(resp.Body)
		return errs.New("%s: status code: %d: %s",
			req.URL.String(), resp.StatusCode, body)
	}

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return errs.Wrap(err)
	}

	return nil
}

func (c *Client) UpdateCreds() (err error) {
	defer mon.Start().Stop(&err)

	if c.Token == nil || c.Token.Expired() {
		return errs.Wrap(c.acquireToken())
	} else if c.Token.ShouldRefresh() {
		return errs.Wrap(c.refreshToken())
	}
	return nil
}

func (c *Client) acquireToken() (err error) {
	defer mon.Start().Stop(&err)

	req := AuthRequest{
		GrantType:    "password",
		ClientId:     ClientId,
		ClientSecret: ClientSecret,
		Email:        c.Email,
		Password:     c.Password,
	}
	var resp AuthResponse

	if err := c.Post(AuthURL, req, &resp); err != nil {
		return errs.Wrap(err)
	}

	c.Token = &Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		Expires:      time.Now().Add(time.Second * time.Duration(resp.ExpiresIn)),
	}

	return nil
}

func (c *Client) refreshToken() (err error) {
	defer mon.Start().Stop(&err)

	req := RefreshRequest{
		GrantType:    "refresh_token",
		ClientId:     ClientId,
		ClientSecret: ClientSecret,
		RefreshToken: c.Token.RefreshToken,
	}
	var resp RefreshResponse

	if err := c.Post(AuthURL, req, &resp); err != nil {
		return errs.Wrap(err)
	}

	c.Token = &Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		Expires:      time.Now().Add(time.Second * time.Duration(resp.ExpiresIn)),
	}

	return nil
}

func (c *Client) Vehicles() (vr *VehiclesResponse, err error) {
	defer mon.Start().Stop(&err)

	if err := c.UpdateCreds(); err != nil {
		return nil, errs.Wrap(err)
	}

	var resp VehiclesResponse
	if err := c.Get(VehiclesURL, &resp); err != nil {
		return nil, errs.Wrap(err)
	}
	return &resp, nil
}

func (c *Client) Data(id int64) (dr *DataResponse, err error) {
	defer mon.Start().Stop(&err)

	if err := c.UpdateCreds(); err != nil {
		return nil, errs.Wrap(err)
	}

	var resp struct {
		Response DataResponse `json:"response"`
	}
	if err := c.Get(fmt.Sprintf(DataURL, id), &resp); err != nil {
		return nil, errs.Wrap(err)
	}
	return &resp.Response, nil
}

func (c *Client) DriveState(id int64) (dsr *DriveStateResponse, err error) {
	defer mon.Start().Stop(&err)

	if err := c.UpdateCreds(); err != nil {
		return nil, errs.Wrap(err)
	}

	var resp struct {
		Response DriveStateResponse `json:"response"`
	}
	if err := c.Get(fmt.Sprintf(DriveStateURL, id), &resp); err != nil {
		return nil, errs.Wrap(err)
	}
	return &resp.Response, nil
}

func (c *Client) ClimateState(id int64) (csr *ClimateStateResponse, err error) {
	defer mon.Start().Stop(&err)

	if err := c.UpdateCreds(); err != nil {
		return nil, errs.Wrap(err)
	}

	var resp struct {
		Response ClimateStateResponse `json:"response"`
	}
	if err := c.Get(fmt.Sprintf(ClimateStateURL, id), &resp); err != nil {
		return nil, errs.Wrap(err)
	}
	return &resp.Response, nil
}

func (c *Client) ChargeState(id int64) (csr *ChargeStateResponse, err error) {
	defer mon.Start().Stop(&err)

	if err := c.UpdateCreds(); err != nil {
		return nil, errs.Wrap(err)
	}

	var resp struct {
		Response ChargeStateResponse `json:"response"`
	}
	if err := c.Get(fmt.Sprintf(ChargeStateURL, id), &resp); err != nil {
		return nil, errs.Wrap(err)
	}
	return &resp.Response, nil
}

func (c *Client) GUISettings(id int64) (gsr *GUISettingsResponse, err error) {
	defer mon.Start().Stop(&err)

	if err := c.UpdateCreds(); err != nil {
		return nil, errs.Wrap(err)
	}

	var resp struct {
		Response GUISettingsResponse `json:"response"`
	}
	if err := c.Get(fmt.Sprintf(GUISettingsURL, id), &resp); err != nil {
		return nil, errs.Wrap(err)
	}
	return &resp.Response, nil
}

func (c *Client) VehicleState(id int64) (vsr *VehicleStateResponse, err error) {
	defer mon.Start().Stop(&err)

	if err := c.UpdateCreds(); err != nil {
		return nil, errs.Wrap(err)
	}

	var resp struct {
		Response VehicleStateResponse `json:"response"`
	}
	if err := c.Get(fmt.Sprintf(VehicleStateURL, id), &resp); err != nil {
		return nil, errs.Wrap(err)
	}
	return &resp.Response, nil
}

func (c *Client) VehicleConfig(id int64) (vcr *VehicleConfigResponse, err error) {
	defer mon.Start().Stop(&err)

	if err := c.UpdateCreds(); err != nil {
		return nil, errs.Wrap(err)
	}

	var resp struct {
		Response VehicleConfigResponse `json:"response"`
	}
	if err := c.Get(fmt.Sprintf(VehicleConfigURL, id), &resp); err != nil {
		return nil, errs.Wrap(err)
	}
	return &resp.Response, nil
}
