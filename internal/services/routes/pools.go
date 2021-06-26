package services_routes_pools

import (
	"context"
	"dex-trades-parser/internal/models"
	"dex-trades-parser/internal/storage"
	"dex-trades-parser/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type PoolService struct {
	ctx    context.Context
	cancel func()
	log    *zap.Logger
	st     *storage.Storage
}

func NewPoolService(
	ctx context.Context,
	cancel func(),
	log *zap.Logger,
	st *storage.Storage,
) (s *PoolService) {
	poolService := &PoolService{
		ctx:    ctx,
		cancel: cancel,
		log:    log,
		st:     st,
	}
	return poolService
}

func (p *PoolService) GetAll(c *gin.Context) {
	var pools []models.Pool
	if err := p.st.DB.Find(&pools).Error; err != nil {
		response.Error(c, http.StatusBadRequest, response.E{
			Code:    response.InvalidJSONBody,
			Message: "invalid request",
		})
		return
	}

	response.Success(c, http.StatusOK, response.S{Data: pools})

}
