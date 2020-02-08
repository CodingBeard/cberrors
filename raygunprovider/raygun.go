package raygunprovider

import (
	"fmt"
	"github.com/MindscapeHQ/raygun4go"
)

type errorProvider struct {
	client *raygun4go.Client
}

func New(client *raygun4go.Client) *errorProvider {
	return &errorProvider{
		client: client,
	}
}

func (p *errorProvider) Error(e error) {
	if p.client != nil {
		_ = p.client.SendError(e)
	}
}

func (p *errorProvider) Recover(e interface{}) {
	if p.client != nil {
		err, ok := e.(error)
		if ok {
			_ = p.client.SendError(err)
		} else {
			_ = p.client.CreateError(fmt.Sprint(e))
		}
	}
}

func (p *errorProvider) Defer() {}