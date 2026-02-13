package keycloak

import (
	"time"

	"github.com/google/uuid"
)

type UserRepresentation struct {
	ID               uuid.UUID           `json:"id"`
	Username         *string             `json:"username,omitempty"`
	FirstName        *string             `json:"firstName,omitempty"`
	LastName         *string             `json:"lastName,omitempty"`
	Email            *string             `json:"email,omitempty"`
	EmailVerified    *bool               `json:"emailVerified,omitempty"`
	Enabled          *bool               `json:"enabled,omitempty"`
	Attributes       map[string][]string `json:"attributes,omitempty"`
	CreatedTimestamp int64               `json:"createdTimestamp,omitempty"` // in unix miliseconds
	TOTP             *bool               `json:"totp,omitempty"`
	RequiredActions  []string            `json:"requiredActions,omitempty"`
	NotBefore        *int32              `json:"notBefore,omitempty"`
	Access           map[string]bool     `json:"access,omitempty"`
}

func (u *UserRepresentation) CreatedAt() time.Time {
	return time.UnixMilli(u.CreatedTimestamp)
}
