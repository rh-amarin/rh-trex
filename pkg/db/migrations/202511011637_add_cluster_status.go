package migrations

// Migrations should NEVER use types from other packages. Types can change
// and then migrations run on a _new_ database will fail or behave unexpectedly.
// Instead of importing types, always re-create the type in the migration, as
// is done here, even though the same type is defined in pkg/api

import (
	"gorm.io/gorm"

	"github.com/go-gormigrate/gormigrate/v2"
)

func addClusterStatuses() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "202511011637",
		Migrate: func(tx *gorm.DB) error {
			// Create cluster_statuses table
			// Stores status history/snapshots aggregated from all adapters
			createTableSQL := `
				CREATE TABLE IF NOT EXISTS cluster_statuses (
					id VARCHAR(255) PRIMARY KEY,
					created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
					deleted_at TIMESTAMPTZ NULL,
					cluster_id VARCHAR(255) NOT NULL,
					phase VARCHAR(255) NOT NULL,
					last_transition_time TIMESTAMPTZ NOT NULL,
					observed_generation INTEGER NOT NULL,
					status_updated_at TIMESTAMPTZ NOT NULL
				);
			`

			if err := tx.Exec(createTableSQL).Error; err != nil {
				return err
			}

			// Create index on deleted_at for soft deletes
			if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_cluster_statuses_deleted_at ON cluster_statuses(deleted_at);").Error; err != nil {
				return err
			}

			// Create index on cluster_id for efficient lookups
			if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_cluster_statuses_cluster_id ON cluster_statuses(cluster_id) WHERE deleted_at IS NULL;").Error; err != nil {
				return err
			}

			// Create index on cluster_id and created_at for history queries
			if err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_cluster_statuses_cluster_created ON cluster_statuses(cluster_id, created_at DESC) WHERE deleted_at IS NULL;").Error; err != nil {
				return err
			}

			// Create foreign key constraint to clusters table
			// Drop first if exists to make migration idempotent
			if err := tx.Exec("ALTER TABLE cluster_statuses DROP CONSTRAINT IF EXISTS fk_cluster_statuses_cluster_id;").Error; err != nil {
				return err
			}
			if err := tx.Exec(`
				ALTER TABLE cluster_statuses 
				ADD CONSTRAINT fk_cluster_statuses_cluster_id 
				FOREIGN KEY (cluster_id) REFERENCES clusters(id) 
				ON DELETE RESTRICT ON UPDATE RESTRICT;
			`).Error; err != nil {
				return err
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			// Drop foreign key constraint first
			if err := tx.Exec("ALTER TABLE cluster_statuses DROP CONSTRAINT IF EXISTS fk_cluster_statuses_cluster_id;").Error; err != nil {
				return err
			}
			// Drop indexes
			if err := tx.Exec("DROP INDEX IF EXISTS idx_cluster_statuses_cluster_created;").Error; err != nil {
				return err
			}
			if err := tx.Exec("DROP INDEX IF EXISTS idx_cluster_statuses_cluster_id;").Error; err != nil {
				return err
			}
			if err := tx.Exec("DROP INDEX IF EXISTS idx_cluster_statuses_deleted_at;").Error; err != nil {
				return err
			}
			// Drop table
			if err := tx.Exec("DROP TABLE IF EXISTS cluster_statuses;").Error; err != nil {
				return err
			}
			return nil
		},
	}
}

