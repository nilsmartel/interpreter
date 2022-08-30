(class Point (x y)
       fun +(self other)
            (Point (+ (x self) (x other)) (+ (y self) (x self))))

(print (let a (Point 1 2)
(let b (Point 2 3)
  (+ a b)
  )))
