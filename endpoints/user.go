package endpoints

type UserURLBuilder struct{}

func (r UserURLBuilder) Build(endpoint string) (string, error) {
	return UserURL(endpoint)
}
