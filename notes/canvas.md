### Prompt

"what would this look like rendered on a canvas instead of a css grid?"

### Plan for Canvas-Based Rendering

This document outlines a plan to refactor the Solitaire game's rendering from a DOM-based approach to a Canvas-based one.

#### 1. HTML Update

*   Replace the `<div id="game-board"></div>` in `solitaire/index.html` with a single `<canvas id="game-canvas"></canvas>` element. A default size of 800x600 will be set.

#### 2. New Canvas API Helpers

*   Create a new file, `js/canvas.go`, to provide a Go-friendly wrapper around the HTML5 Canvas 2D API. This will include functions for:
    *   Getting the 2D rendering context.
    *   Drawing shapes (`fillRect`, `strokeRect`).
    *   Drawing text (`fillText`).
    *   Setting colors, fonts, and styles.

#### 3. Complete `render()` Function Rewrite

*   The `render` function in `solitaire.go` will be completely replaced.
*   The new function will, on every game state change:
    a.  Clear the entire canvas.
    b.  Draw the green background.
    c.  Loop through the game state (piles, cards).
    d.  For each game object, calculate its `(x, y)` coordinates and draw it using the canvas API helpers. For example, drawing a face-up Ace of Spades would involve drawing a white rectangle, then drawing the text "A" in the corner and a "♠" symbol.

#### 4. Event Handling Overhaul

*   All event listeners (`click`, `dblclick`, `dragstart`, `drop`) will be removed from the DOM elements (since they won't exist).
*   A single set of event listeners (`click`, `mousedown`, `mousemove`, `mouseup`) will be attached to the `<canvas>` element itself.
*   **Hit Detection:** When a click occurs, a function will be needed to take the mouse coordinates and loop through the known locations of all the cards and piles on the canvas to determine what the user clicked on.
*   **Manual Drag-and-Drop:** Drag-and-drop will be manually implemented by tracking the mouse position from `mousedown` to `mouseup` and re-rendering the dragged card at the mouse's current position on every `mousemove` event.
