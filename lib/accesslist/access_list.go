package accesslist

import (
	"errors"
	"github.com/aligang/YandexPracticumGoAdvanced/lib/logging"
	"net"
	"net/http"
)

func ValidateSourceIP(allowedSubnet *net.IPNet, r *http.Request) error {
	if ipAddrString := r.Header.Get("X-Real-IP"); ipAddrString != "" {
		ip := net.ParseIP(ipAddrString)
		if allowedSubnet.Contains(ip) {
			return nil
		}
		return errors.New("failed to validate IP address")
	}
	return errors.New("request Header contains no X-Real-IP field")
}

func IPValidationInterceptor(allowedSubnet string) func(http.Handler) http.Handler {
	_, subnet, err := net.ParseCIDR(allowedSubnet)
	if err != nil {
		panic(err)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := ValidateSourceIP(subnet, r)
			if err != nil {
				logging.Warn(err.Error())
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
