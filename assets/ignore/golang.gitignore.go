package ignore

func Golang() string {
	return `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Distribution / packaging
build/
 
""" Test binary, built with go test -c"""
*.test

# Code coverage profiles and other test artifacts
*.out
coverage.*
*.coverprofile
profile.cov

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
go.work.sum

# Flask stuff
instance/
.webassets-cache


# Jupyter Notebook stuff
.ipynb_checkpoints

# Virtual environment
.venv/
venv/
env/
/venv
/env

# Secrets/Configuration (should almost always be ignored)
.env
.flaskenv
settings.ini

# IDE-specific directories and files
# VS Code
.vscode/*
!.vscode/settings.json
!.vscode/tasks.json
!.vscode/launch.json
!.vscode/extensions.json
.history/
# JetBrains IDEs (PyCharm, IntelliJ, etc.)
.idea/

# Windows
Thumbs.db
ehthumbs.db
`
}
