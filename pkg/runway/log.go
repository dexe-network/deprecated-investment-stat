package runway

import (
	"go.uber.org/zap"
)

func (r *Runway) Log() *zap.Logger {
	return r.log
}
