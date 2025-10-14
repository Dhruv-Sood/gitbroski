# Git-Broski 🚀

<p align="center">
  <img src="https://raw.githubusercontent.com/gitbroskie/gitbroski/main/assets/logo.png" alt="GitBroski Logo" width="120"/>
</p>

<p align="center">
  <strong>Broski for your Git!</strong><br>
  <em>A CLI tool to perform various manual tasks with single commands</em>
</p>

<p align="center">
  <a href="https://github.com/gitbroskie/gitbroski/releases"><img src="https://img.shields.io/github/v/release/gitbroskie/gitbroski?style=flat-square" alt="Release"></a>
  <a href="https://github.com/gitbroskie/gitbroski/blob/main/LICENSE.md"><img src="https://img.shields.io/github/license/gitbroskie/gitbroski?style=flat-square" alt="License"></a>
  <a href="https://github.com/gitbroskie/gitbroski/issues"><img src="https://img.shields.io/github/issues/gitbroskie/gitbroski?style=flat-square" alt="Issues"></a>
</p>

---

## Overview

Git-Broski is a command-line tool that streamlines your Git workflow by automating common tasks with simple commands.

> Project Notion Page: https://gratis-neon-644.notion.site/GitBroski-Broski-for-your-Github-Workflow-286e41747653800a9cc4f7b014c6cf51

---

## Installation

### Prerequisites

- Node.js (for npm distribution) or Go (to build locally)

### Install via npm

```bash
npm install -g gitbroski
```

### Install via Go (Local Build)

1. **Clone the Repository**
		```bash
		git clone https://github.com/gitbroskie/gitbroski.git
		cd gitbroski
		```
2. **Install Dependencies**
		```bash
		go mod tidy
		```
3. **Build the Project**
		```bash
		go build -o gitbroski ./cmd
		```
4. **(Optional) Symlink for Global Use (Linux/macOS)**
		```bash
		sudo ln -s /full/path/to/gitbroski /usr/local/bin/gitbroski
		```

### Windows Setup

1. **Clone the Repository**
		```powershell
		git clone https://github.com/gitbroskie/gitbroski.git
		cd gitbroski
		```
2. **Build or Download the Binary**
		- If Go is installed:
			```powershell
			go build -o gitbroski.exe ./cmd
			```
		- If Go is not installed: Download `gitbroski-windows.exe` from the project folder or [Releases](https://github.com/gitbroskie/gitbroski/releases).
3. **Add to PATH (Symbolic Link)**
		- Command Prompt (Admin):
			```cmd
			setx PATH "%PATH%;C:\path\to\your\bin"
			```
		- Git Bash (Admin):
			```bash
			ln -s /c/Users/<YourUserName>/gitbroski/gitbroski.exe /usr/bin/gitbroski
			```
            > ⚠️ Replace /c/Users/<YourUserName>/gitbroski/ with your actual path.
5. **Verify Installation**
		```bash
		gitbroski --version
		```
		Expected output:
		```text
		GitBroski CLI v1.0.0
		```

---

## 🚀 Usage

### 1. Open the Remote Repository

Quickly jump from your command line to the GitHub or GitLab page for your current project.

```bash
gitbroski open
```

### 2. Auto-Generate .gitignore

Effortlessly create a .gitignore file based on a specified language or technology.

```bash
gitbroski ignore <Language>
```

Currently supported:
- python

> Issues are open to add support for Node.js, Golang, and more!

### 3. Easy Empty Commit

Create an empty Git commit (useful for triggering CI/CD pipelines without code changes) and add an optional message.

```bash
gitbroski empty commit <your-Message>
```

---

## 🤝 Contributing

1. Fork and clone the repository.
2. Install dependencies:
		```bash
		go mod tidy
		```
3. Build the project:
		```bash
		go build -o gitbroski ./cmd
		```
4. (Optional) Create a symlink for global use:
		```bash
		sudo ln -s /full/path/to/gitbroski /usr/local/bin/gitbroski
		```
5. Run gitbroski from any folder and contribute!

---

<p align="center">
	<i>Made with ❤️ by the GitBroski community</i>
</p>
