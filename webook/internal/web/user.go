package web

import (
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	uSer        *service.UserService
}

func NewUserHandler(uSer *service.UserService) *UserHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&~])[A-Za-z\d$@$!%*#?&~]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
		uSer:        uSer,
	}
}

func (h *UserHandler) RegisterHandler(server *gin.Engine) {

	routerGroup := server.Group("/users")
	routerGroup.POST("/signup", h.SignUp)
	routerGroup.POST("/login", h.Login)
	routerGroup.POST("/profile", h.Profile)
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	type UserReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var userReq UserReq
	if err := ctx.BindJSON(&userReq); err != nil {
		ctx.String(http.StatusOK, "参数解析异常")
		return
	}
	ok, err := h.emailExp.MatchString(userReq.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}

	if !ok {
		ctx.String(http.StatusOK, "邮箱格式不正确")
		return
	}
	if userReq.Password != userReq.ConfirmPassword {
		ctx.String(http.StatusOK, "两次密码不同，重新输入")
		return
	}
	ok, err = h.passwordExp.MatchString(userReq.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}
	err = h.uSer.SignUp(ctx.Request.Context(), domain.User{
		Email:    userReq.Email,
		Password: userReq.Password,
	})
	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "邮箱冲突: %s", userReq.Email)
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.String(http.StatusOK, "注册成功:邮箱%s", userReq.Email)
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type loginRes struct {
		Email    string
		Password string
	}
	var res loginRes
	if err := ctx.Bind(&res); err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	user, err := h.uSer.Login(ctx.Request.Context(), res.Email, res.Password)
	if err == service.ErrEmailNotSignup {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	if err == service.ErrInvalidUserOrEmail {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	sess := sessions.Default(ctx)
	sess.Set("userId", user.Id)
	err = sess.Save()
	ctx.String(http.StatusOK, "登录成功")
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	userId := sess.Get("userId")
	if v, ok := userId.(int64); !ok {
		ctx.String(http.StatusOK, "用户名获取错误")
	} else {
		ctx.String(http.StatusOK, strconv.Itoa(int(v)))
	}
}
