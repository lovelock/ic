# Interactive SSH client for macOS and Linux users.

## Why I do this?

As an OS X user, I have been eagering for a friendly SSH client for long.

We have iTerm as psudo terminal, which does perfect.
We have open ssh client for connecting to remote hosts, which does quite well, but I have to maintain the config file forever.

Things are much the same to Linux users.

Why no one can make it easier?

## Installation

### `homebrew`

```bash
brew tap https://github.com/lovelock/homebrew-ic
brew install ic
```

### `go get`

```bash
go get github.com/lovelock/ic
```

and then add one line to your `.zshrc` or `.bash_profile` according to the shell you're using.

`PROG=ic source $GOPATH/src/github.com/lovelock/ic/autocomplete/zsh_autocomplete`

or

`PROG=ic source $GOPATH/src/github.com/lovelock/ic/autocomplete/bash_autocomplete`

then `source ~/.zshrc` or `source ~/.bash_profile`.

## Usage

Just type `ic` and press enter, it will tell you almost everything you should know to use it.
