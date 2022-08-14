# plotter-ui

`plotter-ui` is a tool to visualize Fybrik plotter.

The UI is built with [elyra-canvas](https://github.com/elyra-ai/canvas/wiki), an open source tool developed in IBM.

For more info and help, follow `elyra canvas` slack channel at: @elyra-canvas


![A plotter-ui example](docs/plotter-demo.png?raw=true "Title")


## Setup:

### Requirements:
+ `go` version 1.16+

+ `node` version 14.X.X


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

On the first use you'll need to install all the required packages.

In main dir run this commmand:
```
npm install
```
This command will take a while to finish and will create `node_modules/` dir.


## Run application locally:

After you finish installing all the packages, run the application with this command on the main directory (takes several minutes):
```
npm start
```


# Create a pipeline:

In generate-pipline:
```
go run main.go <plotter-path>
```

You can use some of the example plotters in `examples/`