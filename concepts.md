# Concepts

## Language

### Function Overloading
using match

instead of specifying names of arguments, one can simply capture using match arms


e.g.

    (fun git
        ["clone" repo] (...)
        ["push"] (...))

but also across domains!

    (do
        ["a" n] (...)
        [8] (...)
        [1 y 1] (...)
        ["tell" user] (str "hello " user))

## Repl

- undo / redo

- realtime syntax highlighting using token tree
- RAINBOW SYNTAX!!

- shell input window is build like:
    |------------------------|
    |   <input>              |
    |------------------------|
    |                        |
    |   <output> ...         |
  - the input window can grow, to show all lines, currently edited
  - if no input is entered (e.g. input : trim : empty) hide input window

- typing <enter>:
    - if input : tokens : len : == 1
        user probably wants to display a variable
        so evaluate expression and print it (Eval(expr).Str())
    - else
        probably not a full expression
        enter new line to input window.

- typing `)`
    - if input does not start with ')'
        wrap input (e.g. `...input` => `(...input)`)
    - if open parens > closed brackets
        => insert bracket accordingly / intellignetly
            if latestopen == '(' => ')'
                              [  =>  ]
                              {  =>  }


- when typing `;` transform
    - `(x...;...y` => (do x... ) ...y` (y is likely the empty-string)
    - `x...;...y` => (do x... ) ...y`

- stout of commands:
    - indented 2 spaces
    - first line of output starts with >
    - clicking on `>` will hide output, transform `>` into `o`
    - clicking on `o` will show output, transform `o` into `>`
    - clicking on output will open dialog to save output into variable

- every shell command: output to stdout AND return stdout as a string
    - the string is `stdout :filter isprintable : tostring` to enable proper color output while keeping strings sane
    -

- when executing shell commands:
    - current global and local variables are made into the environment variables
    - classes are json'ized maybe ?

- when opening shell
    - current envs are stored into variables
    - if jsonable -> json them maybe? lets see how this goes
