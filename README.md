# Playing around with WASM

Kidnapping this repository to just play around with WASM for a bit. Notes:

In the "example" directory, run:

```
GOOS=js GOARCH=wasm go build -o main.wasm
```

Then from the top-level, run

```
go run server/server.go --dir example
```

# thegrid

(Maybe someday...)

The Grid is a video game concept playground I'm messing with. I have no idea
what it will become over time, but for now I'm just... messing around.

The idea I'm starting with is basically, a Game of Life grid, but... with some
kind of richness. Each cell in the grid obeys its own set of rules, so yes,
you have some "Conway" cells (possibly the default cell / background noise?)
but you also have other cell types.

A simple cell type is a "source" cell, which is just always true. A "sink" cell could always be false. (What if the life in a cell has a color?) A pulsar cell
could toggle on and off (or generally provide a stream of specific values in a
loop.)

Other cell types could be conveyors. A "west" conveyor would simply always take
the value of the cell to its right ("moving" that value west). One for each of
the four directions.

From here... can I move things around enough to make a tower defense game? A
manufactoria type game? I don't know.

