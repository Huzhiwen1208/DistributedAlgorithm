package hosts

import (
	"distributed_election/utils"
	"testing"
)

func TestDistributedElection_Bully_happy_case1(t *testing.T) {
	// arrange
	Hosts.ShutDown(4)
	// action
	Hosts.GetHostById(1).Election()
	// assert
	utils.GetAssertor(t).EqualValues(Hosts[0].MasterId, 3)
	utils.GetAssertor(t).EqualValues(Hosts[1].MasterId, 3)
	utils.GetAssertor(t).EqualValues(Hosts[2].MasterId, 3)
	utils.GetAssertor(t).EqualValues(Hosts[3].MasterId, 4)
}
