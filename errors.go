package cberrors

type ErrorsContainer struct {
	providers      []ErrorProvider
	suppressErrors bool
}

func NewErrorContainer(providers ...ErrorProvider) *ErrorsContainer {
	return &ErrorsContainer{
		providers: providers,
	}
}

type ErrorProvider interface {
	Error(e error)
	Recover(e interface{})
	Defer()
}

func (c *ErrorsContainer) AddProvider(provider ErrorProvider) {
	c.providers = append(c.providers, provider)
}

func (c *ErrorsContainer) MainDefer() {
	c.Recover()
	for _, provider := range c.providers {
		go func() {
			provider.Defer()
		}()
	}
}

func (c *ErrorsContainer) Suppress() {
	c.suppressErrors = true
}

func (c *ErrorsContainer) Unsupress() {
	c.suppressErrors = false
}

func (c *ErrorsContainer) Error(e error) {
	if !c.suppressErrors {
		for _, provider := range c.providers {
			func() {
				defer func() {
					recover()
				}()
				provider.Error(e)
			}()
		}
	}
}

func (c *ErrorsContainer) Recover() {
	e := recover()

	if e != nil {
		for _, provider := range c.providers {
			func() {
				defer func() {
					recover()
				}()
				provider.Recover(e)
			}()
		}
	}
}