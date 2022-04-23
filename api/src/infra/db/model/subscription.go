// Code generated by entc, DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Southclaws/storyden/api/src/infra/db/model/subscription"
	"github.com/google/uuid"
)

// Subscription is the model entity for the Subscription schema.
type Subscription struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// RefersType holds the value of the "refers_type" field.
	RefersType string `json:"refers_type,omitempty"`
	// RefersTo holds the value of the "refers_to" field.
	RefersTo string `json:"refers_to,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubscriptionQuery when eager-loading is set.
	Edges                     SubscriptionEdges `json:"edges"`
	notification_subscription *uuid.UUID
	user_subscriptions        *uuid.UUID
}

// SubscriptionEdges holds the relations/edges for other nodes in the graph.
type SubscriptionEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e SubscriptionEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e SubscriptionEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[1] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subscription) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case subscription.FieldRefersType, subscription.FieldRefersTo:
			values[i] = new(sql.NullString)
		case subscription.FieldDeleteTime, subscription.FieldCreateTime, subscription.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case subscription.FieldID:
			values[i] = new(uuid.UUID)
		case subscription.ForeignKeys[0]: // notification_subscription
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case subscription.ForeignKeys[1]: // user_subscriptions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subscription", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subscription fields.
func (s *Subscription) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subscription.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case subscription.FieldRefersType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refers_type", values[i])
			} else if value.Valid {
				s.RefersType = value.String
			}
		case subscription.FieldRefersTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field refers_to", values[i])
			} else if value.Valid {
				s.RefersTo = value.String
			}
		case subscription.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				s.DeleteTime = value.Time
			}
		case subscription.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case subscription.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case subscription.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field notification_subscription", values[i])
			} else if value.Valid {
				s.notification_subscription = new(uuid.UUID)
				*s.notification_subscription = *value.S.(*uuid.UUID)
			}
		case subscription.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_subscriptions", values[i])
			} else if value.Valid {
				s.user_subscriptions = new(uuid.UUID)
				*s.user_subscriptions = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Subscription entity.
func (s *Subscription) QueryUser() *UserQuery {
	return (&SubscriptionClient{config: s.config}).QueryUser(s)
}

// QueryNotifications queries the "notifications" edge of the Subscription entity.
func (s *Subscription) QueryNotifications() *NotificationQuery {
	return (&SubscriptionClient{config: s.config}).QueryNotifications(s)
}

// Update returns a builder for updating this Subscription.
// Note that you need to call Subscription.Unwrap() before calling this method if this Subscription
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subscription) Update() *SubscriptionUpdateOne {
	return (&SubscriptionClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Subscription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subscription) Unwrap() *Subscription {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("model: Subscription is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subscription) String() string {
	var builder strings.Builder
	builder.WriteString("Subscription(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", refers_type=")
	builder.WriteString(s.RefersType)
	builder.WriteString(", refers_to=")
	builder.WriteString(s.RefersTo)
	builder.WriteString(", delete_time=")
	builder.WriteString(s.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Subscriptions is a parsable slice of Subscription.
type Subscriptions []*Subscription

func (s Subscriptions) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
