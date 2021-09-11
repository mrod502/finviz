package home

type Home struct {
	Signals SignalTable
}

func ParseHome(b []byte) (h *Home, err error) {
	h = new(Home)
	return
}
