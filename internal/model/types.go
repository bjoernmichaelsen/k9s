package model

import (
	"context"

	"github.com/derailed/k9s/internal/client"
	"github.com/derailed/k9s/internal/dao"
	"github.com/derailed/k9s/internal/render"
	"github.com/derailed/tview"
	"k8s.io/apimachinery/pkg/runtime"
)

// Igniter represents a runnable view.
type Igniter interface {
	// Start starts a component.
	Init(ctx context.Context) error

	// Start starts a component.
	Start()

	// Stop terminates a component.
	Stop()
}

// Hinter represent a menu mnemonic provider.
type Hinter interface {
	// Hints returns a collection of menu hints.
	Hints() MenuHints
}

// Primitive represents a UI primitive.
type Primitive interface {
	tview.Primitive

	// Name returns the view name.
	Name() string
}

// Component represents a ui component
type Component interface {
	Primitive
	Igniter
	Hinter
}

// Renderer represents a resource renderer.
type Renderer interface {
	// Render converts raw resources to tabular data.
	Render(o interface{}, ns string, row *render.Row) error

	// Header returns the resource header.
	Header(ns string) render.HeaderRow

	// ColorerFunc returns a row colorer function.
	ColorerFunc() render.ColorerFunc
}

// Cruder performs crud operations.
type Cruder interface {
	// List returns a collection of resources.
	List(ctx context.Context, ns string) ([]runtime.Object, error)

	// Get returns a resource instance.
	Get(ctx context.Context, path string) (runtime.Object, error)
}

// Lister represents a resource lister.
type Lister interface {
	Cruder

	// Init initializes a resource.
	Init(ns, gvr string, f dao.Factory)
}

// Describer represents a resource describer.
type Describer interface {
	// ToYAML return resource yaml.
	ToYAML(ctx context.Context, path string) (string, error)

	// Describe returns a resource description.
	Describe(client client.Connection, gvr, path string) (string, error)
}

// ResourceMeta represents model info about a resource.
type ResourceMeta struct {
	DAO      dao.Accessor
	Renderer Renderer
}
