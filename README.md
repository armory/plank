![](https://cl.ly/1m341B1l0l2P/plank_logo-final.png)

# plank

Spinnaker SDK for Go.

*Plank is a work in progress and will change drastically over the next few months.*

## What is it?
A package used by services that interact with Spinnaker's micro-services. It is not intended to be a client which interacts with Spinnaker's outward facing API.

## Why is it named plank?
Because it's funny.

## How do I use it?
Very carefully. :smiley:

Basic concept is that you instantiate a Plank client thusly:

```go
client := plank.New(nil)
// NOTE: nil defaults to http.Client, but you can sub in anything compatible
//       you'd rather use.
```

You can (or may need to) replace the base URLs for the microservices by
assign them to the keys in the `client.URLs` map:

```go
client.URLs["orca"] = "http://my-orca:8083"
client.URLs["front50"] = config.Front50.BaseURL
```

After that, you just use the Plank functions to "do stuff":

```go
app, err := client.GetApplication("myappname")
pipelines, err := client.GetPipelines(app.Name)
// etc...
```

## Testing

The `mock_plank` directory is there for you to easily mock out the Plank
client when testing your own code.  This was built using the "gomock"
package and can be used like this:

```go
package yourapp

import (
  "testing"
  "github.com/golang/mock/gomock"

  "github.com/armory/plank" // Included for interface, structs
  . "github.com/armory/plank/mock_plank" // The mock client
)

func TestThingUsingPlank(t *testing.T) {
  ctrl := gomock.NewController(t)
  defer ctrl.Finish()

  client := NewMockPlankClient(ctrl)

  // Now, you can set the expected call, the desired response, and check
  // how many times it's called.
  client.EXPECT().
    GetPipelines(gomock.Eq("foo")).
    Return([]plank.Pipeline{}, nil).
    Times(1)

  // Now use the mock client in your code
  err := YourFunction(client)
  // Test results like you'd expect

  // At the end of the test, gomock ensures that that GetPipelines() call
  // on the client was called *only* once, and was called with the argument
  // "foo" (which returned an empty list of pipelines and a nil error).
  // If you client called the function with a different argument, called it
  // more than once (or never called it at all), the test will fail.
}
```


