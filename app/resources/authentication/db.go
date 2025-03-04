package authentication

import (
	"context"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fctx"
	"github.com/Southclaws/fault/ftag"
	"github.com/rs/xid"

	"github.com/Southclaws/storyden/app/resources/account"
	"github.com/Southclaws/storyden/internal/ent"
	model_account "github.com/Southclaws/storyden/internal/ent/account"
	"github.com/Southclaws/storyden/internal/ent/authentication"
)

type database struct {
	db *ent.Client
}

func New(db *ent.Client) Repository {
	return &database{db}
}

func (d *database) Create(ctx context.Context,
	id account.AccountID,
	service Service,
	identifier string,
	token string,
	metadata map[string]any,
) (*Authentication, error) {
	r, err := d.db.Authentication.Create().
		SetAccountID(xid.ID(id)).
		SetService(string(service)).
		SetIdentifier(identifier).
		SetToken(token).
		SetMetadata(metadata).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.AlreadyExists))
		}

		return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.Internal))
	}

	r, err = d.db.Authentication.
		Query().
		Where(authentication.ID(r.ID)).
		WithAccount().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.NotFound))
		}

		return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.Internal))
	}

	return FromModel(r), nil
}

func (d *database) LookupByIdentifier(ctx context.Context, service Service, identifier string) (*Authentication, bool, error) {
	r, err := d.db.Authentication.
		Query().
		Where(
			authentication.IdentifierEQ(identifier),
			authentication.ServiceEQ(string(service)),
		).
		WithAccount().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, false, nil
		}

		return nil, false, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.Internal))
	}

	return FromModel(r), true, nil
}

func (d *database) GetAuthMethods(ctx context.Context, id account.AccountID) ([]Authentication, error) {
	r, err := d.db.Authentication.
		Query().
		Where(authentication.HasAccountWith(model_account.IDEQ(xid.ID(id)))).
		All(ctx)
	if err != nil {
		return nil, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.Internal))
	}

	return FromModelMany(r), nil
}

func (d *database) IsEqual(ctx context.Context, id account.AccountID, identifier string, token string) (bool, error) {
	r, err := d.db.Authentication.
		Query().
		Where(
			authentication.HasAccountWith(model_account.IDEQ(xid.ID(id))),
			authentication.IdentifierEQ(identifier),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.NotFound))
		}

		return false, fault.Wrap(err, fctx.With(ctx), ftag.With(ftag.Internal))
	}

	return r.Token == token, nil
}
