package delivery

//import (
//	"errors"
//	"fmt"
//	"github.com/go-chi/chi/v5/middleware"
//	"net/http"
//	"strings"
//	"time"
//)
//
//func Authorization(jwt auth.JWT) func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//			if whiteList(r.URL.Path) {
//				next.ServeHTTP(w, r)
//				return
//			}
//
//			bearerToken := r.Header.Get("Authorization")
//			splitted := strings.Split(bearerToken, "Bearer ")
//			if len(splitted) != 2 {
//				responseError(w, errs.New(model.ECodeAuthorization, "authorization header not found"))
//				return
//			}
//
//			claim, err := jwt.Verify(splitted[1])
//			if errors.Is(err, auth.ErrTokenExpired) {
//				responseError(w, errs.New(model.ECodeAuthorization, "token is expired"))
//				return
//			}
//
//			if err != nil {
//				responseError(w, errs.New(model.ECodeInternal, err.Error()))
//				return
//			}
//
//			ctx := context.WithValue(r.Context(), auth.JWTKey{}, claim)
//			next.ServeHTTP(w, r.WithContext(ctx))
//		})
//	}
//}
//
//func Logger(ins *telemetry.API) func(next http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		fn := func(w http.ResponseWriter, r *http.Request) {
//			t1 := time.Now()
//			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
//
//			next.ServeHTTP(ww, r)
//
//			ins.Logger().Info(r.Context(), "HTTP Request Response",
//				zap.String("http.method", r.Method),
//				zap.String("http.path", r.URL.Path),
//				zap.String("http.client.ip", r.RemoteAddr),
//				zap.Int("http.status", ww.Status()),
//				zap.String("http.time", fmt.Sprintf("%v", time.Since(t1))),
//			)
//		}
//
//		return http.HandlerFunc(fn)
//	}
//}
//
//func whiteList(path string) bool {
//	var list = []string{"/customer-alfa-account", "/user/verification", "/login",
//		"/internal/multiguna/va/callback", "/account-registration", "/otp/verify", "/otp/send",
//	}
//
//	if path == "/" {
//		return true
//	}
//
//	for _, v := range list {
//		if strings.Contains(path, v) {
//			return true
//		}
//	}
//
//	return false
//}
