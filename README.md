# Example Go lambda function

Turns out Lambda functions in Go are pretty elegant. Here's an example "REST API" style lambda function -

POST JSON - and it shall return JSON! ~ Json Bourne

To build:

```
GOOS=linux go build .
zip -r example.zip Go-Lambda-example
```

Now upload example.zip to lambda, under 'Handler' put the name of the binary (Go-Lambda-example).

test.json contains an example test event which will return

```
{
  "greeting": [
    "Whats up Fella",
    "Whats up G",
    "Whats up Fellow Gopher",
    "Whats up Dude!"
  ]
}
```