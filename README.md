# Playing around with WASM

Kidnapping this repository to just play around with WASM for a bit. Notes:

To compile "example", run:

```
GOOS=js GOARCH=wasm go build -o example/main.wasm github.com/cshabsin/thegrid/example
```

Then from the top-level, run

```
go run example/server/server.go --dir example
```

(You can just keep this server running and recompile... might need to
shift-reload)

Part of the setup may involve updating the wasm_exec stub that Go provides:

```
    cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" example/wasm_exec.js
```

# Bazel notes

To add a new external dependency, e.g. secretmanager:

Add the import to a Go file somewhere in the directory tree. Run 
`go mod tidy`, then run `bazel mod tidy`. Then `bazel run //:gazelle`.
That seems to have done the trick, anyway

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

# Firebase Configuration

The server fetches the Firebase configuration from Google Secret Manager. You will need to create a secret named `firebase-config` in your Google Cloud project.

The value of the secret should be a JSON object with the following structure:

```json
{
  "firebase": {
    "apiKey": "YOUR_API_KEY",
    "authDomain": "YOUR_AUTH_DOMAIN",
    "projectId": "YOUR_PROJECT_ID",
    "storageBucket": "YOUR_STORAGE_BUCKET",
    "messagingSenderId": "YOUR_MESSAGING_SENDER_ID",
    "appId": "YOUR_APP_ID"
  }
}
```

Here is how to find the values for each field in the [Firebase console](https://console.firebase.google.com/):

*   **`apiKey`**: In the Firebase console, go to **Project settings** > **General**. The API key is listed under **Your apps**.
*   **`authDomain`**: In the Firebase console, go to **Authentication** > **Sign-in method**. The auth domain is listed at the top of the page.
*   **`projectId`**: In the Firebase console, go to **Project settings** > **General**. The project ID is listed under **Your project**.
*   **`storageBucket`**: In the Firebase console, go to **Storage**. The storage bucket URL is listed at the top of the page. You need to remove the `gs://` prefix.
*   **`messagingSenderId`**: In the Firebase console, go to **Project settings** > **Cloud Messaging**. The sender ID is listed under **Project credentials**.
*   **`appId`**: In the Firebase console, go to **Project settings** > **General**. The app ID is listed under **Your apps**.