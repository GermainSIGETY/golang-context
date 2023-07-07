# About Contexts in Golang 
> What is it and how to use it

Please Read medium article for a complete course : 
**[https://medium.com/@gsigety/about-contexts-in-golang](https://medium.com/@gsigety/about-contexts-in-golang-b8b4b56ca3c8)**

## What is it ?

Context is struct with method built for Two usages / objectives :

### 1. store additional information related to a request, event, message to handle

Context in Golang is an object that can be used to store information or objects related to processing a request, a message or an event.
It is not a generic way to pass information from a function to another.

Example of additional information :
- information about an authenticated user that requests something
- a logger that prepends messages with a request ID
- a single database transaction for each web request using your database connection

Typical use case is development middlewares for http server librairies : a middleware add values to request's context, when request enters, 
then use this value just before returning the response ; compute and log request's response time or other stuff

### 2. A cancellation system for requests processing 

In some case treatment of requested can be canceled, abandoned if client has already closed its connection.
Why continuing a request to a database, a search engine or third party service for a client that has already left ?

This use case is very practical for an http server, once a client left we could stop everything.

There are different ways of cancellation. let's see them in simple cases, before presenting a more real use case.

## Examples

### [1_USE_CONTEXT_VALUE/main.go](1_USE_CONTEXT_VALUE/main.go)

Usage of a context can carry additional values.

### [2_USE_CONTEXT_WITH_CANCEL/main.go](2_USE_CONTEXT_WITH_CANCEL/main.go)

Usage of a Context that is cancelled.

### [3_USE_CONTEXT_WITH_CANCEL_CAUSE/main.go](3_USE_CONTEXT_WITH_CANCEL_CAUSE/main.go)

Usage of a Context that is cancelled with a cause.

### [4_USE_CONTEXT_WITH_TIMEOUT/main.go](4_USE_CONTEXT_WITH_TIMEOUT/main.go)

Usage of a Context that timeouts.

### [5_USE_CONTEXT_WITH_DEADLINE/main.go](5_USE_CONTEXT_WITH_DEADLINE/main.go)

Usage of a Context with a deadline.

### [6_TANGIBLE_USAGE_WITH_HTTP_SERVER](6_TANGIBLE_USAGE_WITH_HTTP_SERVER/main.go)

A complete example of usage of context with a gin server : 
When a client close its connection, application/serve can detect it thanks to context, and stop doing stuff for a client that already left.

### [7_CONTEXT_HIERARCHY](7_CONTEXT_HIERARCHY/main.go)

Example of hierarchy of contexts.

### Conclusion

Once again we have seen Golang has good design for high traffic management ; language provide tools to optimize leftover clients

