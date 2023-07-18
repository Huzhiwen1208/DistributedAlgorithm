package hosts

type HostsRegister []*Host

var Hosts HostsRegister = []*Host{
	{
		Id: 1,
		IdMap: map[int32]string{
			2: "192.168.0.2",
			3: "192.168.0.3",
			4: "192.168.0.4",
		},
		KeepAlive: true,
		MasterId:  4,
	},
	{
		Id: 2,
		IdMap: map[int32]string{
			1: "192.168.0.1",
			3: "192.168.0.3",
			4: "192.168.0.4",
		},
		KeepAlive: true,
		MasterId:  4,
	},
	{
		Id: 3,
		IdMap: map[int32]string{
			1: "192.168.0.1",
			2: "192.168.0.2",
			4: "192.168.0.4",
		},
		KeepAlive: true,
		MasterId:  4,
	},
	{
		Id: 4,
		IdMap: map[int32]string{
			1: "192.168.0.1",
			2: "192.168.0.2",
			3: "192.168.0.3",
		},
		KeepAlive: true,
		MasterId:  4,
	},
}

func (p *HostsRegister) GetHostById(id int32) *Host {
	return Hosts[id-1]
}

func (p *HostsRegister) ShutDown(id int32) {
	Hosts[id-1].KeepAlive = false
}
