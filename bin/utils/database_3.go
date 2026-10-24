package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.9" )

func K0HwUOw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUugUFqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itv6PueFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MZHcc9rdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paN8dS7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxlfpsJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLo77CyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vK3CRpENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8U4s9PXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zauGJIEwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLtKI6ySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P76aCSSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rY1vSuU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwqRTWntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Twn9Q084Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cx9K9MBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0n9vvEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEcsaELfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c856qclaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzmOGgrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hbcgw7BcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGoJX33rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYZhr2cpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQj3rdMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tdr9WKhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30H47v6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzbi1WTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oH146uHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OrikX1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ckKAWGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ly7wIF5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMEECxvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CIA6nJoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpg3UhCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CKYtnvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzSw4bSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqUN0vIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9brSmuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyC6wTFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pP7cpISgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zrtgh3ztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ohXZnsj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9rQniKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdwO7hiGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGnhiscMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFhtY1WuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90ptv1xXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sexL1GoOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xeuFPTyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbG7vbkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WoiueZaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59w1CARLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qaUudAAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l0laxffvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2ajgfx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBPo4IjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIxDKOpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func laxtrPwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIsgJ5pSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQ53SfGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhMBi1zyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMsGzeliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhdbQsfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V7rYyucIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 301E3zpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OjoRVjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZLJiNaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIueuZ1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5G1lEzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RW2Eb4sTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASe2OotbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIJLxowAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5ODRAmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8E5QRnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngSKLC4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uI3jGl68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUrbiZU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CsAvz1uaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZgoc1ebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xfarhsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JunQnoEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9sVrPfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYAhpRGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R76uAprbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9UlhlmbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BixIIQuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VOGkkUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFNfmiTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSyNRqeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FiZOVjmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1cvSsXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTmikDajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJ3bQaUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRRCoqBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy8ev8zAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3Yhqwg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stU3ySrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQc05ZjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haNNVu2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXXFBdlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OuZl1TjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJYmoIy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZOWaaDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJ6jJaRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msNiyMQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhqgdZxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DK4UmQI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 330rdSgnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSGcEqF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxgVq3qVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22POfpGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89guX09BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ME0VlSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func el9GcBpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xH6dOHS9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Ebwg3TSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdzpqBrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ViWxb2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thP82FQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZIN4k1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXZYdLuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iCLsvrkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSOYIrICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGbrdJKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bmxhRgcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tmf2GJNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9NaAFGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12NQhZUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8OMmpLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSIXISoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6o5M7UIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYgN1vwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOCzPDN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vqn1a8BaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPPOPlMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CpgvZjsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCubbufEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFIxtnfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O84v64SUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YePvvxcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqNw1ZWOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNfjVHguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yR1laRxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sF2Zp9bxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D34OacCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8tz4PS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EwdGte1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXvspcvPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQdLug94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTqUMBAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDQfVg8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDtXNas7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBcpQ01qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjl1KP60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5BBmIlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PusS2ZxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N44Do6FSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6qpTiQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLakfVWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqlVpsRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wg2qSnxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CF2v6r3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RerwOXcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFjciUmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8on6mQnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHXgEKdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s7FT93LfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YsmfGdx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZp5R06SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jo0lqKu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKFr85OoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q05TgkRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g84Jet1RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jifOoa5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8oJY1e0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpR6G9JbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVaNSgVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipPe52GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMxlyGFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yQ9Va7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1218lg0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdKrtwGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxMk1pcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8FXlJ9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPRo1BVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAkhudJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5Bl2MKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FRNjhJ1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmWfqvSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PD6X6IITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoAbqCPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QulRAGViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aglG9CEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVGMvq5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vj76gy1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTvsk7rmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5mr46WMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T70KnbUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FQDKCCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMrNH6PnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mgy4VK5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmNpNJvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUnamEBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2fhpZKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5591V2JoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3kaAbw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ev7BZBg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZPSnkRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhE0Dls2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uN7qO8iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhPCsUdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfK9IkYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ibzAZWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ee1OHuwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4toItFdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYRqNBWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fjDuEr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTmI3KCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 457voy9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tkO7MbUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v6EyfLSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R24Vay1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func twR3Em4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dionMt2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EV3Gr3E5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEC6HZiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rShvUdnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsUDzS37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAuFJXiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0bdSr2SDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upt6FThTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbq5rJQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgobmXUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFBtoigjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yyY84bKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxOzKqAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5UmSR67LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXd9PI9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4BOeOa5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5E9tx6MgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqH4PvPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VboyZl9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8zpadNcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2yuYrpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2OvyUFV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIeO460cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vSqzDvMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLrfgSkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5RsEFjVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3ZxIAQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwqHLmsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IygKh9YrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGk8PGPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeCSqoShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R9YZaRHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfTSz3z3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYRn1FvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjdPxoPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4X655j9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnekXr4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kD06HB6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtGGY5WwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCXPMp2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZyqbglaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2AfB9o3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiN6hdhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxTPgDiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8AGCfvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcMxR62JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTbWiY8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJHCw3P8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ke71q6PQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rL0JQgugWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzZfNjG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WimU51PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfxAYt0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEvWwLWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DjC49OvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6zXKLnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLUcN2A8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4M3Xck5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Q25Ey6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnEiiTMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZTam7XiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dxtne9RIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxyubUo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SzCKQnF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJbs8YSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQCzw1z5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCgrkLpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYhijZvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eEcxAaFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrJX4mGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUFXxzE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4t5BbOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AH2wRfiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I06Iho6HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4jxEMEBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func goF1VHqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tStCmbr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WiRnSQKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20ihsFBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoDuPVEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKiZYZ34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5oQGgxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isVoHSANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QfV6PdxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztlbr1hLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12rCCcAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffhNyvSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8aGeGrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAKGKGybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lRYLZFzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vg2Ema5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZcbEpetFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22HkMwe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func surlnZr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXEKbo4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1movvvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rk0HpN8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKIo7W5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BM0EtJ5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAKgsNAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPe968iQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsjSihANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7cMaXxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqqEWBm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPmMj6AYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afNU138TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFMhlqunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sr9IzAKSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Y5kwV4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTOcJvqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bS0dWoDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADvPKilzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COhXZqLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nI3R5fy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgjlI8dxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bZJ7nnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4vR8tmIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P6AU8OjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0ikvuyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idOKDcUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNlHIQTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Db2Iamh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhWDpPCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jddzdpg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIsCbnsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIYs6GdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpmMeIzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vtFdrwSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxRNsoXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xnpd8UOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24aTA8LAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93lWiAjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ascDUHwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sZm4tixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O2IBJyOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Xwokl3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZewMoRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unlm9LIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 65HWg55BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SErmxgJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bzq0ndTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkhbxXIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8D9d6C4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfVNjuXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSTLDMlXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0OZcpSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EEFOUfXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 906ADYGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbBvXpRyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwVxKw1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1UWnotDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ye5Bha6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlaryUaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZKVgxEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlAaYmy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HyehBng3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPKrr0PCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gEEIeR2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qDxS4osBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aTyqMAfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzwu9gngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxyymYOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0NeYLYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUb8avGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Km4nSWLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BG8iptUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RmBO0cyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XB3FghOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2vjk7UTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8O3fcOHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTmNYVIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tww6bvyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZdmchRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtCW1iaaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ir7eP4QCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMhrDoEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zS1Ihc2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bhjF3PDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQs79feWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nx74YO0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOd1dz0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1V8vWbM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oka6GzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDjccyX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aHd6a1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jcnxoafqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vk5YtA5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49QIrdDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6gieUrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56L3dfL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVSuD91tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f41dbmwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLDYkJtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39zu92oNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n5qRJljSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYiYIGCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sl5kCLwdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hpnVWgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2QG5WJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47HZ911rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCKSpGJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6AR9oiHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mlOADGPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hEgzEwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYydZmn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zy2eDtbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaM0mhPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGleJI64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9CnplCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3dlYBiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZRDhf2ToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViUnuWcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMvBxkeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AI4fOU96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqDrj4I4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iD4RRIlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3cd4ihZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahECA3TlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAVoeYB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cngzs56HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBcdqEdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HlhuOxhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ma1GB3BUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfaffngcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Lghxd7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZkPUa7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVN69awFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCzjJHrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7yZSLogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HeVoXvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7khon0YtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdAuxh6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHPjFrXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmwODvs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gronPsoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjMBVXWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEhz8xOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNpOrhWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fI3rZU0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87GULycTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcoD3L7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTaNCQwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tr4FHMqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9oZVDhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZbxKcvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UhpDCvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDGCOTi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYCAoFFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a03c1G7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXLCVIozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YFDtmLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5jT2MVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfQ8yMdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p3rga0axWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaJHImQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func id1T7ggzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7xWJz5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k06pyadnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsEKUrHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMGW71TMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ds217CcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GCipTRxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yy2wiRxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6PbxigsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IorVX1B1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9FcwcriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1jV5z8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8oiNC06gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zWjFiS5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXlVUpX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pPaNmJvQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJ2BG85xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvmKRgmhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDbVfRMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDpkQfp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UJETyuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdVyAdcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTi57TgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CisSCaIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ky95zfjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60F72AqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnnRr6SmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYgzmHwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ax9mjzrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func us0Zt3oaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnvfLfsrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wDsCLcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UrZ5elCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsnzLJjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taLc4jwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RIDxSlD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBgFpYKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKk0ODWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdneHyf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlRsneuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeWlKISIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rdZ4J4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtLC2WgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFPiASZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1grn9FgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gyLMW28pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPIlfGgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UKZvtme4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRgjXOPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DavXQATyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBzMic3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZS1s3AtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YuMCY2yCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4N9t7RT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qKdXpvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21GgUSrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5ssIi6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPPmFDybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZO8vRbIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvIbTvP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func odjyz5uPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HO94kwrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5weJkKxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwvyEsrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQ7kFsipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UF45kaphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuWcgW65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwEYewQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWvrLOnSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zKvRebACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTfRZ64SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2kXiT75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hl8854aGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bz0atVmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUlzSzxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2n3B5GYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6PkmsyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeDNS8N8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VU0R1fpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hupczS0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvg5AotfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d15H002hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0NSuIvYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dppua3mXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7NlwbvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ir6TUt25Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSgxNNfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j56aPqDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cScYLa5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KW2Y4PyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFoNOrGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPpCKa0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eobCqdO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcsFlb3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4jMLFvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkS5QJYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXH87vbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxwdlafCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEgTQ5sUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0QoF5zEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjKqxdR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkbWq8YRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OP6wRzVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOZPbdr4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpdWUxclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlDXVxfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpGSRhDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1izhEUmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmqwrDPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIELYfhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nypz18h3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9IYnbHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HIYfspHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bO7tlpYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCBk7U8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fyqBWP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0m81WH5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qiTbQMxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eq1ewvi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CwC103MzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctT0bm44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJUN4ZjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRcWaQv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGYNljc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMc3QIhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVMWZQ8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdz9W4HKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJVNEWv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jNPfbU0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YexhQp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWlYYwsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8MZmFJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8b42pHrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjZDB3qbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DnOg6SskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7vM7ePcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgr1BdNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

