# deep-throtle-parser

## Vision:
The parser is able to grab web pages, and pull all links from them -> then
follow these links and pull other links from them, and so on, this is necessary to analyze the content used by the site for its further processing.

### Usage:

1. Clone this repo
2. next run this command: 
```bash
go run main.go ${url} -d ${parsing depth}
```

#### For example:

```bash
go run main.go https://github.com/kitt3911 -d 8
```

or

```bash
go run main.go -d 8 https://github.com/kitt3911
```

### In the plans: 

- [x] Deep multithreaded parsing of links
- [ ] Provide users with exactly the builded console utility for all OS
- [ ] The ability to download and translate content

### Authors:

* [Roman Morozkin](https://github.com/MartIden)
* [Ivan Ignatenko](https://github.com/kitt3911)
