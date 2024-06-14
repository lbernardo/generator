package repository

import (
    "context"
	"go.mongodb.org/mongo-driver/mongo"
	"{{.pkg}}/internal/app/{{.name}}/entity"
)

type I{{.className}}Repository interface {
    Create({{.name}} *entity.{{.className}}) error
}

type {{.className}}Repository struct {
	collection *mongo.Collection
}

func New{{.className}}Repository(db *mongo.Database) I{{.className}}Repository {
	return &{{.className}}Repository{
		collection: db.Collection("{{.name}}s"),
	}
}


func (r *{{.className}}Repository) Create({{.name}} *entity.{{.className}}) error {
    _, err := r.collection.InsertOne(context.Background(), {{.name}})
    return err
}
