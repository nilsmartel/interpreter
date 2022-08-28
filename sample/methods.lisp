(class Point (x y)
    fun str(self) (str "Point" (x self) " " (y self))
)

(print (str (Point 1 2)))