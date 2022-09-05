# Concepts

## Language

## Methods

All methods (and field accesses) will be compiled to global methods.

### Order of Functions

any function `fun x (a b)` where no type information is present,
or any other restriction, will be evaluated last.


## Arrays

operations work like in numpy.

dont forget:
    (get a (> a 1))
    where a is an array
    and (> a 1) is an array of booleans

later: maybe typed arrays make a lot of sense


### Function Overloading
using match

instead of specifying names of arguments, one can simply capture using match arms


e.g.

    (fun git
        ["clone" reponame] (...)
        ["push"] (...))

but also across different types, to create overloadings!

    (fun very-weird
        ["a" n] (...)
        [8] (...)
        [1 y 1] (...)
        ["tell" user] (str "hello " user))

#### Syntax

normal function, without overloading

    (fun square (x) (* x x))

    (fun div-safe (x 0) none
                  (x y) (some (/ x y)))

    (fun git
        ("diff") (...)
        ("clone" reponame))

    (fun locale
        ("de" | 0)  (ok locale.de)
        ("en" | 1)  (ok locale.en)
        (l in ("fr" "sp"))  (err (str "locale " l " not yet supported))
        (x: Int) (err (str "invalid locale id " x))
        (x: String) (err (str "invalid locale " x)))

    (fun option (nil) none
                (x)   (some x))

    (fun sum ([]) 0
             ([x ..rest] (+ x (sum rest))

##### Grammar of argument block

    Block := "(" Arg* ")"
    Arg := Constant | Variable | Array
    Array := "[" Arg* (".." ident)? "]"
    Constant := "nil" | bool | number | const-string
    Variable := ident ((":" Type) | ("in" Binding))?
    Type := ident
    Binding := "(" Constant+ ")"

#### Argument Patterns

An argument can be of type Constant, Variable or Array.
Constants do not count as variables, so that we do not have to bother with them in Bindings

    (fun matches (a: Arg, o: value.Object): Maybe Map(Identifier, value.Object), (...))

    type Arg := Constant | Variable | Array
    type Constant := Nil | Bool | Int | Float | ConstString

Easy to compare in match arms so far!

Variables can be hold additional information.
Either a Binding (e.g. `x in (0 1 2 3)`) to constants, or a type.

    type Variable
    - identifier String
    # one hot encoded
    - type       ClassId?
    - binding    []Constant?

That is fairly easy to compare as well still. Very nice.

Now, what is with pattern matching over arrays?
Arrays can have a number of Arguments themselves.
As such, they can yield a number of variable bindings to values after comparision to them. (Represented as a `Map(Identifier, value.Object)`)
Futhermore an array can hold their remaining arguments.

    type Array
    - args []Arg
    - rest string?

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


### Differences to solar
notes for later


    fun matches (arg, more) Int = arg : splice (first more)

    (fun matches (arg, more): Int (splice arg (first more)))
