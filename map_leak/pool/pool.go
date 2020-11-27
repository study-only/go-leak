package pool

import (
	"errors"
	"sync"
)

// 连接池
type Pool struct {
	Name  string
	peers *sync.Map
}

// 新建
func NewPool(name string) *Pool {
	return &Pool{
		Name:  name,
		peers: new(sync.Map),
	}
}

// 添加连接节点
func (p *Pool) Add(peer *Peer) error {
	_, exist := p.peers.LoadOrStore(peer.WS, peer)
	if exist {
		return errors.New("already exist")
	}

	return nil
}

// 移除连接节点
func (p *Pool) Remove(peer *Peer) error {
	p.peers.Delete(peer.WS)
	return nil
}

// 当前节点数量
func (p *Pool) GetPeerCount() int64 {
	var count int64 = 0
	p.Range(func(*Peer) bool {
		count++
		return true
	})

	return count
}

// 遍历节点
func (p *Pool) Range(f func(peer *Peer) bool) {
	p.peers.Range(func(k, v interface{}) bool {
		peer := v.(*Peer)
		return f(peer)
	})
}
