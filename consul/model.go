package consul

const (
	serviceStatusPass = "passing"
	consulStatusPath  = "/v1/status/leader"
)

type Service struct {
	Name    string
	Address string
	Port    int
}
