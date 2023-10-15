# GoMDP

GoMDP is a Markdown previewer written in Go.

## Features

* Support for links [MySite](https://joshuaejs.me)
* Support for other features

## How to install

```txt
go install github.com/joshuaejs/mdp
```

## What it does

1. Read the contents of the input Markdown file.
2. Use some Go external libraries to parse Markdown and generate a valid HTML block.
3. Wrap the results with and HTML header and footer.
4. Save the buffer to an HTML file that can be viewed in a browser.
