# classes

classes are defined like this

```lisp
(class ClassName (field1 field2 field3)
    fun method1(self) <expr>
    fun method2(self) <expr>
)
```

concretly that means

```lisp
(class Point (x y)
    fun +(self other) (Point 
                        (+ (x self) (x other))
                        (+ (y self) (y other)))
)
```

fields can be accessed like this

```lisp
(let p (Point x y)
    (print (x p)))
```

calling methods can be done like this

```lisp
(let p (Point x y)
    (let p2 (Point x y)
        (print
            (+ p p2)
        )))
```
