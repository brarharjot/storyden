package password

import (
	"context"
	"net/mail"
	"strings"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/fmsg"
	"github.com/Southclaws/fault/ftag"
	"github.com/alexedwards/argon2id"
	"github.com/pkg/errors"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/app/resources/authentication"
)

const AuthServiceName = `password`

var (
	ErrAccountAlreadyExists = errors.New("account already exists")
	ErrPasswordMismatch     = errors.New("password mismatch")
	ErrNotFound             = errors.New("account not found")
)

type Password struct {
	auth    authentication.Repository
	account account.Repository
}

func NewBasicAuth(auth authentication.Repository, account account.Repository) *Password {
	return &Password{auth, account}
}

func (b *Password) Register(ctx context.Context, identifier string, password string) (*account.Account, error) {
	addr, err := mail.ParseAddress(identifier)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.InvalidArgument))
	}

	handle := strings.Split(addr.Address, "@")[0]

	_, exists, err := b.auth.LookupByIdentifier(ctx, AuthServiceName, identifier)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), fmsg.With("failed to get account"))
	}

	if exists {
		return nil, fault.Wrap(ErrAccountAlreadyExists, fctx.With(ctx), ftag.With(ftag.AlreadyExists))
	}

	account, err := b.account.Create(ctx, handle)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), fmsg.With("failed to create account"))
	}

	hashed, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), fmsg.With("failed to create secure password hash"))
	}

	_, err = b.auth.Create(ctx, account.ID, AuthServiceName, identifier, string(hashed), nil)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), fmsg.With("failed to create account authentication instance"))
	}

	return account, nil
}

func (b *Password) Login(ctx context.Context, identifier string, password string) (*account.Account, error) {
	a, exists, err := b.auth.LookupByIdentifier(ctx, AuthServiceName, identifier)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), fmsg.With("failed to get account"))
	}

	if !exists {
		return nil, fault.Wrap(ErrNotFound, fctx.With(ctx), ftag.With(ftag.NotFound))
	}

	match, _, err := argon2id.CheckHash(password, a.Token)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), fmsg.With("failed to compare secure password hash"))
	}

	if !match {
		return nil, fault.Wrap(ErrPasswordMismatch, fctx.With(ctx), ftag.With(ftag.Unauthenticated))
	}

	return &a.Account, nil
}
