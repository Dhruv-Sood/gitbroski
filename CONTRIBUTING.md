# Contributing to Git-Broski

First off, thank you for considering contributing to Git-Broski! Your help makes this project better for everyone.

Following these guidelines will make the contribution process smooth and efficient for all contributors.

## How to Contribute

### 1. Fork and Clone the Repository
Fork the project on GitHub and then clone it locally:

```bash
git clone https://github.com/your-username/gitbroski.git
cd gitbroski
```

### 2. Install Dependencies
Ensure all dependencies are installed:

```bash
go mod tidy
```

### 3. Build the Project
Build the CLI tool locally:

```bash
go build -o gitbroski ./cmd
```

### 4. (Optional) Create a Symlink for Global Use
This allows you to run `gitbroski` from any directory:

```bash
sudo ln -s /full/path/to/gitbroski /usr/local/bin/gitbroski
```

### 5. Make Your Changes
- Create a new branch for your feature or bugfix:
```bash
git checkout -b feature/my-feature
```
- Make your changes and test thoroughly.
- Commit your changes with clear, descriptive messages:
```bash
git commit -m "Add feature X for better usability"
```

### 6. Submit a Pull Request
Push your branch to your fork and open a pull request to the main repository. Make sure to explain your changes clearly.

## Reporting Issues
If you encounter a bug or have a suggestion:
1. Go to the [Issues tab](https://github.com/your-username/gitbroski/issues)
2. Click **New Issue**
3. Provide a descriptive title and detailed description.

## Code of Conduct
Please be respectful and kind in all interactions. Follow the [Code of Conduct](CODE_OF_CONDUCT.md) to help us maintain a welcoming environment.

Thank you for contributing! 🚀

