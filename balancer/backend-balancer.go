package balancer

//BackendBalancer balances the request among loaded backends
type BackendBalancer interface {
	// Next returns backend for the request
	Next() (interface{}, error)

	// Load initialize and load balancer backends. This method will reset everything.
	// Use this method to initialize balancer
	Load([]interface{}) error

	// Reload preserves existing elements only removes/add delta to existing backends.
	// Use this method to periodic refresh of backends
	Reload([]interface{}) error
}
