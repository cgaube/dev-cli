# Dev command utilities

## What ?

Dev CLI is a command-line interface that uses Square's Exoskeleton framework to create a unified `dev` binary. It automatically discovers and executes subcommands from `devcommands` folders in your PATH, allowing you to manage multiple development tools through a single interface.

This approach enables:
- A centralized command entrypoint for all development utilities
- Automatic discovery of subcommands from `devcommands` folders without manual registration
- Support for subcommands written in any language (Node.js, Go, Python, Bash, etc.)
- Tab completion and help system integration
- Modular architecture where each tool can be developed and released independently

The `dev` binary acts as a dispatcher that finds available commands in `devcommands` directories, executes them with appropriate arguments, and returns their exit codes.

## Install

```shell
brew tap cgaube/devcommands
brew install dev-cli
```

## Usage

Once installed, you can use the `dev` command to execute various subcommands:

```shell
# List available commands
dev help

# Execute a specific command
dev <command>

# Get completion for a command
dev <command> --help
```

The dev CLI will automatically discover commands from `devcommands` folders and make them available through the unified interface.

## Command Discovery

Commands are discovered automatically from `devcommands` folders in your PATH. Each command can be:

- A standalone executable binary
- A shell script with proper shebang
- A directory containing an `.exoskeleton` file (for submenus)

This modular approach allows you to organize your development tools in a flexible, extensible way.

### Example Structure

```
~/Projects/
├── devcommands/           # Global commands available in all projects
│   ├── deploy
│   └── test
└── myproject/
    ├── src/
    └── devcommands/       # Commands specific to this project
        └── build
```

Commands in `~/Projects/devcommands` are available globally, while commands in `~/Projects/myproject/devcommands` are only available when working in that project (if that folder is in your PATH).

## Autocompletion

Add this to ~/zshrc

```
# Enable tab completion for the 'dev' command, if it exists
if command -v dev &> /dev/null; then
  source <(dev completion zsh)
fi
```
