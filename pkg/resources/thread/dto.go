package thread

import (
	"time"

	"4d63.com/optional"
	"github.com/Southclaws/dt"

	"github.com/Southclaws/storyden/internal/infrastructure/db/model"
	"github.com/Southclaws/storyden/internal/utils"
	"github.com/Southclaws/storyden/pkg/resources/account"
	"github.com/Southclaws/storyden/pkg/resources/category"
	"github.com/Southclaws/storyden/pkg/resources/post"
	"github.com/Southclaws/storyden/pkg/resources/react"
)

type Thread struct {
	ID        post.PostID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt optional.Optional[time.Time]

	Title    string
	Slug     string
	Short    string
	Pinned   bool
	Author   AuthorRef
	Tags     []string
	Category category.Category
	Posts    []*post.Post
	Reacts   []*react.React
}

type AuthorRef struct {
	ID   account.AccountID
	Name string
}

const Name = "Thread"

func (*Thread) GetRole() string { return Name }

func (*Thread) GetResourceName() string { return Name }

func FromModel(m *model.Post) *Thread {
	var posts []*post.Post

	if p, err := m.Edges.PostsOrErr(); err == nil && len(p) > 0 {
		posts = dt.Map(p, post.FromModel)
	}

	return &Thread{
		ID:        post.PostID(m.ID),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: optional.OfPtr(m.DeletedAt),

		Title:  m.Title,
		Slug:   m.Slug,
		Short:  m.Short,
		Pinned: m.Pinned,
		Author: AuthorRef{
			ID:   account.AccountID(m.Edges.Author.ID),
			Name: m.Edges.Author.Name,
		},
		Tags:     dt.Map(m.Edges.Tags, func(t *model.Tag) string { return t.Name }),
		Category: utils.Deref(category.FromModel(m.Edges.Category)),
		Posts:    posts,
		Reacts:   dt.Map(m.Edges.Reacts, react.FromModel),
	}
}
