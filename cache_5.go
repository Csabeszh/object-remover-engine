package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.5" )

func E9FCegp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6121CchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTIjPdXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3gp77mZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWPw6FdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5c4V56MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoTV82t5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0s8AhXoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsJrzdTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Hy72TMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1W9bFcxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLWUzK9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A9rXYiMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgEx2itpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2x8hUbyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6FKAuFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqrlJSQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRRpamtaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfKB9dvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mCbI4H1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxNES6xMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 184IKHqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5TBiodDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yMstJgzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhgjxrTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IntQAZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uzc36mvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSbCoCseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIQ0X7TwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAJL4VyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpgFyWW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKJfp1VkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYZKENnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flKqNmibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDxAUSTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2CFVkfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4whdxnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7NkC7wrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1jojrkV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qllORNleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKSR8N2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NacUPZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBCRmNzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nk9MBjmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xwbSraaPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hzh9BnDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1f3k6DbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AglRQrYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8oKFvXcFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsuAoHfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p36Hp0qyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjrfwxgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OgC6k5ThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KkBvdPPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FQitYDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8PyhJskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6BIYORiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLMCx1SHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G6udMFcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lsg0vySJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tdp2sJmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fqNNz7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViorOVOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLh2upmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGEodNGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTSPLBdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CG1MYh5KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMQNV8wPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUkQ7LoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHrCaZgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 341NiuJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UI2o5xxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fELN9IMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2zZPUKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRkx9zqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1FwQEavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQiZVvgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzkR7WpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZuJYuM5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgPVlyNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L22kKikeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhfRGKiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKE170lHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiLEe5U2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCAZqxMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHtuAb2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frT9McQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4RQes1jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lCozpejMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJ8y7uBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDeM2pgIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxD6h5cuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5mQ4ZndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OXt85O8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUEJYItSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdpoDsuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrBAjuOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f831fDDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZCi5xQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cXG2oDRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O25xNFl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRnl8wx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuqcUtHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuZuRnASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0uPlw6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAouTnFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tGLeXAO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OqT7ymdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wp9yOgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewSAXxnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjgRYvhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKPMv9K8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJTxJZIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSC5VDEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEgWQZBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzNBOaBWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KrmEvUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Faj94TgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ps7mp4n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJSCLddBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxY3ZNW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8todDo2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSoyoNCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2blgKhV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nncqGkjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IyUidiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxyXA7WdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRy9HjSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y55OfTOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLwXwceGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LYXLFypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rDVwZLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDxDoIpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrI8GKgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKQUB8cyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRXEohYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cRPJc3rLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mim0yddaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4l2NnN6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vp32KqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wyF49eLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPyrkauTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zd4TntNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RT3CcZYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzXlf2IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ss9iayJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8ZsdZXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w73cxvqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqL93tPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFw4iEs8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCGdkf9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QM57nyxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhq2XC7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ml4CTCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KP6JIKq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZ1IIv6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebniaLadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Cbfp4IPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q47D6KzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vF5HE1zsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pA35jB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1IAcgzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suqfoc47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xvMICXjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkiDaM5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWtJIuiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsYLPvC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HcO2UKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEGQNUXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7pXoH0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZ8piNZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNCWvmUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edv7bZzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5t6KgjzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwxpMeWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3CR9qAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hT7Jn0DyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1UuQqdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLLbNzpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Hh8133pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPKJxrXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUjrGRqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUEaKM0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mF2nd0upWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZrXXvxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 07duZYCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwVjxY4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KB6iaALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duakPeT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7oHxB3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxKkO7StWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UjW0RCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GPKlsFbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zv15bnTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQYD2igaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKNMNXXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2zEv0bYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wb9HvfsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bZfbGg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3pZ4PUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7u9RZaklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQO9b9iYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXcl17NkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KODepPUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyCWFJMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xH4nKPHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CbN3nVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DBID6QDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFLpvn6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaxzBtxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnQYZHNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KrqFvKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiMBCRByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rewcjXBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxBSSDrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bapdxTkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWxUG9KPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NIyOC7VIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUVgM1e2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pech50FuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVCT6VSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvrw5bPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otoGwXvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoW7e9XIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lqv1ZAq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tEG4X2jEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9MAQq2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2B1D8hYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HD3d3or6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gyhfvwm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53RZwpnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIPCjYCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFjv9IA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uJljOrCeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aynlXx4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVwUCGZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhUU11xiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6m4mAQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpgGPq77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWFum3i0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TexXgVsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzQ4jiyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vK3HX8N5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XBpeRMq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqTaBnpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miCi3faFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5TEpXROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bT13jNhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CWOVJX8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeaWzDpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzdKCWoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHifPD2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k0ghVVUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CzEY2AgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhDDvG6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ja9FbHQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSFUdsGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNmQu8gbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qcr3mWWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTbugmb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FaTFOQnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CgZQBKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4S8uoh7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfVKdlSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tn8AaOt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oGnnaDwbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxBqdgnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func csETFtRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLCKQibeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RiDqPUAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCdGLn70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7RzryScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DqZ51BcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KjI8fZebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gd6mwLJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pn9mQRhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fek2NDNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dg6t0OJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPuNf2PQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDvW9z0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3nr3xcxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MCn1JrZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KTdwoQ9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4cWav58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9q05ctM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIqtjaZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rCoZwOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffK9cGIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fUFVULm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yibbLMSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1pVjnfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmGR8aGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qcn4GWSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqLXHaU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsL7xZtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOpwgxBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTVgfshGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4yXZ6TqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDxg5TcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4mpbNGzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59tZmqa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKsuWOpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YBuxNhdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqvmGoMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iq0HOjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNXVEgQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxWiYZs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ge8T2IUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIoN2uOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z6KqRsJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUNIxsnzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsBJGmQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9AHrY2qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ui1jq4F1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIY8R66HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtV2aYc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vK7reNNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsNjKtqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fo2ndEXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qI9yLMuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvETLZRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qejnC9TaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6j2b5W6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mjiALmgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpAjX52fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlHIRWNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCCzCrdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAdmPAOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kkTcByqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDysR5PGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRrEpBlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvsW8JCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmEPmTDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iICMhRtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6iyDSfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JnOkwmzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhttzGldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfqL915MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RAtE02HxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpyGV4rLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U4h0VDCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czhD9hHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sf4FBl9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRTcP3xMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erL4Za89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxXofaXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mFHq66JKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0GNf9dRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stV8Vs5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyhWYhRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCsjYR88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fK6DNwGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e7WB2s4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3BLZwagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlLAhF0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UnJMPkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6rRHpVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9CukUNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktjF4Y6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gg9D0NJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIAzXPK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvUax3GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyugbdPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuDbTnahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NltxzMZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6R9cpOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fgz5xYizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSQ23LKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ecJlpnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eypm2yTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztKEuWsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Lk8zhutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvZuuAQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLAjjUl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5Bias6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onleG43UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqQ3dHcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zX4EGCblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuFwFTqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlaLpOhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjcfXB8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHeIBzdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rwgfee7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5OCo4JvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLivUk3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNJOaBA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VyStJM7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nG8eVcXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XgOIdOB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzkboZQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wS1bwPX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyVioif8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdjrsYfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ITZRoTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wxs8fi9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GO7qOLFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdJc6tEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBv0Y1iLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxzcXIi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7r4quK0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7Z3Qs9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFAVOC2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZyi4OyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtwH6RqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4G2yOnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DPp2sa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzqqitKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 169GhOFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnc7nPtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CsukRmRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WLuIu71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpTlOKTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlsgrjocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRz1w2IiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7TWNwVeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfbBwUv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dx7KogPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mk8Ngt8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1ikOF7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0LhwlcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEEs7k84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38nyZvSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdvBcb4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsPu0YovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func laQz9jtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RhZRLQBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dn6VfwhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MfYTd1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNANnNPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AD62EZL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCJU1BK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nz4iH2EFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUqG6lWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1019xATPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnUBcEkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDjMIRncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tIkuv13aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I87hvEXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBSetfLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nlq4b1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNRs6gTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qr2x8IKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3LquktSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E95QYFEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u932xv1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7LSU0QLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mx80LIR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HNboXiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpRmaMw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aKcpfZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ok6dw7TVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZ6A5DIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqo3Q22NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y01qd1XAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJ1crSyMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qSyc3B8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhRyZMgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8IBJgFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqsIAYHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f74uGybuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEcJxBMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DEEGgDlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPO9BgKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pU1nCWrlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsdhpuWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func em2PJXnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hpsXmEtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQj0rcfVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gz1qVDitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vq1NhSEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRXVK33TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHpUx9NWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PlyAV7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vNzSGYMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Phs6j2e1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2ua8nrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7vajsr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoVpBtHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPFX8fyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v8QCBtQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wqQ584LXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEbrkGIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCkRfsQAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJhNnBEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PkO9xZOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o54bXmJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bS9G2yCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jb5Sn2SgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pCV5xmXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCKBPtNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PS55xXSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJHbKc5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMoaFikCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func phZD6BkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yU7Zz2FjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ylrDcqnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4ELKeO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biZ6qMwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrzjYnqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuIJalc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FOSYgo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrntyLlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoVG04eOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VipVg489Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKErQQM5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQbzV1tdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFQJ8yzuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lrqy5CQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgAXLjXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwCuR3DKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vfaKdaJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HDNAHiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXwo6RfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFbGuQ9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVhrvszcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKy9btMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1c6dnAf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4SykDNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0G3IB15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1qgDYylWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SapMdg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auzcxn2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tUu1AsMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edJW5ctHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faWUNFKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLxUbTHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKsnlYpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V8NTlMpAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovY6iaUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9Jvxr6iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jj9tiDpTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fokjYgiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evUAx22kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANXo6JHIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wm0ASfJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqCnUGLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iD2gKAQjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZyIRmcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPwaHByqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2WQ7OPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yqsfw6E5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNUyEDNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rqxf5ueRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VsHucl6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpgpAkSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3VwXg3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vuA2OSPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfeJU36FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lExVxYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmeLiazoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HjoPNOaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzkZG8jjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LdPwnL6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1UQUPW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mUeo5CXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHPMVA0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QVOchGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DOtLd9I3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qpvk73wpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IknJ3dEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bE15KUOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xL4Cicu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOO0hWsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pM4kzmCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdaTsNQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3whhWFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ac4xknqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nogy3GX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXA34Ku4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKQvorooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GnuPW7KwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9fq1sblZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GU87L7bAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQcq1NozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwuT8EvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fEyMTepzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QY0VJPIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CsaPnujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yjfpznZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8XyeAxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlcXDCIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uI2JTEW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtHJRuB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Qjdi0n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mT1GDZIOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89CQb7OVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3LAvUQg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XkPpDaGhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfhvAzVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0b98XY1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ywt7t2VzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeYanW8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tifonc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0Rq8axyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uu4Tb8NUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9m0JMUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbubqGwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTmHHC7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shScHaMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsYPPmWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtxUnrknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKvnnWCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNCWnjyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hi5jbjnZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvAlS1RmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEhsLtfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWwYK2mJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6muZSkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyFC85CPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRRdADcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfntojfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FEuMy8luWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvhjz1GnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HBjMaX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ABpKoCuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqDu4gPaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGD5Ygf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOZthUndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHfQxqf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

