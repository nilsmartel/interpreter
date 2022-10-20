(class Point (x y)
       (fun + [self other]
            (Point (+ (x self) (x other)) (+ (y self) (x self)))))

(class BufferedWriter
       (_vec w)
       (fun write [self data]
            (do
                (push (_vec self) data)
                (if (> (len _vec) 1024) (do
                                          (write (w self) (_vec self))
                                          (assign (_vec self) [])
                                          true) false))

(print (let a (Point 1 2)
(let b (Point 2 3)
  (+ a b)
  )))
