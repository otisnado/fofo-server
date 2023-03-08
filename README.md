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
  - `JWTKEY` JWT key to sign tokens
- Gow package for development

## How to start project

#### Development
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

## Start Docker container
You can find image on [Docker Hub](https://hub.docker.com/r/otisnado/fofo-server)
```
docker run -d -p 9090:8080 -e DBHOST=localhost -e DBPORT=3306 -e DBNAME=fofo -e DBUSER=root -e DBPASS=1234 -e DBTZ=America%2FEl_Salvador --name fofo-server otisnado/fofo-server:latest
```

## List of features
- [x] Projects CRUD
- [x] Languages CRUD
- [x] Users CRUD
- [x] Groups CRUD
- [x] Docker support
- [ ] Docker compose file
- [x] Mysql database support
- [ ] PostgreSQL database support
- [x] CI Pipeline in GitHub Actions
- [x] Docker Hub repository
- [ ] Base project templates
- [ ] Java language support (To generate projects)
- [ ] Golang language support (To generate projects)
- [ ] Base CI/CD templates for GitHub Actions
- [ ] Base CI/CD templates for Azure DevOps Pipelines
- [ ] Spring CLI support to generate projects based on Spring Boot
- [ ] Buffalo CLI support to generate prjects based on Buffalo Framework
- [ ] Connect to GitHub to create projects repositories
- [ ] Connecto to Azure DevOps Repos to create projects repositories
- [x] Open API v2 definition
- [x] Swagger endpoint
- [ ] User login support
- [x] JWT Support
- [x] Default admin user creation at first start
- [ ] RBAC