package campaing

import (
	internalerrors "emailn/internal/interna-lErrors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaing struct {
	Id        string    `validate:"required"`
	Name      string    `validate:"min=3,max=30"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=3,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
}

//init func, "constructor in Java"

func NewCampaing(name string, content string, emails []string) (*Campaing, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	campaing := &Campaing{
		Id:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

	err := internalerrors.ValidateStruct(campaing)

	if err == nil {
		return campaing, nil
	}
	return nil, err
}
