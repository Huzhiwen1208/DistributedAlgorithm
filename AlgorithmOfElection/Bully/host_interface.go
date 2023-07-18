package hosts

type HostInterface interface {
	Election() (int32, error)
	SendVictory() error
}
