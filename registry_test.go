//go:build go1.27

package parco

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

type (
	registryAnimal struct {
		Specie string
		Age    uint8
	}

	registryFlat struct {
		Address string
		Rooms   uint8
	}

	testRegistry struct {
		registry *Registry[uint8]
		animal   *Model[registryAnimal]
		flat     *Model[registryFlat]
	}
)

const (
	registryAnimalID uint8 = 1
	registryFlatID   uint8 = 2
)

func newTestRegistry() testRegistry {
	animal := NewModel[registryAnimal](ObjectFactory[registryAnimal]()).
		Field(SmallVarchar(),
			func(a *registryAnimal) string { return a.Specie },
			func(a *registryAnimal, v string) { a.Specie = v }).
		Field(UInt8(),
			func(a *registryAnimal) uint8 { return a.Age },
			func(a *registryAnimal, v uint8) { a.Age = v })
	flat := NewModel[registryFlat](ObjectFactory[registryFlat]()).
		Field(SmallVarchar(),
			func(f *registryFlat) string { return f.Address },
			func(f *registryFlat, v string) { f.Address = v }).
		Field(UInt8(),
			func(f *registryFlat) uint8 { return f.Rooms },
			func(f *registryFlat, v uint8) { f.Rooms = v })
	registry := NewRegistry[uint8](UInt8()).
		Register(registryAnimalID, animal).
		Register(registryFlatID, flat)
	return testRegistry{registry: registry, animal: animal, flat: flat}
}

func TestRegistry(t *testing.T) {
	t.Run("roundtrip dispatches on the id", func(t *testing.T) {
		fx := newTestRegistry()
		var buf bytes.Buffer
		require.NoError(t, fx.registry.Compile(Entry[uint8]{ID: registryAnimalID, Value: registryAnimal{Specie: "cat", Age: 3}}, &buf))
		require.NoError(t, fx.registry.Compile(Entry[uint8]{ID: registryFlatID, Value: registryFlat{Address: "main st", Rooms: 2}}, &buf))

		first, err := fx.registry.Parse(&buf)
		require.NoError(t, err)
		second, err := fx.registry.Parse(&buf)
		require.NoError(t, err)

		require.Equal(t, Entry[uint8]{ID: registryAnimalID, Value: registryAnimal{Specie: "cat", Age: 3}}, first)
		require.Equal(t, Entry[uint8]{ID: registryFlatID, Value: registryFlat{Address: "main st", Rooms: 2}}, second)
	})

	t.Run("compile with an unregistered id writes nothing", func(t *testing.T) {
		fx := newTestRegistry()
		var buf bytes.Buffer

		err := fx.registry.Compile(Entry[uint8]{ID: 9, Value: registryAnimal{}}, &buf)

		require.ErrorIs(t, err, ErrUnknownType)
		require.Zero(t, buf.Len())
	})

	t.Run("compile with a value of the wrong type", func(t *testing.T) {
		fx := newTestRegistry()
		var buf bytes.Buffer

		err := fx.registry.Compile(Entry[uint8]{ID: registryAnimalID, Value: registryFlat{}}, &buf)

		require.ErrorIs(t, err, ErrUnknownType)
	})

	t.Run("parse with an unregistered id", func(t *testing.T) {
		fx := newTestRegistry()

		_, err := fx.registry.Parse(bytes.NewReader([]byte{9}))

		require.ErrorIs(t, err, ErrUnknownType)
	})

	t.Run("registry nests as a field", func(t *testing.T) {
		fx := newTestRegistry()
		type envelope struct {
			Seq     uint8
			Payload Entry[uint8]
		}
		model := NewModel[envelope](ObjectFactory[envelope]()).
			Field(UInt8(),
				func(e *envelope) uint8 { return e.Seq },
				func(e *envelope, v uint8) { e.Seq = v }).
			Field(fx.registry,
				func(e *envelope) Entry[uint8] { return e.Payload },
				func(e *envelope, v Entry[uint8]) { e.Payload = v })
		value := envelope{Seq: 7, Payload: Entry[uint8]{ID: registryFlatID, Value: registryFlat{Address: "x", Rooms: 1}}}
		var buf bytes.Buffer
		require.NoError(t, model.Compile(value, &buf))

		parsed, err := model.Parse(&buf)

		require.NoError(t, err)
		require.Equal(t, value, parsed)
	})

	t.Run("registering the same id twice panics", func(t *testing.T) {
		fx := newTestRegistry()

		require.Panics(t, func() { fx.registry.Register(registryAnimalID, fx.animal) })
	})
}
