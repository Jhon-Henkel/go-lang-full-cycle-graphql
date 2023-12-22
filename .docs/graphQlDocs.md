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
- Find All
  ```graphql
  query queryCategories {
    categories {
      id
      name
      description
    }
  }
  ```
  
- Find All with Courses
  ```graphql
  query queryCategoryWithCourses {
    categories {
      id
      name
      description
      courses {
        id
        name
      }
    }
  }
  ```
  
## Courses
- New Course:
    ```graphql
    mutation createCourse {
      createCourse(
        input: {
          name: "Go Expert"
          description: "Torne-se expert em Go-lang"
          categoryId: "bfa339a2-af2e-4af9-a7a0-16c57b1e82ea"
        }
      )
      {
        id
        name
        description
      }
    }
    ```
  
- Find All
    ```graphql
    query queryCouses {
      courses {
        id
        name
        description
      }
    }
    ```
  
- Find All with Category
    ```graphql
        query queryCoursesWithCategory {
      courses {
        id
        name
        description
        category {
          id
          name
          description
        }
      }
    }
    ```