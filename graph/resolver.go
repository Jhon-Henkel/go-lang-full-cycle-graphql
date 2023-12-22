package graph

import "github.com/Jhon-Henkel/go-lang-full-cycle-graphql/internal/database"

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Course
}
