package app

type ctxKey int

const (
	ctxAccount ctxKey = iota
	ctxProfile
)

// API provides application resources and handlers
type API struct {
	Profile *ProfileResource
}
