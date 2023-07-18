package hosts

import (
	"log"
	"math/rand"
	"time"
)

type Host struct {
	Id        int32
	IdMap     map[int32]string // id -> ip
	KeepAlive bool
	MasterId  int32
}

// 选举操作
func (p *Host) Election() error {
	// 参加选举的节点
	log.Printf("Election- the node: %v", p.Id)
	// 1. 获取比自己大的结点
	greaterIds := p.GetGreaterIds()
	// 如果自己是最大,则直接宣誓主权
	if len(greaterIds) == 0 {
		return p.SendVictory()
	}

	// 否则, 对每个比自己大的节点都发送Election消息
	var rejected = false
	for id, ip := range p.IdMap {
		log.Printf("Election- connecting ip: %v", ip)
		host := Hosts.GetHostById(id)
		if !host.KeepAlive {
			log.Printf("Election- host: %v is not keep alive", host.Id)
			continue
		}

		// 发送消息
		t0 := time.Now()
		agree := host.AcceptElection(p.Id)
		cost := time.Since(t0).Seconds()
		if cost > 6 { // 超过6秒则认为同意
			log.Printf("host:{%v} connected exceed 6s", ip)
			agree = true
		}

		if !agree {
			// 对首个不同意的节点开始选举
			rejected = true
			host.Election()
			break
		}
	}

	if !rejected {
		return p.SendVictory()
	}

	return nil
}

func (p *Host) GetGreaterIds() []int32 {
	var result []int32
	for id := range p.IdMap {
		if id > p.Id {
			result = append(result, id)
		}
	}
	return result
}

// 接收到Election消息后,是否同意fromId做主节点
func (p *Host) AcceptElection(fromId int32) bool {
	if fromId > p.Id {
		return true
	}

	// 如果发来选举消息的节点比自己的节点小, 则发送一个Alive消息回去(false, 不同意)
	if fromId < p.Id {
		var n = rand.Int31n(6)
		time.Sleep(time.Duration(n) * time.Second) // 随机停止0-5秒

		return false
	}

	return false
}

// func (p *Host) AcceptAlive(fromId int32) error {

// 	return nil
// }

func (p *Host) AcceptVictory(fromId int32) error {
	p.MasterId = fromId
	log.Printf("Host %d has accepted the victory id: %v", p.Id, fromId)
	return nil
}

func (p *Host) SendVictory() error {
	for id, ip := range p.IdMap {
		log.Printf("SendVictory- connecting ip: %v", ip)
		host := Hosts.GetHostById(id)
		if !host.KeepAlive {
			log.Printf("SendVictory- host %v is not alive", host.Id)
			continue
		}
		// 发送消息
		host.AcceptVictory(p.Id)
	}

	p.MasterId = p.Id

	return nil
}
