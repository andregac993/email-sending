package campaign

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	name     = "Nome da Campanha"
	content  = "Corpo da campanha"
	contacts = []string{"example@example.com", "example2@example.com"}
)

func Test_NewCampaign_Create_Campaign(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.Equal(campaign.Name, name)
	assertions.Equal(campaign.Content, content)
	assertions.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_ID_Not_Null(t *testing.T) {
	assertions := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedAt_Must_Be_Now(t *testing.T) {
	assertions := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assertions.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_Must_Validate_Name(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	if assertions.NotNil(err, "Expected an error when campaign name is empty") {
		assertions.Equal("nome é obrigatório", err.Error())
	}
}

func Test_NewCampaign_Must_Validate_Content(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	if assertions.NotNil(err, "Expected an error when campaign content is empty") {
		assertions.Equal("conteudo é obrigatório", err.Error())
	}
}

func Test_NewCampaign_Must_Validate_Contact(t *testing.T) {
	assertions := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	if assertions.NotNil(err, "Expected an error when contacts are empty") {
		assertions.Equal("e-mail é obrigatório", err.Error())
	}
}
