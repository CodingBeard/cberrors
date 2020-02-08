package sentryprovider

import (
	"github.com/getsentry/sentry-go"
	"time"
)

type errorProvider struct{}

func New() *errorProvider {
	return &errorProvider{}
}

func (p *errorProvider) Error(e error) {
	sentry.CaptureException(e)
}

func (p *errorProvider) Recover(e interface{}) {
	sentry.CurrentHub().Recover(e)
}

func (p *errorProvider) Defer() {
	sentry.Flush(time.Second * 5)
}


