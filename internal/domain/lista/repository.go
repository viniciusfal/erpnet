package lista

import "context"

//go:generate mockery --name=ListaRepository --with-expecter=true
type ListaRepository interface {
	Save(ctx context.Context, lista *Lista) error
	Update(ctx context.Context, lista *Lista) error
	FindById(ctx context.Context, idLista string) (*Lista, error)
	FindAll(ctx context.Context) ([]Lista, error)
	Delete(ctx context.Context, idLista string) error
}

type ItemRepository interface {
	Save(ctx context.Context, item *Item) error
	Update(ctx context.Context, item *Item) error
	FindById(ctx context.Context, idItem string) (*Item, error)
	FindAll(ctx context.Context) ([]Item, error)
	Delete(ctx context.Context, idItem string) error
}

type CategoriaRepository interface {
	Save(ctx context.Context, categoria *Categoria) error
	Update(ctx context.Context, categoria *Categoria) error
	FindById(ctx context.Context, idCategoria int64) (*Categoria, error)
	FindAll(ctx context.Context) ([]Categoria, error)
	Delete(ctx context.Context, idCategoria int64) error
}
