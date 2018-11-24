package health

// Possible states of a health check
const (
	HEALTHY   = "healthy"
	UNHEALTHY = "unhealthy"
)

// Checks describes a list of interfaces that can be pinged
type Checks map[string]interface{ Ping() bool }

// Create will turn a list of health checks in to a map of check->health
func Create(healthChecks Checks) map[string]string {

	r := map[string]string{}
	for name, check := range healthChecks {
		if r[name] = HEALTHY; !check.Ping() {
			r[name] = UNHEALTHY
		}
	}
	return r
}

// CreateSummarized is will turn a list of health checks into healthy/unhealthy
func CreateSummarized(healthChecks Checks) string {
	return Summarize(Create(healthChecks))
}

// Summarize is will turn a map of name->health into healthy/unhealthy
func Summarize(status map[string]string) string {
	for _, status := range status {
		if status == UNHEALTHY {
			return UNHEALTHY
		}
	}
	return HEALTHY
}
