package pool

type WS struct {
	WID string
}

type Peer struct {
	UID string
	WS  *WS
}

func NewPeer(id string) *Peer {
	return &Peer{
		UID: "u:id",
		WS: &WS{
			WID: "ws:" + id,
		},
	}
}
