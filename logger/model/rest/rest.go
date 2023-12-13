package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeServerError Code = 1000 + iota
	CodeNotFound
	CodeInvalidParams
	CodeAuth
	CodeUnknownError
)

var _httpStatus = map[Code]int{
	CodeNotFound:      http.StatusNotFound,
	CodeInvalidParams: http.StatusBadRequest,
	CodeAuth:          http.StatusUnauthorized,
	CodeServerError:   http.StatusInternalServerError,
}

type Code uint16

func (ec Code) HttpStatus() int {
	if hc, ok := _httpStatus[ec]; ok {
		return hc
	}
	return http.StatusInternalServerError
}

type Rest struct {
	Code  Code        `json:"code" binding:"required"`
	Msg   string      `json:"msg,omitempty"`
	Total int64       `json:"total,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	List  interface{} `json:"list,omitempty"`
}

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, Rest{Code: http.StatusOK, Data: v})
}

func Msg(ctx *gin.Context, msg string) {
	CodeMsg(ctx, http.StatusOK, msg)
}

func CodeMsg(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, Rest{Code: Code(code), Msg: msg})
}

func CodeData(ctx *gin.Context, code int, v interface{}) {
	ctx.JSON(http.StatusOK, Rest{Code: Code(code), Data: v})
}

func CodeDataMsg(ctx *gin.Context, code int, msg string, v interface{}) {
	ctx.JSON(http.StatusOK, Rest{Code: Code(code), Data: v, Msg: msg})
}

func Failed(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, Rest{Code: CodeInvalidParams, Msg: msg})
}

func OK(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Rest{Code: http.StatusOK, Msg: "ok"})
}

func Unauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, Rest{Code: CodeAuth, Msg: "token 无效"})
}

func Data(ctx *gin.Context, v interface{}, total int64) {
	ctx.JSON(http.StatusOK, Rest{Code: http.StatusOK, Data: v, Total: total})
}

func List(ctx *gin.Context, v interface{}, total int64) {
	ctx.JSON(http.StatusOK, Rest{Code: http.StatusOK, List: v, Total: total})
}

func (ae Rest) Error() string {
	return fmt.Sprintf("ERR-%05d: %s", ae.Code, ae.Msg)
}

func New(code Code, format string, args ...interface{}) Rest {
	return Rest{
		Code: code,
		Msg:  fmt.Sprintf(format, args...),
	}
}

func Wrap(code Code, err error) Rest {
	return Rest{
		Code: code,
		Msg:  err.Error(),
	}
}
