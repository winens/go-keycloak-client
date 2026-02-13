package keycloak

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (c *Client) GetUser(userId uuid.UUID) (*UserRepresentation, error) {
	res, err := c.httpClient.Get(c.adminApiUrl + "/users/" + userId.String())
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrUnknown
	}

	var user UserRepresentation
	err = json.NewDecoder(res.Body).Decode(&user)
	return &user, err
}
