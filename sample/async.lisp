
(defn do-something [x: Int] (+ x 8))

(let x (<3 do-something 8)
  (let multi (+ (.. x) (do-something 9))
    (println multi))

; should print 33
; expected output: 33
