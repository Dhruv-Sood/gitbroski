# 🐙 Git-Broski

**Broski for your Git!**  
A CLI tool to perform common Git tasks with simple commands.

## Features

- Open your current Git repository in your default browser
- Simplify repetitive Git tasks with single commands
- Fully cross-platform

## Installation

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
go build -o gitbroski ./cmd
```

### 4. (Optional) Create a Symlink for Global Use
```bash
sudo ln -s /full/path/to/gitbroski /usr/local/bin/gitbroski
```
> This allows running `gitbroski` from **any directory**.

## Usage
Open the current Git repository in your default browser:

```bash
gitbroski open
```

## Contributing
We welcome contributions! Check out our [CONTRIBUTING.md](CONTRIBUTING.md) guide to get started.

## License
MIT

