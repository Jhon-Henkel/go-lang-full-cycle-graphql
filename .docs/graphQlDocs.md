## Category
- New Category:
    ```graphql
    mutation createCategory {
      createCategory(
        input: {
          name: "tecnologia"
          description: "cursos de tecnologia"
        }
      )
      {
        id
        name
        description
      }
    }
    ```