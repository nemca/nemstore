# nemstore

The `nemstore` utility allows you to work with encrypted files on the file system. Apache 2.0 licensed.

For example, you can store protected files in a public cloud (like Google Drive, Dropbox, etc.) and safely synchronize them.

#### Install
```
go get -u github.com/nemca/nemstore
```

#### Usage
You must create config file:
```
➜  cat $HOME/.nemstore.yaml
# Set directory to store files
StorageDir: /Users/mbr/Dropbox/nemstorege
```

`nemstore` use `EDITOR` environment variable. Please add it to your shell profile.

Show help:
```
➜  nemstore
The utility allows you to work with encrypted files on the file system.

Usage:
  nemstore [command]

Available Commands:
  cat         Print encrypted file
  create      Create new encrypted file
  edit        Edit encrypted file
  help        Help about any command
  less        Open file in `less` pager
  ls          List of stored files
  rm          Remove encrypted file
  version     Show the version number

Flags:
      --config string   config file (default is $HOME/.nemstore.yaml)
  -h, --help            help for nemstore

Use "nemstore [command] --help" for more information about a command.
```
