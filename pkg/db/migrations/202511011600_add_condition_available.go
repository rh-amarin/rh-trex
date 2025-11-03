package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func addConditionAvailable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202511011636",
		Migrate: func(tx *gorm.DB) error {
			// Create condition_available table
			// Stores individual status conditions from adapters
			createTableSQL := `
				CREATE TABLE IF NOT EXISTS condition_available (
					id VARCHAR(255) PRIMARY KEY,
					created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					deleted_at TIMESTAMPTZ NULL,
					cluster_id VARCHAR(255) NOT NULL,
					type VARCHAR(255) NOT NULL,
					status VARCHAR(255) NOT NULL,
					reason VARCHAR(255) NULL,
					message TEXT NULL,
					observed_generation INTEGER NOT NULL,
					condition_created_at TIMESTAMPTZ NOT NULL,
					condition_updated_at TIMESTAMPTZ NOT NULL
				);
			`

			if err := tx.Exec(createTableSQL).Error; err != nil {
				return err
			}

			// Create index on deleted_at for soft deletes
			if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_condition_available_deleted_at ON condition_available(deleted_at);").Error; err != nil {
				return err
			}

			// Create index on cluster_id for efficient lookups
			if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_condition_available_cluster_id ON condition_available(cluster_id) WHERE deleted_at IS NULL;").Error; err != nil {
				return err
			}

			// Create foreign key constraint to clusters table
			if err := tx.Exec(`
				ALTER TABLE condition_available 
				ADD CONSTRAINT fk_condition_available_cluster_id 
				FOREIGN KEY (cluster_id) REFERENCES clusters(id) 
				ON DELETE RESTRICT ON UPDATE RESTRICT;
			`).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Drop foreign key constraint first
			if err := tx.Exec("ALTER TABLE condition_available DROP CONSTRAINT IF EXISTS fk_condition_available_cluster_id;").Error; err != nil {
				return err
			}
			// Drop indexes
			if err := tx.Exec("DROP INDEX IF EXISTS idx_condition_available_cluster_id;").Error; err != nil {
				return err
			}
			if err := tx.Exec("DROP INDEX IF EXISTS idx_condition_available_deleted_at;").Error; err != nil {
				return err
			}
			// Drop table
			if err := tx.Exec("DROP TABLE IF EXISTS condition_available;").Error; err != nil {
				return err
			}
			return nil
		},
	}
}

