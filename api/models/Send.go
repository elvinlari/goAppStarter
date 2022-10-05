package models

import (
	"errors"
	"html"
	"net"
	"strings"

	// "net"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Send struct {
	User    string `json:"user"`
	From    string `json:"source"`
	To      string `json:"dest"`
	Message string `json:"message"`
	MssId   string `json:"msgID"`
	MetaDt  string `json:"METADATA"`
}

func (p *Send) Prepare() {
	p.User = html.EscapeString(strings.TrimSpace(p.User))
	p.From = html.EscapeString(strings.TrimSpace(p.From))
	p.To = html.EscapeString(strings.ReplaceAll(p.To, "+", ""))
	p.To = html.EscapeString(strings.ReplaceAll(p.To, " ", ""))
	// p.Message = html.EscapeString(strings.TrimSpace(p.Message))
	if p.MssId == "" {
		p.MssId = genUUID()
	}
}

func genUUID() string {
	id := uuid.New()
	return string(id.String())
}

// FromRequest extracts the user IP address from req, if present.
func FromRequest(r *http.Request) (string, error) {
	// ip,_, err := net.SplitHostPort(r.RemoteAddr)
	// if err != nil {
	//   return "", fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	// }
	// if ip == "" {
	//   return "", fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	// }
	// fmt.Println("IP IS: ", ip)
	// return ip, nil

	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", fmt.Errorf("token not found")
}

func RequestIp(r *http.Request) (string, error) {
	// ip, _, err := net.SplitHostPort(r.RemoteAddr)
	// if err != nil {
	// 	return "", fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	// }
	// if ip == "" {
	// 	return "", fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	// }
	// fmt.Println("IP IS: ", ip)
	// forward := r.Header.Get("X-Forwarded-For")
	// fmt.Println("IP forwarded: ", forward)
	// return ip, nil

	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	fmt.Println("Real IP IS: ", ip)
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	fmt.Println("Forwarded IPs ARE: ", ips)
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	fmt.Println("RemteAddr IP IS: ", ip)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")

}

func (p *Send) Validate() error {
	if p.User == "" {
		return errors.New("Username is required")
	}
	if p.From == "" {
		return errors.New("Source address is required")
	}
	if p.To == "" {
		return errors.New("destination address is required")
	}
	if p.Message == "" {
		return errors.New("Message is required")
	}
	return nil

}
