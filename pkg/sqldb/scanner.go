package sqldb

type Scanner interface {
	Scan(dest ...any) error
}
