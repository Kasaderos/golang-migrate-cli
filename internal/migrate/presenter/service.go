package presenter

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

type Service interface {
	MigrateUp(context.Context) error
	MigrateDrop() error
	ListMigrations() error
	MigrateVersion() (string, error)
	GoToVersion(ctx context.Context, version uint) error
}

type service struct {
	db *migrate.Migrate
}

func New(db *migrate.Migrate) Service {
	return &service{db}
}

func (s *service) MigrateUp(ctx context.Context) error {
	return s.db.Up()
}

func (s *service) MigrateDrop() error {
	fmt.Println("ARE SURE ABOUT THAT??? y/N")
	answer := ""
	fmt.Scanf("%s\n", &answer)
	if answer == "y" {
		return s.db.Drop()
	}
	return nil
}

func (s *service) ListMigrations() error {
	return fmt.Errorf("not implemented")
}

func (s *service) MigrateVersion() (string, error) {
	version, dirty, err := s.db.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return err.Error(), nil
		}
		return "", err
	}
	if dirty {
		return fmt.Sprintf("%d dirty", version), nil
	}
	return fmt.Sprintf("%d clean", version), nil
}

func (s *service) GoToVersion(ctx context.Context, version uint) error {
	err := s.db.Migrate(version)
	if err == migrate.ErrNoChange {
		fmt.Println(err)
		return nil
	}
	return err
}
