package socket

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

// HTTPServer http server over socket
type HTTPServer struct {
	sock string
	ln   net.Listener
}

// NewHTTPServer create server
func NewHTTPServer(addr string) (*HTTPServer, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	host := u.Host
	if u.Scheme == "unix" {
		host = u.Path
	}
	ln, err := net.Listen(u.Scheme, host)
	if err != nil {
		return nil, err
	}
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("来者何人"))
	})
	go http.Serve(ln, nil)
	return &HTTPServer{addr, ln}, nil
}

// Close server
func (s *HTTPServer) Close() {
	s.ln.Close()
}

// HTTPClient ....
type HTTPClient struct {
	addr   string
	client *http.Client
}

// NewHTTPClient ...
func NewHTTPClient(addr string) (*HTTPClient, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	tr := new(http.Transport)
	dialer := &net.Dialer{
		Timeout: time.Second,
	}
	tr.DialContext = func(ctx context.Context, _, _ string) (net.Conn, error) {
		return dialer.DialContext(ctx, u.Scheme, u.Path)
	}
	client := &http.Client{
		Transport: tr,
	}
	return &HTTPClient{u.Path, client}, nil
}

// Ping request
func (h *HTTPClient) Ping() ([]byte, error) {
	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		return nil, err
	}
	req.Host = "fungo"
	req.URL.Host = h.addr
	req.URL.Scheme = "http"
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
