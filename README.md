# Piet Mondrian like pictures in Go

## Usage

simple examples with lines and rectangles

```bash
$ go run main.go
```

to start a web app

```bash
$ cd app
$ go build
$ ./app
```

## Development

Build docker on a local machine

```bash
$ DOCKER_BUILDKIT=0 docker build --tag 'mondrian:latest' --no-cache .
$ docker run --publish 8080:8080 mondrian:latest
```



## TODOs

- @Vladimir: Make frontend display the generated Mondrian image on the same page where the input form is, instead of redirecting to new page.
- @Vladimir: Limit the complexity input on the frontend and adjust description on first page for the user:
  - for style="with lines" limit to (0.001, 0.05)
  - for style="without lines" limit to (0.001, 0.08)
- @Vladimir: Deploy app to Azure (or somewhere else?)
- @Vladimir: Look into masks if there is enough time.
- @Irena: Write Algorithm Description.
- @Irena: Find a solution of how to make the code executable from all machines (go binary? Docker?)

- Generate our favorite image and print in A3 format and when generating the image, adjust the pixels to look nice on A3 (A3 is 3508 pixels x 4961 pixels).
    We could do a few test prints on A3 to see which Mondrian canvas size should we pick.

### Mask examples

- [many examples](https://golang.hotexamples.com/examples/image.draw/-/DrawMask/golang-drawmask-function-examples.html)
- [masks blogpost](https://medium.com/@damithadayananda/image-processing-with-golang-8f20d2d243a2)
- [image blur repo](https://github.com/brunocramos/go-image-blur/blob/master/image-blur.go)


## Inspiration references

- [Mondrian Process](https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.564.8410&rep=rep1&type=pdf)
- [GoMondrian](https://github.com/8lall0/GoMondrian)
- [generativeart](https://github.com/jdxyw/generativeart)

# Authors

- [pythonmonty](https://github.com/pythonmonty)
- [vvkorz](https://github.com/vvkorz)