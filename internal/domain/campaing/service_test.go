package campaing

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/interna-lErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaing *Campaing) error {
	args := r.Called(campaing)
	return args.Error(0)
}

var (
	newCampaing = contract.NewCampaing{
		Name:    "Test Y",
		Content: "Body.. ..",
		Emails:  []string{"test1@email.com", "test2@email.com"},
	}

	service = Service{}
)

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock

	id, err := service.Create(newCampaing)

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaing(t *testing.T) {
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaing *Campaing) bool {
		if campaing.Name != newCampaing.Name || campaing.Content != newCampaing.Content || len(campaing.Contacts) != len(newCampaing.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Repository = repositoryMock
	service.Create(newCampaing)

	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaing{})

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock
	_, err := service.Create(newCampaing)
	assert.True(errors.Is(internalerrors.ErrInternal, err))

}
