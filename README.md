# Extract version info from git repo.

This small program will generate a file named gitversion.go
when it is run in a directory with git support.

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

``` go install github.com/jkvatne/version@latest``` 

And ensure that the exe file created is available in the path.


In the Goland IDE, to run it automatically before building the exe file:
- Edit the settings with Ctrl-Alt-S, Tools, External Tools, then click '+'
- Set name and description. Set Program to 'version.exe', and click OK
- Set working directory to '$ProjectFileDir$'
- Edit the run configuration(s)
- At the bottom, in the 'Before launch' list, click '+'
- Select 'Run external tool' and select the tool you added (version).")To run the program in Goland automaticaly for each build,
add an external tool named ``` version.exe``` in the run configuration,

(C)Jan Kåre Vatne 2024 (jkvatne@online.no)
