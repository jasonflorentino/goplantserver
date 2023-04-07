# Plant Server GO Edition

**_Teaching myself [GO](https://go.dev/) by porting my little [Plant Server](https://github.com/jasonflorentino/plant-server) application to golang_**

This year has already been slammed, so we're learning and building bit by bit...

First thoughts so far are that golang is pretty great! Still getting stuck on this module system, but coming from JavaScript I guess I'm used to the trouble. And I'm already feeling worried I'll get annoyed writing `if err !== nil` all the time. That said, what little code I've written is already starting to make me not want to go back to writing JS/TS. 😅

— Jason, March 2023

# Usage

As the name might suggest, you'll need to be able to run or build [Go](https://go.dev/) code on your computer.

Some routes require an API key. Make sure you set it in your environment before running the app:

```bash
export PLANTSERVER_API_KEY=mysecretkey
```

Protected routes are accessed with an `api_key` query parameter on the request URL:

```bash
http://localhost:8080/plants?api_key=mysecretkey
```

Run the app from your shell:

```bash
go run main.go
```

Or build and run an executable:

```bash
go build main.go; ./main
```

The app will be running at http://localhost:8080/
