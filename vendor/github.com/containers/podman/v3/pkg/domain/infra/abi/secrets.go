package abi

import (
	"context"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/containers/common/pkg/secrets"
	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/pkg/errors"
)

func (ic *ContainerEngine) SecretCreate(ctx context.Context, name string, reader io.Reader, options entities.SecretCreateOptions) (*entities.SecretCreateReport, error) {
	data, _ := ioutil.ReadAll(reader)
	secretsPath := ic.Libpod.GetSecretsStorageDir()
	manager, err := secrets.NewManager(secretsPath)
	if err != nil {
		return nil, err
	}
	driverOptions := make(map[string]string)

	if options.Driver == "" {
		options.Driver = "file"
	}
	if options.Driver == "file" {
		driverOptions["path"] = filepath.Join(secretsPath, "filedriver")
	}
	secretID, err := manager.Store(name, data, options.Driver, driverOptions)
	if err != nil {
		return nil, err
	}
	return &entities.SecretCreateReport{
		ID: secretID,
	}, nil
}

func (ic *ContainerEngine) SecretInspect(ctx context.Context, nameOrIDs []string) ([]*entities.SecretInfoReport, []error, error) {
	secretsPath := ic.Libpod.GetSecretsStorageDir()
	manager, err := secrets.NewManager(secretsPath)
	if err != nil {
		return nil, nil, err
	}
	errs := make([]error, 0, len(nameOrIDs))
	reports := make([]*entities.SecretInfoReport, 0, len(nameOrIDs))
	for _, nameOrID := range nameOrIDs {
		secret, err := manager.Lookup(nameOrID)
		if err != nil {
			if errors.Cause(err).Error() == "no such secret" {
				errs = append(errs, err)
				continue
			} else {
				return nil, nil, errors.Wrapf(err, "error inspecting secret %s", nameOrID)
			}
		}
		report := &entities.SecretInfoReport{
			ID:        secret.ID,
			CreatedAt: secret.CreatedAt,
			UpdatedAt: secret.CreatedAt,
			Spec: entities.SecretSpec{
				Name: secret.Name,
				Driver: entities.SecretDriverSpec{
					Name: secret.Driver,
				},
			},
		}
		reports = append(reports, report)
	}

	return reports, errs, nil
}

func (ic *ContainerEngine) SecretList(ctx context.Context) ([]*entities.SecretInfoReport, error) {
	secretsPath := ic.Libpod.GetSecretsStorageDir()
	manager, err := secrets.NewManager(secretsPath)
	if err != nil {
		return nil, err
	}
	secretList, err := manager.List()
	if err != nil {
		return nil, err
	}
	report := make([]*entities.SecretInfoReport, 0, len(secretList))
	for _, secret := range secretList {
		reportItem := entities.SecretInfoReport{
			ID:        secret.ID,
			CreatedAt: secret.CreatedAt,
			UpdatedAt: secret.CreatedAt,
			Spec: entities.SecretSpec{
				Name: secret.Name,
				Driver: entities.SecretDriverSpec{
					Name:    secret.Driver,
					Options: secret.DriverOptions,
				},
			},
		}
		report = append(report, &reportItem)
	}
	return report, nil
}

func (ic *ContainerEngine) SecretRm(ctx context.Context, nameOrIDs []string, options entities.SecretRmOptions) ([]*entities.SecretRmReport, error) {
	var (
		err      error
		toRemove []string
		reports  = []*entities.SecretRmReport{}
	)
	secretsPath := ic.Libpod.GetSecretsStorageDir()
	manager, err := secrets.NewManager(secretsPath)
	if err != nil {
		return nil, err
	}
	toRemove = nameOrIDs
	if options.All {
		allSecrs, err := manager.List()
		if err != nil {
			return nil, err
		}
		for _, secr := range allSecrs {
			toRemove = append(toRemove, secr.ID)
		}
	}
	for _, nameOrID := range toRemove {
		deletedID, err := manager.Delete(nameOrID)
		if err == nil || errors.Cause(err).Error() == "no such secret" {
			reports = append(reports, &entities.SecretRmReport{
				Err: err,
				ID:  deletedID,
			})
			continue
		} else {
			return nil, err
		}
	}

	return reports, nil
}
