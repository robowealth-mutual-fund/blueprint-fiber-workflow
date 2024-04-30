package container

import "go.uber.org/dig"

func (c *Container) Configure() error {
	for _, ctor := range constructors {
		if err := c.container.Provide(ctor.function, ctor.options...); err != nil {
			return err
		}
	}

	return nil
}

type constructor struct {
	function interface{}
	options  []dig.ProvideOption
}

func newConstructor(f interface{}, opts ...dig.ProvideOption) constructor {
	options := make([]dig.ProvideOption, len(opts))

	copy(options, opts)

	return constructor{
		function: f,
		options:  options,
	}
}
