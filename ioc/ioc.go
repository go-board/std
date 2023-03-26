package ioc

import (
	"fmt"

	"github.com/go-board/std/result"
)

type ctor[T any] struct{ provide Provider[T] }

func newCtor[T any](provider Provider[T]) ctor[T] { return ctor[T]{provide: provider} }

// Provider is a function that create a value of type T
type Provider[T any] func(*Factory) (T, error)

// Factory is a container for providers
type Factory struct{ serviceCollection map[string]any }

func typename[T any]() string { return fmt.Sprintf("%T", *new(T)) }

// AddSingleton add singleton provider to Factory
func AddSingleton[T any](f *Factory, val T) { AddNamedSingleton(f, typename[T](), val) }

// AddNamedSingleton add named singleton provider to Factory
func AddNamedSingleton[T any](f *Factory, name string, val T) {
	f.serviceCollection[name] = newCtor(func(*Factory) (T, error) { return val, nil })
}

// AddPrototype add prototype provider to Factory
func AddPrototype[T any](f *Factory, fn Provider[T]) { AddNamedPrototype(f, typename[T](), fn) }

// AddNamedPrototype add named prototype provider to Factory
func AddNamedPrototype[T any](f *Factory, name string, fn Provider[T]) {
	f.serviceCollection[name] = newCtor(fn)
}

// Get retrieve typed service from Factory
func Get[T any](f *Factory) result.Result[T] { return GetNamed[T](f, typename[T]()) }

// MustGet retrieve typed service from Factory, or panic if error
func MustGet[T any](f *Factory) T { return MustGetNamed[T](f, typename[T]()) }

// GetNamed retrieve named typed service from Factory
func GetNamed[T any](f *Factory, name string) result.Result[T] {
	ctorInst, ok := f.serviceCollection[name]
	if !ok {
		return result.Errorf[T](fmt.Sprintf("service %s not registered", name))
	}
	realCtor, ok := ctorInst.(ctor[T])
	if !ok {
		return result.Errorf[T](fmt.Sprintf("%T is not a valid ctor[%s]", ctorInst, typename[T]()))
	}
	return result.FromPair(realCtor.provide(f))
}

// MustGetNamed retrieve named typed service from Factory, or panic if error
func MustGetNamed[T any](f *Factory, name string) T {
	service := GetNamed[T](f, name)
	if service.IsErr() {
		panic(service.Error())
	}
	return service.Value()
}

// Configure create a new Factory, and run specific configure function
func Configure(configure func(*Factory)) *Factory {
	factory := &Factory{serviceCollection: map[string]any{}}
	configure(factory)
	return factory
}
