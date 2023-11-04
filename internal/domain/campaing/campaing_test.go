package campaing

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaing X"
	content  = "Body... "
	contacts = []string{"email1@e.com", "email2@e.com", "email3@e.com"}
	fake     = faker.New()
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

func Test_NewCampaing_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing("", content, contacts)

	assert.Equal("name is required with min 3", err.Error())
}

func Test_NewCampaing_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(fake.Lorem().Text(33), content, contacts)

	assert.Equal("name is required with max 30", err.Error())
}
func Test_NewCampaing_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, "", contacts)

	assert.Equal("content is required with min 3", err.Error())
}

func Test_NewCampaing_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, fake.Lorem().Text(1040), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaing_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, nil)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaing_MustValidateContactsEmailInvalid(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaing(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
