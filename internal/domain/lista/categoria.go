package lista

type Categoria struct {
	ID   int64
	Nome string
}

type InputCategoria struct {
	Nome string
}

func CreateCategoria(input InputCategoria) (*Categoria, error) {

	if input.Nome == "" {
		return nil, ERRRNOMECATEGORIAOBRIGATORIO
	}

	c := Categoria{
		Nome: input.Nome,
	}

	return &c, nil
}
