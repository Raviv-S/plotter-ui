# Plotter-ui Examples

In order to test the examples in the ui, you should change the `src/App.js` accordingly:

On the import list:
```
import pipeline from "./generated-pipeline.json"; \\-> Change to "../examples/plotter-#-pipeline.json"
import paramDef from "./generated-params.json"; \\-> Change to "../examples/plotter-#-params.json"
```

In order to run a demo plotter, run this in `generated-pipeline` dir:
```
go run main.go ../examples/plotter-#.yaml
```


The plotters examples are scenarios taken from:

https://docs.google.com/document/d/1vsiFiVvc2Al8dIN5TopHZPd43-fkpCFT1qgS97hqSvA/edit?usp=sharing

