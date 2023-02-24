package tunnel

import (
	"context"

	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/pkg/errors"
)

func (ic *ContainerEngine) AutoUpdate(ctx context.Context, options entities.AutoUpdateOptions) (*entities.AutoUpdateReport, []error) {
	return nil, []error{errors.New("not implemented")}
}
