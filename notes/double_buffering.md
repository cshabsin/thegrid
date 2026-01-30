That's a very sharp and important question. You are right to be skeptical, as the code doesn't *explicitly* create a second buffer. The "double buffering" in this modern web context is handled *implicitly* by the browser's rendering engine, and it's a key reason why `requestAnimationFrame` is so effective.

Here’s a more precise breakdown of what's happening:

1.  **The Visible Canvas (Front Buffer):** This is what the user actually sees on their screen. It's the final, composed image for the current frame.

2.  **The Drawing Buffer (Back Buffer):** Every `<canvas>` element has an internal, off-screen bitmap where all drawing commands are executed. You can think of this as the "back buffer." When you call `ctx.ClearRect()`, `ctx.LineTo()`, or `ctx.Stroke()`, you are not modifying the visible screen directly. You are modifying this hidden drawing buffer.

3.  **The "Buffer Swap":** The magic happens when your `requestAnimationFrame` callback function finishes executing. The browser knows you're done with all your drawing commands for that frame. It then takes the completed image from the drawing buffer (the back buffer) and paints it onto the screen (the front buffer) in a single, atomic operation.

Here is the sequence for our Lissajous app:

*   **Frame N is visible on screen.**
*   `requestAnimationFrame` fires your `renderFrame` function to prepare Frame N+1.
*   `ctx.ClearRect(...)`: The *hidden* drawing buffer is completely cleared. The user still sees Frame N.
*   `for` loop with `ctx.LineTo(...)`: A new curve is drawn line-by-line onto the *hidden* buffer. The user still sees Frame N.
*   `ctx.Stroke()`: The path is rendered into the *hidden* buffer. The user still sees Frame N.
*   Your `renderFrame` function exits.
*   **The browser now performs the "swap":** It takes the fully rendered image from your hidden drawing buffer and displays it on the screen. The user now sees Frame N+1, which appeared all at once without any visible clearing or partial drawing.

So, you are correct that we aren't manually creating a second canvas. We don't need to. The browser's rendering pipeline for the `<canvas>` element gives us this double-buffering behavior for free, preventing the flickering that would happen if we were drawing directly to the visible screen.