### Generic UI Board Refactoring Plan

This document outlines a plan to refactor the UI rendering logic from `solitaire.go` into a generic, reusable `Board` component within the `cardkit` framework. This will make it significantly easier to build UIs for other card games like Spider Solitaire.

### The Goal

The end state is a `solitaire.go` file that is dramatically simplified. It should only be responsible for creating the game-specific logic (`klondike.Klondike`) and passing it to the generic `ui.Board` for rendering.

```go
// Conceptual goal for solitaire.go
package main

import (
    "github.com/cshabsin/thegrid/cardkit/ui"
    "github.com/cshabsin/thegrid/solitaire/klondike"
)

func main() {
    game := klondike.New()
    board := ui.NewBoard(game)
    game.AddListener(board.Render)
    board.Render()
    select {}
}
```

### Analysis of Deficiencies and Proposed Solutions

**1. The Entire `render()` Function**

*   **Problem:** The current `render()` function in `solitaire.go` contains all the logic for displaying the game state, but this logic is almost entirely generic (iterate piles, clear DOM, append card elements). A new game would require duplicating this entire function.
*   **Solution:** Move this rendering loop into a new `Board` component in a `cardkit/ui` package. The `Board.Render()` method will become the single source of truth for displaying any game that conforms to the required interface.

**2. Pile-Specific Styling Logic**

*   **Problem:** Game-specific styling rules (e.g., fanning waste cards horizontally, stacking tableau cards vertically) are currently hard-coded inside the `render()` function.
*   **Solution:** The game logic itself should provide these layout rules to the generic `Board`. We will define a `PileLayout` struct that the `Board` can query for each pile.

    ```go
    // In a new cardkit/ui package
    type PileLayout struct {
        Direction      LayoutDirection // e.g., Horizontal or Vertical
        CardOffset     int             // e.g., 20 or 30 pixels
        MaxVisible     int             // For the waste pile
    }
    ```

    The `klondike.Klondike` struct will implement a method like `GetPileLayout(pileName string) PileLayout`.

**3. The `main()` Function's Setup Boilerplate**

*   **Problem:** The `main` function is full of boilerplate for setting up the UI, such as creating the `GameUI` struct, the pile `div`s, and the `CardToDOM` map.
*   **Solution:** The generic `ui.Board` should handle all of this automatically. The `ui.NewBoard(game)` constructor will inspect the provided game object and create the necessary DOM structure.

### The Refactoring Plan

1.  **Create `cardkit/ui` Package:** Create a new directory for the generic UI components.
2.  **Define Game Interface:** Define a `Game` interface that any playable game (like Klondike or Spider) must implement. This interface will expose methods like `GetAllPiles()`, `GetPileLayout()`, and `AddListener()`.
3.  **Create `Board` Component:** Implement the `ui.Board` struct. Its constructor will build the initial DOM, and its `Render()` method will contain the intelligent, dynamic rendering logic.
4.  **Refactor `klondike.Klondike`:** Make the `Klondike` struct implement the new `Game` interface.
5.  **Refactor `solitaire.go`:** Simplify the `main` function to the target state described above.
