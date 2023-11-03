package campaing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaing X"
	content  = "Body... "
	contacts = []string{"email1@e.com", "email2@e.com", "email3@e.com"}
)

func Test_NewCampaing_CreateCampaing(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

	//withou external lib testfy
	// if campaing.Id != "1" {
	// 	t.Errorf("expected 1")
	// } else if campaing.Name != name {
	// 	t.Errorf("expected correct name")
	// } else if len(campaing.Contacts) != len(contacts) {
	// 	t.Errorf("expected correct contacts")
	// }
}

func Test_NewCampaing_IdIsNotEmpty(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.NotEmpty(campaing.Id)
}

func Test_NewCampaing_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.Greater(campaing.CreatedOn, now)
}

func Test_NewCampaing_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaing_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaing_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, []string{})

	assert.Equal("contacat is required", err.Error())
}
