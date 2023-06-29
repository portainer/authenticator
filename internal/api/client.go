package api

import (
	"crypto/tls"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/portainer/authenticator/internal/types"
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

func NewPortainerApiClientModel(options *types.Options) (*PortainerApiClientModel, error) {
	httpTransport := httptransport.NewWithClient(*options.PortainerURL, "/api", []string{"https"}, &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: *options.InsecureTls}}})
	c := portainerclient.New(httpTransport, strfmt.Default)

	return &PortainerApiClientModel{
		Hostname:      *options.PortainerURL,
		Username:      "",
		Password:      "",
		Client:        c,
		HttpTransport: httpTransport,
	}, nil
}
