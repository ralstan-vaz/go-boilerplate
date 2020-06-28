## Go-Boilerplate

An easier way to bootstrap web projects.  
The philosophies in this boilerplate are not official standards defined by the Go core dev team.  
They are purely defined through my experience in Go in the past years  

## Concepts Covered
The following concepts are currently covered in this boilerplate.  
The rest will be added in time.
- [x] Directory structure
- [x] Dependency injection
- [x] Request interceptors
- [x] Vendoring
- [x] Commenting
- [ ] Logging
- [ ] Error
- [ ] Tracing
- [ ] Documenting
- [ ] Unit Testing
- [ ] Integration testing



## Directory structure

### apis
It contains controllers for the pkgs that need to be exposed. Each protocol can have different implementations. Each protocol's implementation can reside in its respective directories.

### config
It contains the configuration model. This model will be used in the project.
.yml is used as the configuration file as it is much readable and also supports comments.
The yml config bind to the model which then passed through the project.

### initiate
It would contain code to start up the project. All dependencies would also be created here and then passed to the respective packages.

### pkg
It would mainly contain library code. It could contain multiple packages some of which may depend on other packages. Eventually, some pkg will be used by /apis to get exposed.

### builder 
It contains the scripts need to build the code. Containerization scripts would reside here.

### vendor
All dependencies would be here.

## Flow
The flow of the program would look like this.

```js
 main.go --> initiate --> config --> apis --> pkg
```
