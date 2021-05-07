# Basic golang App :airplane:

This basic golang application deploy a basic http server using [Gorilla Mux](https://github.com/gorilla/mux.git)

## HTTP Server API :book:

+ **/** : Home page
+ **/user/{name}** : return the value of the name variable
+ **/magic-number/{num}** : generate a magic number using num var

## Build :hammer:

### Go build

    cd app 
    go build

If you want to rename the build, run insteand

    go build -o <name>

### Docker image

    docker build -t <image_name> .
