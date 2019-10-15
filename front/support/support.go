// Contextを操作するパッケージ
package support

import (
	"context"
)

type contextKeyUser struct{}

func AddUserToContext(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, contextKeyUser{}, email)
}
