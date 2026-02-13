package keycloak

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/subosito/gotenv"
)

func TestMain(m *testing.M) {
	gotenv.Load()
	os.Exit(m.Run())
}

func TestNewClient_Connection(t *testing.T) {
	client := NewClient(&Config{
		BaseURL:      os.Getenv("BASE_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Realm:        os.Getenv("REALM"),
	})

	userId := uuid.MustParse("719d4d81-0316-4131-bbfb-620cf70c61df")
	user, err := client.GetUser(userId)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("user found with username: %s", *user.Username)
}
