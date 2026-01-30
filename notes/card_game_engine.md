### Card Game Engine Refactoring Plan

This document outlines a plan to refactor the Solitaire codebase to create a generic, reusable card game engine.

#### 1. Create a `pkg/card` Package

*   **Purpose:** To represent the most fundamental building blocks of a card game.
*   **Actions:**
    *   Create a new directory: `pkg/card`.
    *   Move the `Card`, `Suit`, and `Rank` types from `solitaire/game/game.go` into this new package.
    *   Create a `Deck` type in this package with methods like `NewStandard52()`, `Shuffle()`, and `Draw()`.

#### 2. Create a `pkg/pile` Package

*   **Purpose:** To represent a generic, ordered stack of cards.
*   **Actions:**
    *   Create a new directory: `pkg/pile`.
    *   Define a `Pile` type, which will be a wrapper around a slice of `*card.Card`.
    *   Implement universal methods for the `Pile` type, such as `Push(c *card.Card)`, `Pop() *card.Card`, `Peek() *card.Card`, `Len() int`, and `Shuffle()`.

#### 3. Isolate Klondike-Specific Logic

*   **Purpose:** To separate the specific rules of Klondike Solitaire from the core engine.
*   **Actions:**
    *   Rename the `solitaire/game` directory to `solitaire/klondike`.
    *   Rename the `Game` struct within this package to `Klondike`.
    *   Refactor the `Klondike` struct to use the new generic packages. Its fields (`Stock`, `Waste`, `Foundations`, `Tableau`) will become `*pile.Pile` types.
    *   All Klondike-specific rules (e.g., `CanMoveToTableau`, `DrawCards`, `CheckWin`) will remain in the `klondike` package.

#### 4. Update the UI Layer

*   **Purpose:** To connect the user interface to the newly refactored game logic.
*   **Actions:**
    *   Update `solitaire/solitaire.go` to import and use the new `solitaire/klondike` package.
    *   The UI will now interact with the `klondike.Klondike` struct instead of the old `game.Game` struct.

#### 5. Update Build Files

*   **Purpose:** To ensure the project still builds correctly after the refactoring.
*   **Actions:**
    *   After all files are moved and refactored, run Gazelle to automatically update all `BUILD.bazel` files.
    *   This will create new build targets for the `pkg/card` and `pkg/pile` packages and update the dependency graphs across the project.
