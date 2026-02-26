package network

func WritePacket(tun Tunnel, packet []byte) error {
	_, err := tun.Write(packet)
	return err
}
