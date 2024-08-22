package middle

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net/http"
	"slices"
	"strings"
	"test/common"
)

// 权限校验中间件
var ErrInvalidToken = common.NewCodeError(400, common.LoginOutOfDate)
var ErrInvalidResource = common.NewCodeError(-1, common.ResourceNotAuth)
var ErrInvalidAuth = common.NewCodeError(-1, common.UserAuthNotEnough)
var whitePath = []string{
	"/admin-user/user/login",
}

type AuthMiddleware struct {
	Rds redis.Client
}

func NewAuthMiddleware(rds redis.Client) *AuthMiddleware {
	return &AuthMiddleware{Rds: rds}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		urlPath := r.RequestURI
		method := r.Method

		//白名单放行
		if slices.Contains(whitePath, urlPath) {
			next.ServeHTTP(w, r)
			return
		}
		if len(token) == 0 {
			common.Response(w, nil, ErrInvalidToken)
			return
		}
		//get请求放行
		if strings.EqualFold(method, "get") {
			next.ServeHTTP(w, r)
			return
		}
		userInfo, err := common.GetJwtUserInfo(r.Context().Value("userInfo"))
		if err != nil {
			common.Response(w, nil, err)
			return
		}
		// 判断redis中的token跟请求的token是否一致，不一致的话返回报错信息
		//redis := database.GetRedisDB()
		key, _ := m.Rds.Get(context.Background(), "USER_JWT:"+string(userInfo.UserId)).Result()
		if !strings.EqualFold(key, token) {
			common.Response(w, nil, ErrInvalidToken)
			return
		}
		roleNames := userInfo.RoleNames
		//没有资源权限的话直接返回错误
		if len(roleNames) < 1 {
			common.Response(w, nil, ErrInvalidAuth)
			return
		}

		//判断用户角色是不是高级运营以上
		var success bool
		for _, name := range roleNames {
			if strings.EqualFold(name, "超级管理员") || strings.EqualFold(name, "高级运营人员") {
				success = true
				break
			}
		}
		//未发现资源存在的情况，直接返回错误信息
		if !success {
			common.Response(w, nil, ErrInvalidAuth)
			return
		}
		/*pathValue := r.Context().Value("paths")
		var paths []string
		if jsonPath, ok := pathValue.([]string); ok {
			paths = jsonPath
		} else {
			response.Response(w, nil, ErrInvalidToken)
			return
		}
		//没有资源权限的话直接返回错误
		if len(paths) < 1 {
			response.Response(w, nil, ErrInvalidResource)
			return
		}
		//判断请求的路径是否再资源权限内
		var newstr string
		var success bool
		for _, path := range paths {
			newstr = strings.ReplaceAll(newstr, path, "**")
			if strings.Contains(path, "**") && strings.HasPrefix(urlPath, newstr) {
				success = true
				break
			} else if strings.Contains(path, urlPath) {
				success = true
				break
			}
		}
		//未发现资源存在的情况，直接返回错误信息
		if !success {
			response.Response(w, nil, ErrInvalidResource)
			return
		}*/
		next.ServeHTTP(w, r)
		//// 将鉴权信息添加到请求上下文中
		//ctx := context.WithValue(r.Context(), "token", token)
		//next.ServeHTTP(w, r.WithContext(ctx))
	}
}
