package middlewares

import (
	"daimon/api"
	"daimon/pkg/ctx"
	"daimon/pkg/e"
	"daimon/pkg/jwt"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	Token        = "token"
	RefreshToken = "refreshtoken"
)

func JWT() ctx.HandlerFunc {
	return func(c *ctx.Context) {

		token := c.Request.Header.Get(Token)
		if token == "" {
			c.Abort(e.TokenError.Add("token is empty"))
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.Abort(e.TokenError.Add(err.Error()))
			return
		}
		// TODO 后续改造此逻辑
		if time.Unix(claims.ExpiresAt, 0).Add(time.Hour * 48).Before(time.Now()) {
			c.Abort(e.TokenError.Add("token expire more then 48h"))
			return
		}
		// ok set user id  and phone
		c.SetPhone(claims.Phone)
		c.SetUserID(claims.UserID)

		if claims.ExpiresAt < time.Now().Unix() {
			// TODO token timeout
			c.Set("refresh_token", jwt.GenToken(&api.User{UserId: claims.UserID, Phone: claims.Phone}, time.Hour*2))
		}

		c.Set("request_id", uuid.NewV4().String())

		c.Next()

	}
}
