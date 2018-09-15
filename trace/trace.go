package trace

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"time"
)

// TraceTime struct
type TraceTime struct {
	dnsStartTime time.Time
	dnsEndTime   time.Time

	tlsHandshakeStartTime time.Time
	tlsHandshakeEndTime   time.Time

	connectionStartTime time.Time
	connectionEndTime   time.Time

	ttfbEndTime time.Time

	startTime time.Time
	endTime   time.Time
}

func (t *TraceTime) GetDnsTime() time.Duration {
	return t.dnsEndTime.Sub(t.dnsStartTime)
}

func (t *TraceTime) GetTLSHandshakeTime() time.Duration {
	return t.tlsHandshakeEndTime.Sub(t.tlsHandshakeStartTime)
}

func (t *TraceTime) GetConnectionTime() time.Duration {
	return t.connectionEndTime.Sub(t.connectionStartTime)
}

func (t *TraceTime) GetTTFBTime() time.Duration {
	return t.ttfbEndTime.Sub(t.startTime)
}

func (t *TraceTime) GetTotalTime() time.Duration {
	return t.endTime.Sub(t.startTime)
}

// Trace
func Trace(url string) (*TraceTime, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	traceTime := TraceTime{}

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) {
			traceTime.dnsStartTime = time.Now()
		},
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			traceTime.dnsEndTime = time.Now()
		},

		TLSHandshakeStart: func() {
			traceTime.tlsHandshakeStartTime = time.Now()
		},
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			traceTime.tlsHandshakeEndTime = time.Now()
		},

		ConnectStart: func(network, addr string) {
			traceTime.connectionStartTime = time.Now()
		},
		ConnectDone: func(network, addr string, err error) {
			traceTime.connectionEndTime = time.Now()
		},

		GotFirstResponseByte: func() {
			traceTime.ttfbEndTime = time.Now()
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	traceTime.startTime = time.Now()
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		return nil, err
	}
	traceTime.endTime = time.Now()

	return &traceTime, nil
}
