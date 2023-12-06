# Graphql with Gin Framework using Golang and Postgresql database using GORM

### Contents
   - [Contents](#contents)
   - [Introduction](#introduction)
   - [How Graphql queries called using golang and gqlgen tool?](#how-graphql-queries-called-using-golang-and-gqlgen-tool) 
   - [Documentation](#documentation) 
   - [Getting Started with gqlgen](#getting-started-with-gqlgen) 
   - [Finally Running the application](#finally-running-the-application)
   - [Conclusion](#conclusion)
### Introduction

[GraphQL](https://graphql.org/) is a runtime query language created by Facebook as an alternative to using REST API’s. While [GO language](https://go.dev/) is popular for its support for concurrency and scaling as core needed. 

In this blog, we will be exploring writing API’s with [GraphQL](https://graphql.org/) and [GO language](https://go.dev/). In this session we will be exploring how to initialize simple [GraphQL](https://graphql.org/) application using [Go](https://go.dev/) with [gqlgen](https://gqlgen.com/) library.  Also, we will be exploring how to use custom resolvers and custom structuring of application. So, let’s get started! 

### How Graphql queries called using golang and gqlgen tool? 

As we all know, in REST API’s we can call different URLs for different requests, but using [gqlgen tool](https://gqlgen.com/) we make all request to the single URL using POST request only. Yes, you heard it right only single URL for all different requests. Here, you might be thinking then how it differentiates which method to call?! No worries, while sending request we need to send a request body in string format expressing which query or mutation with its parameter and what response is needed.  

Here comes the next question, how would anyone know what the parameters are needed for request and which parameters we can get as a response? 

### Documentation 

The [gqlgen](https://gqlgen.com/) tool has a feature i.e., it generates documentation for us using [GraphQL](https://graphql.org/) or GraphQL Playground. Isn’t it great we don’t need things like swagger for documentation which we create while building REST API’s [gqlgen](https://gqlgen.com/) gives flexibility for self-documenting, whatever schemas are defined.

### Getting Started with gqlgen 

Before starting with how to use gqlgen tool for creating graphql based API application, I will highly recommend you read [GraphQL Documentation](https://graphql.org/learn/) if you don’t know the basics of graphql or datatypes graphql uses. Also, if you are new to GO language please refer to [Go Documentation](https://go.dev/learn/). 

So, let's start with brief overview of [graphql](https://graphql.org/),  the most basic component of [GraphQL](https://graphql.org/) schemas is object, which represents a kind of object you might want to fetch from your API or to be precise your service. Below is the representation of an object: 
```graphql
type ObjectName {
    # Int is the datatype of the variabe you can put '!' mark
    # to make vaiable as required field.
    Variable: Int
}
```
Next thing you want to know is about to special types, basically heart of [graphql](https://graphql.org/) without which you might not even run any API or service. Every [GraphQL](https://graphql.org/) service has query or mutatuion type. You might be thinking what so special about this type? So, these two types are used to call services which makes them heart of [GraphQL](https://graphql.org/). Query type is used for viewing data, if we compare to REST API it is same as GET query. 

At this point you might get a question, where mutation type is used as query could easily do any processing of data as I am saying query type is same as GET request in REST API.  So, you might be aware of that it’s suggested by convention that one doesn’t use GET request for modification of any data similarly it goes for query. Hence, we use mutation for modification of data. 

Mutation also provides feature for inputting an object by creating custom object of input type, it is same as Object as shown above but instead of type keyword, we replace it with input. I will show you how to write query or mutation but for understanding how input are created below is the representation of input object:
```graphql
input ObjectName {
    # It is actually an object but used for inputting
    # in services called by any query or mutatuion type
    Variable: Int
}
```
Moving on to using [gqlgen tool](https://gqlgen.com/), for creating application. You can refer to the [official documentation of gqlgen tool](https://gqlgen.com/). 

Let’s get started creating our application, 

> Step 1: Initializing a new go module.
Assuming You have already created a directory for the project if not
```shell
mkdir example

cd example
```
Now create a new go module
```shell
go mod init example
```
> Step 2: Run below given line in your terminal to get gqlgen tools. 
```shell
mkdir -p tools && printf '// +build tools\npackage tools\nimport (_ "github.com/99designs/gqlgen"\n _ "github.com/99designs/gqlgen/graphql/introspection")' | gofmt > tools/tools.go

go mod tidy
```
> Step 3: Initialise gqlgen config and basic structure 
```shell
go run github.com/99designs/gqlgen init

go mod tidy
```
The above code will generate the following structure: 
```text
example
├── graph
│   ├── model
│   │   └── models_gen.go
│   ├── generated.go
│   ├── resolver.go
│   ├── schema.graphqls
│   └── schema.resolver.go
├── tools
│   └── tools.go
├── go.mod
├── go.sum
├── gqlgen.yml
└── server.go 
```
> Step 4: We will be creating our custom structure
First of all we will adjust some files and create some directories before actually writting some code.

1. Create a directory *example/models* copy  **example/graph/model/model_gen.go** file to *example/models*.  
    It will look like this:

    ```text
    example
    ├── models
    └── models_gen.go
    ```
2. Create a directory *example/schema*  copy **example/graph/schmea.graphqls** file to *example/schemas*.  
    It will look like this
    ```text
    example
    ├── schemas
    └── schema.graphqls
    ```
3. Now Remove graph directory, and do changes in **example/gqlgen.yml** file.
```yaml
schema:
  - schemas/*.graphqls

# Where should the generated server code go?
exec:
  filename: generated/generated.go
  package: generated

# Where should any generated models go?
model:
  filename: models/models_gen.go
  package: models

# Where should the resolver implementations go?
resolver:
  layout: follow-schema
  dir: generated
  package: generated
  filename_template: "{name}_resolvers.go"

# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.
autobind:
#  - "example/graph/model"

# This section declares type mapping between the GraphQL and go type systems
#
# The first line in each type will be used as defaults for resolver arguments and
# modelgen, the others will be allowed when binding to fields. Configure them to
# your liking
models:
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
```

4. Run generate command
```shell
go run github.com/99designs/gqlgen generate
```

5. The final structure will look like this after generate command:
```text
example
├── generated
│   ├── generated.go
│   ├── resolver.go
│   └── schema_resolvers.go
├── models
│   └── models_gen.go
├── schemas
│   └── schema.graphqls
├── tools
│   └── tools.go
├── go.mod
├── go.sum
├── gqlgen.yml
└── server.go 
```
`Note that we are still far away from achieving our final structure.`

6. Now create remaining directories and empty files for now except for files in generated directory as shown in below structure.
```text
example
├── database
│   ├── postgres
│   │   └── postgres.go
│   ├── repository
│   │   ├── company_repository.go
│   │   ├── employee_repository.go
│   │   ├── employee_task_repository.go
│   │   └── task_repository.go
├── dataloaders
│   ├── companyloader_gen.go    // Auto Created Do not create file
│   ├── dataloaders.go          
│   ├── employeeloader_gen.go   // Auto Created Do not create file
│   └── taskloader_gen.go       // Auto Created Do not create file
├── domain
│   ├── company_domain.go
│   ├── domain.go
│   ├── employee_domain.go
│   └── task_domain.go
├── generated
│   ├── company_resolvers.go    // Auto Created Do not create file
│   ├── employee_resolvers.go   // Auto Created Do not create file
│   ├── generated.go            // Auto Created Do not create file
│   ├── mutation_resolvers.go   // Auto Created Do not create file
│   ├── query_resolvers.go      // Auto Created Do not create file
│   ├── resolver.go             // Auto Created Do not create file
│   └── task_resolvers.go       // Auto Created Do not create file
├── models
│   ├── company_model.go
│   ├── employee_model.go
│   ├── models_gen.go
│   ├── pagination_model.go
│   └── task_model.go
├── schemas
│   ├── company.graphqls
│   ├── employee.graphqls
│   ├── mutation.graphqls
│   ├── query.graphqls
│   ├── schema.graphqls
│   └── task.graphqls
├── tools
│   └── tools.go
├── go.mod
├── go.sum
├── gqlgen.yml
└── server.go 
```
> Step 5: Assuming above structure is created. We will directly start creating models we need in models directory.

1. First with *models/company_model.go* create *Company* Model as shown below:

```go
package models

type Company struct {
    ID        int         `gorm:"column:id;AUTO_INCREMENT;primaryKey;" json:"id"`
    Name      string      `gorm:"column:name;type:varchar(255);not null;unique;" json:"name"`
    Website   *string     `gorm:"column:website;type:varchar(255);unique;" json:"website,omitempty"`
    Employees []*Employee `gorm:"foreignKey:CompanyID;references:ID;constraint:OnDelete:CASCADE;" json:"employees"`
}

// Used for Union Implementation 
func (Company) IsResultUnion() {}
```

2. Simmilarly create *Employee* and *Task* model inside *models/employee_model.go* and *models/tasks_model.go* respectively.

```go
package models

type Employee struct {
    ID          int      `gorm:"column:id;AUTO_INCREMENT;primaryKey;" json:"id"`
    Name        string   `gorm:"column:name;type:varchar(255);not null;unique;" json:"name"`
    Designation string   `gorm:"column:designation;type:varchar(255);not null;" json:"designation"`
    Email       string   `gorm:"column:email;type:varchar(255);not null;unique;" json:"email"`
    CompanyID   int      `gorm:"column:company_id;not null;" json:"company_id"`
    Company     *Company `json:"company,omitempty"`
    Tasks       []*Task  `gorm:"many2many:employee_tasks;" json:"tasks,omitempty"`
}

// Used for Union Implementation 
func (Employee) IsResultUnion() {}
```

```go
package models

import "time"

type Task struct {
    ID          int         `gorm:"column:id;AUTO_INCREMENT;primaryKey;" json:"id"`
    Name        string      `gorm:"column:name;type:varchar(255);not null;unique;" json:"name"`
    Description *string     `gorm:"column:description;type:varchar(255);" json:"description"`
    StartDate   time.Time   `gorm:"column:start_date;not null;" json:"start_date"`
    EndDate     *time.Time  `gorm:"column:end_date;" json:"end_date,omitempty"`
    Employees   []*Employee `gorm:"many2many:employee_tasks;" json:"employees"`
}

func (Task) IsResultUnion() {}
```
> Step 6: Now we will be creating types in graphql to bind above created objects

Create *Company*, *Employee* and *Task* type in *schemas/company.graphqls*, *schemas/employee.graphqls* and *schemas/task.graphqls* respectively as shown below.

```graphql
type Company{
    id:         ID!        
    name:       String!     
    website:    String   
    employees:  [Employee!]!
}
```
```graphql
type Employee {
    id:             ID!
    name:           String!
    designation:    String!
    email:          String!
    company_id:	    Int!
    company:	    Company
    tasks:          [Task!]
}
```
```graphql
type Task{
    id:             ID!
    name:           String!
    description:    String
    start_date:     Time!
    end_date:       Time
    employees:      [Employee!]!
}
```
> Step 6: Create a mutation and query type and create services we needed.

1. First we will create services for query in schemas/query.graphqls as below:
```graphql
type Query{
    employees(filter:FilterInput, page:Int = 1):Pagination
    tasks(filter:FilterInput,page:Int=1):Pagination
    companies(filter:FilterCompanyInput,page:Int=1):Pagination
}
```

2. Now we will create services for mutation in schemas/mutation.graphqls as below:
```graphql
type Mutation{
    registerCompany(input:CreateCompany!):Company
    registerEmployee(input:CreateEmployee!):Employee
    createTask(input:CreateTask):Task
    assignTask(task_id:Int!,employee_ids:[Int!]!):Task
}
``` 
3. The Final step we will be creating necessary Input Object and Response Structure needed both in query and mutatuion type by their services in schemas/schmea.graphqls
```graphql
scalar Time
scalar Int64

input FilterInput{
    employee_name:          String
    employee_designation:   String
    task_name:              String
    employee_email:         String
    task_start_date:        Time
    task_end_date:          Time  
}

input FilterCompanyInput{
    employee_name:      String
    company_name:       String
    employee_email:     String
    company_website:    String
}

input CreateCompany{
    name:       String!     
    website:    String
}

input CreateEmployee{
    name:           String!
    designation:    String!
    email:          String!
    company_id:     Int!
}

input CreateTask{
    name:           String!
    description:    String
    start_date:     Time!
    end_date:       Time
    employee_ids:   [Int!]!
}

type TaskList{
    tasks: [Task!]!
}

type EmployeeList{
    employees: [Employee!]!
}

type CompanyList{
    companies: [Company!]!
}

union ResultUnion = Task | Employee | TaskList | EmployeeList | Company | CompanyList

type Pagination{
    total_rows:     Int64!
    total_pages:    Int!
    page:           Int!
    sort:           String!
    limit:          Int!
    rows:           ResultUnion
}
```
 We are completed with creating services and schemas neccessary for running graphqls now its turn for generating our equivalent go codes. But before that we need to do changes in our gqlgen.yml file to bind our custom created models and create some custom resolvers as well.

 > Step 7: Updating gqlgen.yml

1. Earlier we added autobind option uncomment below line and instead of **example/graph/model** write **example/models**. 

```yaml
autobind:
 - "example/models"
```
This will bind our custom created models with graphql model in short turining autobind on searches for the models in  the specified directory

2. After updating auto bind we need to bind grqphqls models with custom models created in golang so we will update below models line for linking models also keep it in mind that we will be creating custom resolvers for all models so we will add resolvers true for all models.

```yaml
models:
  Employee:
    model: 
      - example/models.Employee
    fields:
      tasks:
        resolver: true
      company:
        resolver: true
  Task:
    model: 
      - example/models.Task
    fields:
      employees:
        resolver: true
  Company:
    model: 
      - example/models.Company
    fields:
      employees:
        resolver: true
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.ID
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32
  Int:
    model:
      - github.com/99designs/gqlgen/graphql.Int
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int32

  models:
  Int64:
    model:
      - github.com/99designs/gqlgen/graphql.Int64
      - github.com/99designs/gqlgen/graphql.Int
```

`Note that I have also added Int64 datatype as pagination object needs datatype Int64`

 > Step 8: Generating equivalent go resolvers for all schemas. Run below command to auto generate all necessary resolvers needed.

 ```shell
 go run github.com/99designs/gqlgen generate
 ```

 You will see that all files in example/generated are auto created. Also, we need some changes in example/server.go file 
  
Search for below line

 ```go
 srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
 ```

 replace above line with below code you can see the minor change of replacing package name as we have restructure graph to generated we also need to do same changes in server.go file


 ```go
 srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &generated.Resolver{}}))
 ```

 > Step 9: Now we will create necessary functions in domain and add implementation logic.

 In this step I will not go deep by showing any logic instead I will just give you reference for which functions will be under which directory.

 I will create functions name same as auto created functions with same return values so you can identify in which resolver function needs to be called.

 But before that we will create a domain structure with following neccessary repositories that will also give you idea of which repositories are needed.

 ```go
type Domain struct {
    TaskRepository     repository.TaskRepository
    EmployeeRepository repository.EmployeeRepository
    M2MRepository      repository.M2MRepository
    ComapnyRepository  repository.ComapnyRepository
}
 ```

 I will also show you one sample structure of how I have created a repository structure

 ```go
 // example/database/repository/employee_task_repository.go
 type M2MRepository struct {
	DB *gorm.DB
}
 ```

 `Note that I have used gorm.io for database you can use any orm you want`

 Also add domain object in example/ generated/resolvers.go.Resolver file

 ```go
 type Resolver struct {
	Domain domain.Domain
}
 ```
Below are the functions with its respective files:

| Domain File Name  | Domain Function Name  | Resolver Function Name | Generated Resolver File |
|-------------------|-----------------------|------------------------|-------------------------|
| company_domain.go |  RegisterNewCompany   |     RegisterCompany    |   mutaion_resolver.go   | 
| company_domain.go |  RetrieveCompanies    |     Companies          |   query_resolver.go     |
| employee_domain.go|  CreateNewEmployee    |     CreateEmployee     |   mutaion_resolver.go   |
| employee_domain.go|  CreateNewEmployee    |     Employees          |   query_resolver.go     |              
| task_domain.go    |  CreateNewTask        |     CreateTask         |   mutaion_resolver.go   |
| task_domain.go    |  AssigntaskToEmployees|     AssignTask         |   mutaion_resolver.go   |
| task_domain.go    |  RetrieveTasks        |     Tasks              |   query_resolver.go     |    

You might have a question why create domain structure and create its functionality when we can directly implement logic in generated resolver files, yes you can directly implement logic in auto generated resolvers but there are some risks of doin it as it is auto generated your implmentation might get override when you regenerate the code.

Think of a scenerio where you have written more than 50 lines of code and your implementation get override. Hence, I prefer creating a logic in different file although gqlgen tool gives us recvery optioin but why take chance its like prevention is better than cure. 

Also if we implement our logic in different file would also give us some advantage as well where we want to create same application with some other technology for example REST API you can directly call this functions and no need to rewrite the whole code again for REST API as implementation logic will remain same.

> Step 10: Add created domain to the resolvers in example/server.go file

You might find below line
```go
srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &generated.Resolver{}}))
```
Replace it with
```go
srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &generated.Resolver{
    Domain: domain.Domain{
        TaskRepository:     repository.TaskRepository{DB: DB},
        EmployeeRepository: repository.EmployeeRepository{DB: DB},
        M2MRepository:      repository.M2MRepository{DB: DB},
        ComapnyRepository:  repository.ComapnyRepository{DB: DB},
        },
    }}))
```
I am not showing you database implementaion neither how to call it in as I know you are smart enough to predict how I did that. So dont forget to add database implementation as I was using postgres I have added function in example/database/postgres/postgres.go file you can do the same for any DB Engines and call DB variable in main function

> Step 10: Custom Resolvers Implementation

A question might arrise why we need Custom Reolvers or What is the use of custom resolvers?

-> To answer this question, you need to recall one of the structure you have created earlier. I will take employee as an example

```graphql
type Employee {
    id:             ID!
    name:           String!
    designation:    String!
    email:          String!
    company_id:	    Int!
    company:	    Company
    tasks:          [Task!]
}
```
Here, you can see we have two sub objects that is Task which is in slice of Task and Company.

What Resolver does is it simply fetches you the data of this sub objects or models based on the foreign key relation and the logic you provide.

Yes, I know some of you might get a question why need of resolvers we can directly fetch it when we call our main query, yes you can do that but think of a scenerio where a sub object has many variables or a object has many objects, in this case database query would get too long and might not be a feasible option where a user only needs data for main object not for sub object.

You might ask then we can simply create a new service for fetching that data. Again my answer will be yes you can create that but thats the beauty of graphql you don't need to do that as custom resolver can solve that problem why call different queries when you can get data in same service.

I hope now your doubt are clear, so moving on to implementation part of custom resolvers.

I am using [DataLoaden library by vektah](https://github.com/vektah/dataloaden), you can use any thing is you want else directly write implementation logic in custom resolvers

We need to create three data loaders i.e. for company, task and employee

First of all you need to get the library so run,
```shell
go get -u github.com/vektah/dataloaden
```
Now for creating dataloaders, first of all I will give you a brief how to run it.

```shell
go run github.com/vektah/dataloaden CompanyLoader int '*example/models.Company'
```
Above is the how you run dataloaden to get CompanyLoader
The above command has 4 parts `go run github.com/vektah/dataloaden` line will run the dataloaden `CompanyLoader` is the name you want to give to your dataloader `int` is you can say is the primary key's datatype but infact it is the datatype of the variable which we will be using to fetch the Data of that Object and last  `*example/models.Company` is the datatype of the object or the object you want to return.

You have already seen the working of datatloaden so lets create data loader for all models.

1. First of all go to dataloader directory
```shell
cd dataloader
```
2. Create a empty file dataloader.go, if you have not created and write
```go
package datatloader
```
3. Run Below three commands
```shell
go run github.com/vektah/dataloaden CompanyLoader int '*example/models.Company'

go run github.com/vektah/dataloaden EmployeeLoader int '[]*example/models.Employee'

go run github.com/vektah/dataloaden TaskLoader int '[]*example/models.Task'
```
This command will generate dataloaders for all objects. That all you are done with creating data loaders now you just need to integrate data loader for that you need to create a middleware for dataloader before that we will be doing some changes with example/server.go.

> Step 11: Adding a framework or router for using middleware

I am using GIN Framework for using middleware, for simplicity but you can use any framework which uses http library.

`Note that you cannot directly use framework which doesnot implement http library.`

Add GIN library using,
```shell
 go get -u github.com/gin-gonic/gin
```
Also add GIN CORS library,
```shell
go get -u github.com/gin-contrib/cors
```
Modify example/server.go file with below code
```go
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strconv"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "gorm.io/gorm"
    "example/database/postgres"
    "example/database/repository"
    "example/dataloaders"
    "example/domain"
    "example/generated"
    "example/models"
)

const defaultPort = "8080"

func graphqlHandler(DB *gorm.DB) gin.HandlerFunc {
    srv := handler.NewDefaultServer(generated.NewExecutableSchema(
        generated.Config{Resolvers: &generated.Resolver{
            Domain: domain.Domain{
            TaskRepository:     repository.TaskRepository{DB: DB},
            EmployeeRepository: repository.EmployeeRepository{DB: DB},
            M2MRepository:      repository.M2MRepository{DB: DB},
            ComapnyRepository:  repository.ComapnyRepository{DB: DB},
    }}}))
    return func(ctx *gin.Context) {
        handlerFunc := dataloaders.DataLoaderMiddleware(DB, srv)
        handlerFunc.ServeHTTP(ctx.Writer, ctx.Request)
    }
}

func playGroundHandler() gin.HandlerFunc {
    handlerFunc := playground.Handler("GraphQL playground", "/query")
    return func(ctx *gin.Context) {
        handlerFunc.ServeHTTP(ctx.Writer, ctx.Request)
    }
}

func init() {
    if err := godotenv.Load(".env"); err != nil {
        panic("unable to load .env file")
    }
}

func main() {
    dropCreate := flag.Bool("db_drop_create", false,
        "Specifies if database needs to drop"+
        "  table before migration while starting application.")
    dbUp := flag.Bool("db_up", false,
        "Specifies if database needs migrations while starting application.")
    dbDown := flag.Bool("db_down", false,
        "Specifies if database needs to drop while starting application.")
    flag.Parse()
    DB, err := postgres.New(os.Getenv("DB_DSN"))
    if err != nil {
        panic(err)
    }
    if *dbDown || *dropCreate {
        log.Println("Droping All Tables")
        err = DB.Migrator().DropTable(&models.Company{}, &models.Employee{}, 
            &models.Task{})
        if err != nil {
            panic(err)
        }
    }
    if *dbUp || *dropCreate {
        log.Println("Auto Migrating All Tables")
        err = DB.AutoMigrate(&models.Company{}, &models.Employee{}, &models.Task{})
        if err != nil {
            panic(err)
        }
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }
    models.PageLimit, _ = strconv.Atoi(os.Getenv("PAGE_LIMIT"))
    models.OrderBy = os.Getenv("ORDER_BY")

    engine := gin.Default()
    engine.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:" + port},
        AllowCredentials: true,
    }))
    engine.POST("/query", graphqlHandler(DB))
    engine.GET("/", playGroundHandler())
    go fmt.Printf("[GIN-debug] Connect to http://localhost:%s/ for GraphQL playground\n", 
        port)
    if engine.Run() != nil {
        panic("unable to run gin engine")
    }
}
```
Do not panic for `dataloaders.DataLoaderMiddleware(DB, srv)` as our next step is creating this middleware.

> Step 12: Creating DataLoader Middleware  

Add below code in example/dataloaders/dataloader.go file
```go
package dataloaders

import (
	"context"
	"net/http"
	"time"
	"gorm.io/gorm"
	"example/models"
)

var (
	EmployeeLoaderByTaskKey    = "employeeloaderbytask"
	EmployeeLoaderByCompanyKey = "employeeloaderbycompany"
	taskLoaderKey              = "taskloader"
	companyLoaderKey           = "companyloader"
)

func DataLoaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        emplyeeLoaderByTask := EmployeeLoader{
            maxBatch: 100,
            wait:     1 * time.Millisecond,
            fetch: func(keys []int) (employees [][]*models.Employee, errs []error) {
                taskKeyMap := make(map[int][]*models.Employee, len(keys))
                for i := range keys {
                    var emps []*models.Employee
                    err := db.Model(models.Employee{}).
                        Joins("inner join employee_tasks on employee_id = id").
                        Where("task_id = ?", keys[i]).Find(&emps).Error
                    if err != nil {
                        errs = append(errs, err)
                        return
                    }
                    taskKeyMap[keys[i]] = emps
                }
                result := make([][]*models.Employee, len(keys))
                for i, id := range keys {
                    result[i] = taskKeyMap[id]
                }
                return result, nil
            },
        }
        taskLoader := TaskLoader{
            maxBatch: 100,
            wait:     1 * time.Millisecond,
            fetch: func(keys []int) (tasks [][]*models.Task, errs []error) {
               employeeKeyMap := make(map[int][]*models.Task, len(keys))
                for i := range keys {
                var tsks []*models.Task
                err := db.Model(models.Task{}).
                    Joins("inner join employee_tasks on task_id = id").
                    Where("employee_id = ?", keys[i]).Find(&tsks).Error
                if err != nil {
                    errs = append(errs, err)
                    return
                }
                employeeKeyMap[keys[i]] = tsks
            }
            result := make([][]*models.Task, len(keys))
            for i, id := range keys {
                result[i] = employeeKeyMap[id]
            }
            return result, nil
            },
        }
        emplyeeLoaderByCompany := EmployeeLoader{
            maxBatch: 100,
            wait:     1 * time.Millisecond,
            fetch: func(keys []int) (employees [][]*models.Employee, errs []error) {
                companyKeyMap := make(map[int][]*models.Employee, len(keys))
                for i := range keys {
                    var emps []*models.Employee
                    err := db.Model(models.Employee{}).
                        Where("company_id = ?", keys[i]).Find(&emps).Error
                    if err != nil {
                        errs = append(errs, err)
                        return
                    }
                    companyKeyMap[keys[i]] = emps
                }
                result := make([][]*models.Employee, len(keys))
                for i, id := range keys {
                    result[i] = companyKeyMap[id]
                }
                return result, nil
            },
        }
        companyLoader := CompanyLoader{
            maxBatch: 100,
            wait:     1 * time.Millisecond,
            fetch: func(keys []int) (companies []*models.Company, errs []error) {
                err := db.Model(models.Company{}).Where("id IN ?", keys).Find(&companies).Error
                if err != nil {
                    errs = append(errs, err)
                    return
                }
                companyMap := make(map[int]*models.Company, len(keys))
                for i := range companies {
                    companyMap[companies[i].ID] = companies[i]
                }
                result := make([]*models.Company, len(keys))
                for i, id := range keys {
                    result[i] = companyMap[id]
                }
                return result, nil
            },
        }
        ctx := context.WithValue(r.Context(), EmployeeLoaderByTaskKey, &emplyeeLoaderByTask)
        ctx = context.WithValue(ctx, taskLoaderKey, &taskLoader)
        ctx = context.WithValue(ctx, EmployeeLoaderByCompanyKey, &emplyeeLoaderByCompany)
        ctx = context.WithValue(ctx, companyLoaderKey, &companyLoader)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func GetEmployeeLoader(ctx context.Context, key string) *EmployeeLoader {
    return ctx.Value(key).(*EmployeeLoader)
}

func GetTaskLoader(ctx context.Context) *TaskLoader {
    return ctx.Value(taskLoaderKey).(*TaskLoader)
}

func GetCompanyLoader(ctx context.Context) *CompanyLoader {
    return ctx.Value(companyLoaderKey).(*CompanyLoader)
}
```
You might be thinking we have created middleware which stores data loader but when do we call GetLoader functions to fetch the Object data.No worries I will show you how to call it in next step.


> Step 13: Invoking Dataloaders in Custom Resolvers

I will show you how to invoke  custom resolver only in one resolver as the way of invoking is very simple you might get it how to do the same for rest others.

Lets take a simple example of laptop resolver to fetch Employees:

Go to example/generated/company_resolver.go, there you will see Employees Function just remove panic code and add below code

```go
return dataloaders.GetEmployeeLoader(ctx, dataloaders.EmployeeLoaderByCompanyKey).Load(obj.ID)
```

And that's it you are done with Laptop Resolver.

Incase, you get stuck with any resolver below is the code for all rest resolvers,

```go
// example/generated/employee_resolver.go
// Inside Company Function
return dataloaders.GetCompanyLoader(ctx).Load(obj.CompanyID)


// Inside Task Function
return dataloaders.GetTaskLoader(ctx).Load(obj.ID)
```
```go
// example/generated/task_resolver.go
// Inside Employees Function
return dataloaders.GetEmployeeLoader(ctx, 
    dataloaders.EmployeeLoaderByTaskKey).Load(obj.ID)
```
If You have come this far, congratulations you have successfully implemented all functionalities needed to run all services.

 `Note: If you get stuck anywhere you can refer to my github repo`
### Finally Running the application
Yes, your wait is over, we are finally going to ru our application.
To run our application simply run 
```shell
go run server.go
```
In terminal you will get GraphQL Playground locally hosted link,
click on that link you will redirect on to a GraphQL Playground where you can perform all GraphQL related queries.

First of all I will show you two different ways to call a Query or Mutation

I will take query type for Companies service as an example to demonstrate two different ways:

> 1. Directly giving value to the Service Called

```graphql
query DirectInputCompanies {
  companies(
    filter: {
      employee_email: ".m", 
      company_name: "sim", 
      company_website: ".com", 
      employee_name: "ab"
    },
    page: 1
  ) {
    total_rows
    total_pages
    page
    sort
    limit
    rows {
      ... on  CompanyList{
        companies{
          name
          website
        }
      }
    }
  }
}
```

   Here, you can see that I have called companies query by directly passing parameter values to the call this is the one and the simple way to call any service either of mutation or query type.

   Another way is also simple instead of directly giving parameters value in the service method you declare a temporary variable and pass value to that variable.

> 2. Passing  Values By Temporary Parameter to the service called

```graphql
query Companies($filter:FilterCompanyInput,$page:Int){
  companies(filter:$filter,page:$page){
    total_rows
    total_pages
    page
    sort
    limit
    rows {
      ... on CompanyList {
        companies {
          id
          name
         employees{
          name
          designation
          email
          }
        }
      }
    }
  }
}
```

```json
//Variables
{
  "page": 1,
  "filter": {
    "employee_email": ".m",
    "company_name": "sim",
    "company_website": ".com",
    "employee_name": "ab"
  }
}
```

    
   In the above function you can see the neat and clear query without any value added in the service called itself, instead passed values to the temporary vriables created by the service.

   To add the values go to down to the Variables Tab and paste this Variables Code.

As you can see there is no difference in how you call the query, both does the same thing.

Now just for reference I will show you ho to invoke mutation service below is the example of it.
```graphql
mutation RegisteCompany($input: CreateCompany!) {
  registerCompany(input: $input) {
    id
    name
    website
  }
}
```

```graphql
// Variable
{
  "input": {
    "name": "Example",
    "website": "example.com"
  }
}
```

### Conclusion
In this article, we covered overview of what is [graphql](https://graphql.org/), creating [GraphQL](https://graphql.org/) using [gqlgen tool](https://gqlgen.com/) and [golang](https://gqlgen.com/). Also we learned different features of [gqlgen tool](https://gqlgen.com/), auto binding models created in [golang](https://gqlgen.com/) with [GraphQL](https://graphql.org/) types, creating custom resolvers useful for auto fetching sub-object's data, using **`UNION`** type demonstraing how pagination can be sent to response which can fetch you different object's data which reduces creating pagination type for all data differently and different wats to invoke graphql query. 


