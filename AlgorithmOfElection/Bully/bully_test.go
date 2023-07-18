package bully

import (
	hosts "distributed_election/AlgorithmOfElection/Hosts"
	"distributed_election/utils"
	"testing"
)

func TestDistributedElection_Bully_happy_case1(t *testing.T) {
	// arrange
	hosts.Hosts.ShutDown(4)
	// action
	hosts.Hosts.GetHostById(1).Election()
	// assert
	utils.GetAssertor(t).EqualValues(hosts.Hosts[1].MasterId, 3)
}
