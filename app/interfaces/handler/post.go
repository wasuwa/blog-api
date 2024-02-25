package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wasuwa/blog-api/app/usecase"
)

// PostHandler 投稿のハンドラーインターフェース
type PostHandler interface {
	HandlePostFindByID(ctx echo.Context) error
	Routes(e *echo.Echo)
}

type postHandler struct {
	usecase usecase.PostUseCase
}

// NewPostHandler postHandlerを初期化します
func NewPostHandler(pu usecase.PostUseCase) PostHandler {
	return &postHandler{usecase: pu}
}

// Post 投稿
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// HandlePostFindByID IDに紐づく投稿を取得します
func (ph *postHandler) HandlePostFindByID(ctx echo.Context) error {
	ctx.QueryParam("id")
	if _, err := ph.usecase.FindByID(context.TODO(), 0); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	post := &Post{
		ID:    "",
		Title: "title",
		Body:  "body",
	}
	return ctx.JSON(http.StatusOK, post)
}

// Routes ルーティング
func (ph *postHandler) Routes(e *echo.Echo) {
	e.GET("/posts/:id", ph.HandlePostFindByID)
}
