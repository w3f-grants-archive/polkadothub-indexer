package system_event

import (
	"github.com/figment-networks/polkadothub-indexer/client"
	"github.com/figment-networks/polkadothub-indexer/model"
	"github.com/figment-networks/polkadothub-indexer/store"
	"github.com/figment-networks/polkadothub-indexer/types"
	"github.com/figment-networks/polkadothub-indexer/usecase/http"
	"github.com/figment-networks/polkadothub-indexer/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	_ types.HttpHandler = (*getForAddressHttpHandler)(nil)
)

type getForAddressHttpHandler struct {
	client *client.Client

	useCase *getForAddressUseCase

	systemEventDb store.SystemEvents
}

func NewGetForAddressHttpHandler(c *client.Client, systemEventDb store.SystemEvents) *getForAddressHttpHandler {
	return &getForAddressHttpHandler{
		client: c,

		systemEventDb: systemEventDb,
	}
}

type GetForAddressRequest struct {
	Address string                 `uri:"address" binding:"required"`
	After   *int64                 `form:"after" binding:"-"`
	Kind    *model.SystemEventKind `form:"kind" binding:"-"`
}

func (h *getForAddressHttpHandler) Handle(c *gin.Context) {
	var req GetForAddressRequest
	if err := c.ShouldBindUri(&req); err != nil {
		logger.Error(err)
		http.BadRequest(c, errors.New("invalid address"))
		return
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error(err)
		http.BadRequest(c, errors.New("invalid kind or/and after"))
		return
	}

	resp, err := h.getUseCase().Execute(req.Address, req.After, req.Kind)
	if err != nil {
		logger.Error(err)
		http.ServerError(c, err)
		return
	}

	http.JsonOK(c, resp)
}

func (h *getForAddressHttpHandler) getUseCase() *getForAddressUseCase {
	if h.useCase == nil {
		h.useCase = NewGetForAddressUseCase(h.systemEventDb)
	}
	return h.useCase
}