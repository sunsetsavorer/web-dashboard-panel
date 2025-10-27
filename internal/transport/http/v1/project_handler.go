package v1

import (
	"net/http"

	"github.com/sunsetsavorer/web-dashboard-panel/internal/interfaces"
	"github.com/sunsetsavorer/web-dashboard-panel/internal/router"
	"github.com/sunsetsavorer/web-dashboard-panel/internal/transport"
)

type ProjectHandler struct {
	*BaseHandler
	svc interfaces.ProjectService
}

func NewProjectHandler(
	baseHandler *BaseHandler,
	svc interfaces.ProjectService,
) *ProjectHandler {

	return &ProjectHandler{
		BaseHandler: baseHandler,
		svc:         svc,
	}
}

func (h *ProjectHandler) RegisterRoutes(router *router.Router) {

	router.HandleFunc("/api/v1/projects/", h.getList)
}

func (h *ProjectHandler) getList(w http.ResponseWriter, r *http.Request) {

	res, err := h.svc.GetList()
	if err != nil {
		h.WriteJsonResponse(
			w,
			transport.ErrorsResponse{
				Errors: map[string]string{
					"other": err.Error(),
				},
			},
			http.StatusBadRequest,
		)
		return
	}

	h.WriteJsonResponse(w, res, http.StatusOK)
}
