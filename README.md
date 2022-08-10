# plotter-ui

`plotter-ui` is a tool to visualize Fybrik plotter.

The UI is built with [elyra-canvas](https://github.com/elyra-ai/canvas/wiki).

## Setup:

### Requirements:
`go` version 1.16+

`node` version 14.X.X


### Setup required:

Install `nvm`:

```
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash
```

Install `node v14.20.0`:

```
nvm install --lts
nvm install v14.20.0
nvm alias default 14.20.0
nvm use 14.20.0
```

## Build package:

In main repo dir:
```
npm install
```

## Run server locally:

In main repo dir:
```
npm start
```

# Create a pipeline:

In generate-pipline:
```
go run main.go <plotter-path>
```
