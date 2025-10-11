# Dev command utilities


## Install

```shell
brew tap cgaube/devcommands
```

## Autocompletion

Add this to ~/zshrc

```
# Enable tab completion for the 'dev' command, if it exists                      
if command -v dev &> /dev/null; then                                             
  source <(dev completion zsh)                                                   
fi       
```
