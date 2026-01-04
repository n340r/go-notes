package domain

type Category struct {
	id   int
	name string
}

type NewCategoryData struct {
	ID   int
	Name string
}

func (b Category) ID() int {
	return b.id
}

func (b Category) Name() string {
	return b.name
}

func newCategory(data NewCategoryData) (Category, error) {
	return Category{
		id:   data.ID,
		name: data.Name,
	}, nil
}
