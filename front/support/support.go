// Contextを操作するパッケージ
package support

import (
	"context"
)

type contextKeyUser struct{}

func AddUserToContext(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, contextKeyUser{}, email)
}

func GetUserFromContext(ctx context.Context) string {
	u := ctx.Value(contextKeyUser{})
	email, ok := u.(string)
	if !ok {
		return ""
	}
	return email
}
