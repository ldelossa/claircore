package migrations

const (
	// this migration truncates the manifest_index
	// and scanned_manifest tables and adds
	// a unique index to manifest_index.
	// this is required since the manifest_index
	// table currently bloats with duplicate
	// records.
	//
	// after this migration is complete manifests
	// will need to be re-indexed for notifications
	// on these manifests to work correctly.
	//
	// index reports will still be served
	// without a re-index being necessary.
	migration3 = `
LOCK manifest_index;
LOCK scanned_manifest;
TRUNCATE manifest_index;
TRUNCATE scanned_manifest;
CREATE UNIQUE INDEX manifest_index_unique ON manifest_index (package_id, COALESCE(dist_id, 0), COALESCE(repo_id, 0), manifest_id);
`
)
