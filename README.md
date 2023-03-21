# Nepackage

A tool to have a main place where start a new project

## Requeriments
- Go 1.19.* o superior
- Docker
- MySQL 8
- Environments variables:
  - `DSN` "connectionString" to MySQL database. Default to `''`
  - `JWTKEY` JWT key to sign tokens
- Gow package for development

## How to start project

#### Development
1. Clone Nepackage repository
```
git clone {repository} 
```
2. cd to direcotory
```
cd ./nepackage
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
You can find image on [Docker Hub](https://hub.docker.com/r/otisnado/nepackage)
```
docker run -d -p 9090:8080 --name nepackage otisnado/nepackage:latest
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
- [x] In memory default SQLite database support
- [x] CI Pipeline in GitHub Actions
- [x] Docker Hub repository
- [ ] Base project templates
- [ ] Java language support (To generate projects)
- [ ] Golang language support (To generate projects)
- [ ] Base CI/CD templates for GitHub Actions
- [ ] Base CI/CD templates for Azure DevOps Pipelines
- [x] Spring CLI support to generate projects based on Spring Boot
- [ ] Buffalo CLI support to generate prjects based on Buffalo Framework
- [ ] Connect to GitHub to create projects repositories
- [ ] Connecto to Azure DevOps Repos to create projects repositories
- [x] Open API v2 definition
- [x] Swagger endpoint
- [x] User login support
- [x] JWT Support
- [x] Default admin user creation at first start
- [x] RBAC