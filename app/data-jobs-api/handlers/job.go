package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/joshakeman/portal-demo/foundation/web"

	job "github.com/joshakeman/portal-demo/business/data"
)

type jobHandlers struct {
	db *sql.DB
}

func (h *jobHandlers) list(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	jobs, err := job.List(ctx, h.db)
	if err != nil {
		switch err {
		case job.ErrInvalidID:
			return err
			// return web.NewRequestError(err, http.StatusBadRequest)
		case job.ErrNotFound:
			return err
			// return web.NewRequestError(err, http.StatusNotFound)
		default:
			return err
		}
	}

	return web.Respond(ctx, w, jobs, http.StatusOK)
}
