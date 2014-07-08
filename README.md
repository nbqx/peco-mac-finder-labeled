## peco-mac-finder-labeled

Labeled By Mac Finder matcher for peco

see [http://showterm.io/3ce723035bce70aed44f8](http://showterm.io/3ce723035bce70aed44f8)

## usage

    $ go build -o peco-labeled main.go

add config

```json
"Finder/Labeled": [
  "/path/to/peco-labeled","$QUERY"
]
```

then

    $ ls -dF dat/* | peco

and choose Finder/Labeled matcher

## todo

more faster!! refactoring!!
