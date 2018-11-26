package swarm

import "net"

type Packet struct {
	addr *net.UDPAddr
	data []byte
}

func (p *Packet) IP() string {
	return string(p.addr.IP)
}

func (p *Packet) Port() int {
	return p.addr.Port
}

func (p *Packet) Data() []byte {
	return p.data
}
