package notifier

import "context"

type Notifier interface {
	Send(context.Context, string, string) error
}
