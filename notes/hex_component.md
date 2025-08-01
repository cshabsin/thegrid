# Refactoring HexGrid to Individual Components

This note captures a conversation about refactoring the `HexGrid` from a monolithic SVG path to a grid of individually manipulable `Hex` objects.

---

### Prompt 1: Making Hexes Individual

> What would it look like to make the HexGrid be comprised of indivdual Hexes that can be manipulated individually? I'd like to do something like set the fill on individual hexes, for example.

### Answer 1: Refactoring to a Data Model of `Hex` Objects

That's a fantastic question. It gets to the core of making the hex map truly interactive and dynamic. To achieve what you're asking for, we need to change the fundamental data model of the `HexGrid`.

Here’s what the refactoring would look like:

1.  **Introduce a `Hex` Object:** Right now, the `HexGrid` is just a mathematical concept that we use to draw a single, monolithic SVG `<path>` for the grid mesh. I'll introduce a new `Hex` struct. Each instance of this struct will represent a single hexagon on the map and will be able to hold its own properties, such as its fill color, an ID, or a CSS class.

2.  **Refactor `HexGrid`:** The `HexGrid` will be refactored from a simple calculator into a container object that holds a 2D slice of these new `Hex` objects. It will effectively become a `[][]Hex`.

3.  **Individual Hex Rendering:** I'll remove the current `GridMesh` method, which draws the whole grid at once. Instead, I'll add a `ToElement()` method to the new `Hex` struct. This will allow us to generate an individual SVG element for each hex. We can then iterate through the `HexGrid` and render each hex one by one, applying its specific properties.

4.  **Update the `explorers` App:** The `explorers` app currently draws a generic hexagon in the `NewParsec` function. I'll update it to:
    *   Fetch the specific `Hex` object from the `HexGrid` for the given coordinates.
    *   Use the `Hex.ToElement()` method to render it.
    *   This will be the point where you could set the fill. For example, before calling `ToElement()`, you could do `hex.Fill = "url(#starfield)"` or `hex.Fill = "blue"`.

This is a significant but powerful change. It moves us from a static, one-piece grid to a dynamic grid composed of individual, addressable objects. This will make it much easier to implement features like highlighting a single hex, changing its color on the fly, or applying different patterns to different hexes.

---

### Prompt 2: The Doubled Border Problem

> I worry that in an earlier version of this, neighboring hexes had a "doubled" border where both hexes drew lines - the single SVG path was the workaround I came up for back then, but maybe there's something better to do?

### Answer 2: The Two-Layer Rendering Solution

That is an excellent and critical point. You've correctly identified the primary drawback of rendering individual hexes. The "doubled border" problem is a classic issue in grid rendering, and your previous solution of a single SVG path is the standard, correct way to avoid it.

However, that approach is fundamentally incompatible with styling individual hexes. We are at a crossroads with two competing goals:

1.  **Perfect, single-thickness borders.** (Favors the single path approach).
2.  **Individually styleable hexes.** (Requires individual hex objects).

We can, in fact, have the best of both worlds. The standard solution for this exact problem is to use a **two-layer approach**:

1.  **The Fill Layer:** We will create a `[][]Hex` data model as discussed. We will iterate through it and draw an individual SVG hexagon for *each* hex. However, these hexes will be drawn **with no border** (`stroke="none"`). Their only job is to provide the fill (color, pattern, etc.). This gives us our individual control.

2.  **The Border Layer:** After drawing all the fills, we will draw the original, monolithic grid mesh path **on top** of the fill layer. This single path acts as a "wireframe" overlay, providing the clean, crisp, single-thickness borders for the entire grid.

This way, we get the visual perfection of the single-path border while still having a data model of individual hexes that we can color and manipulate underneath it.
