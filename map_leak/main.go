package main

import (
	"github.com/study-only/go-leak/map_leak/pool"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"
)

var p = pool.NewPool("ws")

func main() {
	go enter()
	go leave()
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func enter() {
	for i := 0; i < 1000000; i++ {
		peer := newPeer(i)
		if err := p.Add(peer); err != nil {
			log.Printf("ERROR add peer %d", i)
		}
		log.Printf("peer %d added", i)
	}

	log.Printf("peers all added")
	for {
		count := 0
		p.Range(func(peer *pool.Peer) bool {
			log.Printf("uid=%s, wid=%s", peer.UID, peer.WS.WID)
			count++
			return true
		})
		log.Printf("has %d peers", count)
		time.Sleep(10 * time.Second)
	}
}

func leave() {
	for {
		count := 0
		p.Range(func(peer *pool.Peer) bool {
			if err := p.Remove(peer); err != nil {
				log.Printf("ERROR remove peer %s", peer.UID)
			}
			count++
			return true
		})
		log.Printf("remove %d peers", count)
		time.Sleep(time.Second)
	}
}

func newPeer(id int) *pool.Peer {
	uid := strconv.Itoa(id)
	return pool.NewPeer(uid)
}
