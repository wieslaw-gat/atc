package migrations

import "github.com/concourse/atc/dbng/migration"

func AddPausedToJobs(tx migration.LimitedTx) error {
	_, err := tx.Exec(`ALTER TABLE jobs ADD COLUMN paused boolean DEFAULT(false)`)

	return err
}
