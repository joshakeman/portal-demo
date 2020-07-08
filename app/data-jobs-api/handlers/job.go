package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/joshakeman/portal-demo/foundation/web"

	"github.com/joshakeman/portal-demo/business/job"
)

type jobHandlers struct {
	db *sql.DB
}

func (h *jobHandlers) list(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// ctx, span := global.Tracer("service").Start(ctx, "handlers.product.retrieve")
	// defer span.End()

	jobs, err := job.List(ctx, h.db)
	if err != nil {
		switch err {
		case job.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		case job.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		default:
			return err
		}
	}

	return web.Respond(ctx, w, jobs, http.StatusOK)
}
