package primitives

import "strings"

type IpAddress string

func (ip IpAddress) Network() string {
	parts := strings.Split(string(ip), ".")

	return strings.Join(parts[:3], ".")
}
