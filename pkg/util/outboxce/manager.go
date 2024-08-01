//go:generate go run github.com/vektra/mockery/v2
package outboxce

import (
	"context"
	"database/sql"

	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/tink-crypto/tink-go/v2/tink"
)

type AEADFunc func(ob Outbox) (tink.AEAD, error)

type RelayFunc func(ctx context.Context, ce []event.Event) error

type Manager interface {
	Store(ctx context.Context, tx *sql.Tx, ob Outbox) (err error)
	StoreAsEncrypted(ctx context.Context, tx *sql.Tx, ob Outbox) (err error)

	Observe(ctx context.Context, relayFunc RelayFunc) (err error)
}
