package joke

type Repository interface {
	Random() (Model, error)
	Query(query string) (ModelQuery, error)
}
