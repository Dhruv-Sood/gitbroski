# Git-Broski 

https://gratis-neon-644.notion.site/GitBroski-Broski-for-your-Github-Workflow-286e41747653800a9cc4f7b014c6cf51

**Broski for your Git!**  
A CLI tool to perform various manual tasks with single commands

## Installation
This section guides you through setting up gitbroski locally.

### Prerequisites
1. Ensure Node.js is installed and configured on your system.

### Local Setup
Install the tool globally using npm by running the following command in your terminal:
```bash
npm install -g gitbroski
```
## Usage
Here are the key commands for gitbroski to enhance your Git workflow:

#### 0. Show Version

Print the installed gitbroski version. Supports aliases `version`, `--version`, and `-v`.
```bash
gitbroski --version
# or
gitbroski version
# or
gitbroski -v
```

#### 1. Open the Remote Repository

Quickly jump from your command line to the GitHub or GitLab page for your current project.
```bash
gitbroski open
```
#### 2. Auto-Generate .gitignore
Effortlessly create a .gitignore file based on a specified language or technology.
```bash
gitbroski ignore <Language>
```
Currently, this command supports:
- python
- (Issues are open to add support for node.js, golang, and many more languages!)

#### 3. Easy Empty Commit
Create an empty Git commit (useful for triggering CI/CD pipelines without code changes) and add an optional message.
```bash
gitbroski empty commit <your-Message>
```


## Installation (Local setup)

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/gitbroski.git
cd gitbroski
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Build the Project
```bash
# Basic build (version defaults to "dev")
go build -o gitbroski ./cmd

# Optional: embed version at build time
go build -ldflags "-X gitbroski/internal/version.Version=$(jq -r .version package.json 2>/dev/null || echo 0.0.0)" -o gitbroski ./cmd
```

### 4. (Optional) Create a Symlink for Global Use
```bash
sudo ln -s /full/path/to/gitbroski /usr/local/bin/gitbroski
```
> This allows you to run `gitbroski` from **any directory**.

- Opens the **current Git repository** in your default browser.

## Contribution Guide
1. **Fork** and **clone** the repository.
2. Install dependencies:
```bash
go mod tidy
```
3. Build the project:
```bash
go build -o gitbroski ./cmd
```
4. Create a symlink for global use (optional):
```bash
sudo ln -s /full/path/to/gitbroski /usr/local/bin/gitbroski
```
5. Run `gitbroski` from **any folder** and contribute!

## License
MIT
