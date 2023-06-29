package api

import (
	"crypto/tls"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	portainerclient "github.com/portainer/client-api-go/v2/client"
)

type PortainerApiClientModel struct {
	Hostname      string
	ApiKey        string
	Username      string
	Password      string
	InsecureTls   bool
	Client        *portainerclient.PortainerClientAPI
	HttpTransport *httptransport.Runtime
}

func NewPortainerApiClientModel(hostname string, username string, password string, insecureTls bool) (*PortainerApiClientModel, error) {
	httpTransport := httptransport.NewWithClient(hostname, "/api", []string{"https"}, &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureTls}}})
	c := portainerclient.New(httpTransport, strfmt.Default)

	return &PortainerApiClientModel{
		Hostname:      hostname,
		Username:      username,
		Password:      password,
		Client:        c,
		HttpTransport: httpTransport,
	}, nil
}
