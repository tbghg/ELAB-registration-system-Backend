package comment

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/thread/comment"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func deleteCommentById(ctx *gin.Context) {
	request := comment.DeleteRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	err = comment.Delete(ctx, &request)
	if err != nil {
		slog.Error("error in comment.Delete, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{Ok: true})
}