# Ed's adapted fly.io and go

A minimal Go application for [fly.io Getting Started](https://fly.io/docs/getting-started/golang/) documentation and tutorials.

To get started:

1. clone this repo
2. `flyctl launch` to launch the app on fly.io for the first time
3. `flyctl deploy` to deploy the app to fly.io after the first launch
4. `flyctl open` to open the app in your browser
5. `flyctl logs` to see the logs
6. `flyctl rsync` to sync files to the fly.io app
7. view the deployed app with flyctl open


In one terminal, run this to start the server:
```
air app.go
```
Air will do live-reload of the Go server on file changes

In another terminal (split terminal), run this to have automatic live reload in the browser:
```
browser-sync start --proxy "127.0.0.1:8080" --files "templates/*.tmpl" --reloadDelay 2000 --no-open --no-notify
```

The reloadDelay is to avoid browser-sync from reloading the page too quickly while air is rebuilding the Go app.

You don't have to do the browser-sync bit if you're happy to manually refresh the page.s