# cleanstack

cleanstack is an attempt of implementing clean architecture in go project

the concept introduce by Robert C. Martin (Uncle Bob) 
more at : https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

# Description

Rules Of clean architecture are :

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

# Project cleanstack

This project has 4 Domain layer :

- Domain Layer
- Repository Layer
- Usecase Layer
- Presenter Layer

![clean architecture](https://github.com/maps90/cleanstack/raw/main/clean-archpng)

### Domain Layer
domain layer will used in all layer. this layer will store any presenter api contract, it also acts as entity / model for the repository.

### Repository Layer

Repository will store any Database handler. Querying, or Creating/ Inserting into any database will stored here. This layer will act for CRUD to database only. No business process happen here. Only plain function to Database.

### Usecase Layer

This layer will act as the business process handler. Any process will handled here. This layer will decide, which repository layer will use. And have responsibility to provide data to serve into delivery. Process the data doing calculation or anything will done here.

### Presenter Layer
This layer will act as the presenter. Decide how the data will presented. Could be as REST API, or HTML File, or gRPC whatever the delivery type. Validate the input and sent it to Usecase layer.









