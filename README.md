# Fofo Server

A server to have main place where start a new project


## Requeriments
- Go 1.19.* o superior
- Docker
- MySQL 8
- Environments variables:
  - `DBHOST` (Default localhost)
  - `DBPORT` (Default 3306)
  - `DBNAME` (Default fofo)
  - `DBUSER` (Default root)
  - `DBPASS` (Default 1234)
  - `DBTZ` (Default America%2FEl_Salvador)
- Gow package for development

## How to start project

#### Devlopment
1. Clone FoFo Server repository
```
git clone {repository} 
```
2. cd to direcotory
```
cd ./fofo-server
```

3. Donwloag dependencies
```
go get
```

4. Install Gow package
```
go install github.com/mitranim/gow@latest
```
5. Run main.go with Gow
```
gow run main.go
```

#### Docker
```
docker run -d -p 9090:8080 -e DBHOST=localhost -e DBPORT=3306 -e DBNAME=fofo -e DBUSER=root -e DBPASS=1234 -e DBTZ=America%2FEl_Salvador --name fofo-server otisnado/fofo-server:latest
```

## API Endpoints
This API has the nexts endpoints:

#### /projects
- GET -- Get a list of all projects created with this tool, return the info in JSON
- POST -- Add a new project, requires JSON structure

### /projects/:id
- PATCH -- Update data for a projects, requires project id

#### /projects/:id
- GET -- Get a project by its ID, return the info in JSON


## List of features
- [x] Projects CRUD
- [x] Languages CRUD
- [x] Docker support
- [ ] Docker compose file
- [ ] Mysql database support
- [ ] PostgreSQL database support
- [ ] Base project templates
- [ ] Java language support (To generate projects)
- [ ] Golang language support (To generate projects)
- [ ] Base CI/CD templates for GitHub Actions
- [ ] Base CI/CD templates for Azure DevOps Pipelines
- [ ] Spring CLI support to generate projects based on Spring Boot
- [ ] Buffalo CLI support to generate prjects based on Buffalo Framework
- [ ] Connect to GitHub to create projects repositories
- [ ] Connecto to Azure DevOps Repos to create projects repositories
- [ ] User login support
