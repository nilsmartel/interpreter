# interpreter

this is a simple scripting language designed to teach myself a little bit about interpreters and because it's fun.

Ultimately I am interested about creating a feature set, that would make this language a nice fit to replace bash.
Though that is a non priority, as this is a fun project to work on, more than anything else 


On first glance it seems very lisp-y, and in some regards it is.

This code for example will be run just fine

```lisp
(print (if true (str "hello "  "world " 78) "this should never be here"))
```
