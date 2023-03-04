# Fofo Server

A server to have main place where start a new project


## Requeriments
## How to start project
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
