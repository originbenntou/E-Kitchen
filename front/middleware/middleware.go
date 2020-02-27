package middleware

import (
	"github.com/originbenntou/E-Kitchen/front/session"
	"github.com/originbenntou/E-Kitchen/front/support"
	pbUser "github.com/originbenntou/E-Kitchen/proto/user"

	"log"
	"net/http"
	"time"
)

const loggingFmt = "Method:%s\tPath:%s\tElapsedTime:%s\tStatusCode:%d\n"

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}

		defer func() {
			log.Printf(loggingFmt,
				r.Method,
				r.URL.String(),
				time.Since(start),
				lrw.statusCode,
			)
		}()
		next.ServeHTTP(w, r)
	})
}

func NewAuthentication(
	userClient pbUser.UserServiceClient,
	sessionStore session.Store) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			sessionID := session.GetSessionIDFromRequest(r)
			// セッション有効ならユーザーID取得
			v, ok := sessionStore.Get(sessionID)
			if !ok {
				log.Println("session get failed")
				http.Redirect(w, r, "/signin", http.StatusFound)
				return
			}
			userID, ok := v.(uint64)
			if !ok {
				log.Println("sessionId convert into userId failed")
				http.Redirect(w, r, "/signin", http.StatusFound)
				return
			}

			ctx := r.Context()
			resp, err := userClient.FindUser(ctx, &pbUser.FindUserRequest{
				UserId: userID,
			})
			if err != nil {
				log.Println(err)
				http.Redirect(w, r, "/signin", http.StatusFound)
				return
			}
			// contextを介して他機能にユーザー情報を共有
			ctx = support.AddUserToContext(ctx, resp.Email)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
