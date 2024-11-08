# Extract verision info from git.

This small program will generate a file named gitversion.go
when it is run in any directory with git support.

The file gitversion.go will have "package main" at the top, so it will be 
part of the main package automaticaly.

The only content of the generated "giversion.go" file is a constant ```Version```
The constant is a string consisting of the latest tag when no changes are 
commited after the tagged commit.

It the tag is not on the latest commit, the number of commits is added, 
together with the hash. If local changes are present after the commit, 
the ```-dirty``` string is added.  

Example:
```
package main

// Version is extracted from the latest tag in the git repository
const Version="V0.1-107-gce0a8d27-dirty"
```

To use this program, run 
```
go install
``` 

And ensure that the exe file created is available in the path.

Add the following line to the pre-build script, or for GoLand, 
add an external tool named ``` version.exe``` in the run configuration,
and set working directory to ``` $ProjectFileDir$``` 
 