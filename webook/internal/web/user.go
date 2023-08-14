package web

import (
	"errors"
	"fmt"
	"github.com/TangTangHC/basic-go-study/webook/internal/domain"
	"github.com/TangTangHC/basic-go-study/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"unicode/utf8"
)

var (
	ErrMsgMoreThanMaxLen = errors.New("%s最多可输入%d个字")
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
	routerGroup.GET("/profile", h.Profile)
	routerGroup.POST("/edit", h.Edit)
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
	//sess := sessions.Default(ctx)
	//sess.Set("userId", user.Id)
	//sess.Options(sessions.Options{
	//	MaxAge: 30,
	//})
	//err = sess.Save()
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
		UId:       user.Id,
		UserAgent: ctx.Request.UserAgent(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signingString, err := token.SignedString([]byte("TyrmfzW2KnkH0HRfIH6lzd5XsQtrM31O"))
	ctx.Header("x-jwt-token", signingString)
	ctx.String(http.StatusOK, "登录成功")
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	//sess := sessions.Default(ctx)
	//sUserId := sess.Get("userId")
	//userId, ok := sUserId.(int64)
	//if !ok {
	//	ctx.String(http.StatusOK, "用户名获取错误")
	//	return
	//}
	token, _ := ctx.Get("token")
	v, ok := token.(*UserClaims)
	if !ok {
		// todo 监控这里，为什么解析失败
	}

	type userRes struct {
		Email     string `json:"email"`
		NikeName  string `json:"nikeName"`
		Birthday  string `json:"birthday"`
		Signature string `json:"signature"`
	}
	user, err := h.uSer.Profile(ctx, v.UId)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	ctx.JSON(http.StatusOK, &userRes{
		Email:     user.Email,
		NikeName:  user.NikeName,
		Birthday:  user.Birthday,
		Signature: user.Signature,
	})
}

func (h *UserHandler) Edit(ctx *gin.Context) {
	type userReq struct {
		NikeName  string `json:"nikeName"`
		Birthday  string `json:"birthday"`
		Signature string `json:"signature"`
	}
	var req userReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if l := utf8.RuneCountInString(req.NikeName); l > 20 {
		ctx.String(http.StatusOK, fmt.Sprintf(ErrMsgMoreThanMaxLen.Error(), "nikeName", 20))
		return
	}

	if l := utf8.RuneCountInString(req.Signature); l > 500 {
		ctx.String(http.StatusOK, fmt.Sprintf(ErrMsgMoreThanMaxLen.Error(), "signature", 500))
		return
	}

	if len(req.Birthday) > 0 {
		_, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			ctx.String(http.StatusOK, "birthday 格式错误")
			return
		}
	}

	sess := sessions.Default(ctx)
	sessVal := sess.Get("userId")
	userId, ok := sessVal.(int64)
	if !ok {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	err := h.uSer.Edit(ctx.Request.Context(), domain.User{
		Id:        userId,
		NikeName:  req.NikeName,
		Birthday:  req.Birthday,
		Signature: req.Signature,
	})
	if err != nil {
		ctx.String(http.StatusOK, "客户信息编辑失败")
		return
	}
	ctx.String(http.StatusOK, "客户信息编辑成功")
}

type UserClaims struct {
	jwt.RegisteredClaims
	UId       int64
	UserAgent string
}
