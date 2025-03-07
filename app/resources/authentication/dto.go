package authentication

import (
	"github.com/samber/lo"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/internal/ent"
)

type Service string

type Authentication struct {
	Account    account.Account `json:"account"`
	Service    Service         `json:"service"`
	Identifier string          `json:"identifier"`
	Token      string          `json:"-"`
	Metadata   interface{}     `json:"metadata,omitempty"`
}

func FromModel(m *ent.Authentication) *Authentication {
	return &Authentication{
		Account:    *account.FromModel(*m.Edges.Account),
		Service:    Service(m.Service),
		Identifier: m.Identifier,
		Token:      m.Token,
		Metadata:   m.Metadata,
	}
}

func FromModelMany(m []*ent.Authentication) []Authentication {
	return lo.Map(m, func(t *ent.Authentication, i int) Authentication {
		return *FromModel(t)
	})
}
