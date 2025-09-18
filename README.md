# Go WebAssembly Keystroke Interceptor ⌨️

This tool was initially designed for video engineers working on a film/TV production to eliminate the pressure on actors to hit each key perfectly on scenes requiring them to use a keyboard. The project is a simple web application demonstrating how to use Go compiled to WebAssembly (Wasm) to manipulate the DOM in a web browser. It features a text box that intercepts all keyboard input and types out a predetermined phrase, one character per key press. This can be placed within any existing UI. 

The application is served by a minimal, self-contained web server also written in Go.

## Features

-   **Keystroke Interception:** Captures `keydown` events on a text area.
-   **DOM Manipulation:** Modifies the text area's value from Go.
-   **Go Wasm:** All client-side logic is written in Go.
-   **Go HTTP Server:** A simple file server to host the application.

***

## Getting Started

Follow these instructions to get the project running on your local machine.

### Prerequisites

You must have the Go compiler installed on your system. You can download it from the [official Go website](https://go.dev/dl/).

### Installation & Running

1.  **Clone the Repository**
    ```sh
    git clone [https://github.com/amillerrr/wasm-typer.git](https://github.com/amillerrr/wasm-typer.git)
    cd wasm-typer
    ```

2.  **Copy the Wasm JavaScript File**
    Go requires a JavaScript "glue" file to run Wasm in the browser. Copy it from your Go installation into the project directory:
    ```sh
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
    ```
    NOTE: Some installations will use $(go env GOROOT)/lib/wasm/wasm_exec.js

3.  **Compile the WebAssembly Module**
    This command compiles the `main.go` file into a `.wasm` binary that the browser can execute:
    ```sh
    GOOS=js GOARCH=wasm go build -o main.wasm main.go
    ```

4.  **Run the Go Web Server**
    This command starts the local server to host your `index.html`, `.js`, and `.wasm` files:
    ```sh
    go run server.go
    ```
    You should see a confirmation message: `Starting server at http://localhost:8080`.

5.  **View in Browser**
    Open your favorite web browser and navigate to:
    **[http://localhost:8080](http://localhost:8080)**

    Now, click inside the text box and press any key to see the magic happen!

***

## Customization

You can easily customize the application's behavior.

### Changing the Typed Phrase

To change the phrase that gets typed out, simply edit the `targetPhrase` constant at the top of the **`main.go`** file:

```go
// in main.go
const targetPhrase = "Your new custom phrase goes here!"
```
After changing the phrase, you must re-compile the Wasm module (Step 3) and restart the server (Step 4).

### Optional: Using a Custom Domain (e.g., `tester.link`)

If you want to access your local server using a custom domain instead of `localhost:8080`, follow these steps.

1.  **Edit Your `hosts` File**
    You need administrator/root privileges for this step. Add the following line to your system's `hosts` file to tell your computer that `tester.link` points to your local machine.
    -   **macOS/Linux:** `/etc/hosts`
    -   **Windows:** `C:\Windows\System32\drivers\etc\hosts`

    ```
    127.0.0.1   tester.link
    ```

2.  **Modify the Server Port**
    Web browsers use port 80 for standard HTTP traffic by default. Change the port number in **`server.go`** from `:8080` to `:80`:
    ```go
    // in server.go
    log.Fatal(http.ListenAndServe(":80", nil))
    ```

3.  **Run the Server with Privileges**
    On macOS and Linux, you must use `sudo` to run a server on a port below 1024.
    ```sh
    sudo go run server.go
    ```

4.  **Access the New Domain**
    You can now access your application by visiting **[http://tester.link](http://tester.link)** in your browser.

***

## Project File Structure

```
.
├── main.go            # Go source code for the WebAssembly module (client-side)
├── server.go          # Go source code for the local HTTP server
├── index.html         # The main HTML file that loads the Wasm app
├── wasm_exec.js       # The JS glue file required by Go Wasm
├── main.wasm          # The compiled Wasm binary (ignored by git)
└── README.md          # This file
```
