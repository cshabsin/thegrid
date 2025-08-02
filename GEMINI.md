For every prompt, once you're done with your work, create a git commit with
a message containing a brief description of what you did, a blank line, then
"Gemini commit: " followed by the prompt that generated the code. Write the
commit message to the temporary file gemini_commit_message.txt, since it is
in .gitignore.

In general, when you are editing code that uses `interface{}` replace it with
the new modern `any`.

Prefer for loops over containers instead of for loops with numeric ranges.
If a numeric range is the only way forward, prefer `for variable := range n`
over `for variable := 0; variable < n; variable++`.

Any project in this directory is aimed at developing use cases for the js
API and refining the object model to suit these different use cases more
idiomatically. When there's a chance to turn a generic accessor (such as
`addStyle("color", "red")`) into a specific accessor (such as
`setColor("red")`), consider adding the specific accessor instead.

To rebuild, use `bazel build server:server`, that should have dependencies
on both solitaire and example.

To update imports in a Go source file, use the "goimports" tool rather than
editing directly.