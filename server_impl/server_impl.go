package impl

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/tk3413/tk-weight-calc/calculator"
	api "github.com/tk3413/tk-weight-calc/server_gen"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	logger *slog.Logger
}

func NewServer(opts ...Option) api.ServerInterface {
	s := &Server{}
	for _, o := range opts {
		o(s)
	}
	if s.logger == nil {
		s.logger = slog.Default()
	}
	return s
}

// (GET /weights)
func (s *Server) GetWeights(w http.ResponseWriter, r *http.Request, params api.GetWeightsParams) {
	s.logger.Debug("GetWeights endpoint called", "weight", params.Weight)

	resp := calculator.CalculateWeights(params.Weight)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func HandlerFromMux(si api.ServerInterface, m api.ServeMux) http.Handler {
	return api.HandlerFromMux(si, m)
}

type Option func(*Server)

func WithLogger(l *slog.Logger) Option {
	return func(s *Server) {
		s.logger = l
	}
}
