// Package migrations
package migrations

import "embed"

//go:embed *.sql
var EmbedMigrations embed.FS
