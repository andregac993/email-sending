package campaign

import (
	"errors"
	"github.com/rs/xid"
	"time"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedAt time.Time
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("nome é obrigatório")
	} else if content == "" {
		return nil, errors.New("conteudo é obrigatório")
	} else if len(emails) == 0 {
		return nil, errors.New("e-mail é obrigatório")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}, nil
}
