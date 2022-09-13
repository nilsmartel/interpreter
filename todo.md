# TODO

## Method Calls

- Pattern Matching types ( implemented )

- Parser for functions
- execution.Eval
- Parser for methods

- Rewrite Function.Eval (get rid of settings args -> vars manually. Use Pattern Matching interface instead)

### Parser for functions

    (fun xyz (<args>) <expr>)

    args := <arg> (..<ident>)?
    arg := <constant> | <variable> | <array>
    variable := <ident> ("in" (<args>) | ":" <ident>)?
    constnt := <string-const> | <number> | <bool>



