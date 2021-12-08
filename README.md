<p align="center">
  <img alt="Stampede" src="assets/logo.svg" />
</p>

<p align="center">
<a href="https://img.shields.io/badge/license-MIT-%23339933"><img src="https://img.shields.io/badge/license-MIT-%23339933" alt="License"></a>
<a href="https://img.shields.io/github/contributors/bashovski/stampede-cli?color=%23011762"><img src="https://img.shields.io/github/contributors/bashovski/stampede-cli?color=%23011762" alt="Contributors"></a>
</p>

# Stampede CLI 🦕
- To see the source code of Stampede Framework, check https://github.com/bashovski/stampede

# About
- Stampede is a framework, or an eco-system written in TypeScript for Deno, made with emphasis on delivering new features quicker than ever before.
- CLI for Stampede is instrumental for creating a new Stampede project as improves overall dev-experience by allowing you to create important modules for your next project.

# Installation

```shell script
brew tap bashovski/stampede
brew install stampede
stampede --help # Test if it works
```

## Building From Source

Prerequisite: Install Go

```
git clone https://github.com/bashovski/stampede-cli.git
cd stampede-cli
go mod init bashovski/stampede-cli
go build
sudo mv stampede-cli /usr/local/bin/stampede
```

Now you can run stampede-cli: ```stampede --help```.
