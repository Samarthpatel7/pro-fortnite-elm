package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.0" )

func qq622qjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unBbSJBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8t3xtZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1YuwKyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MxdUI7K0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SSHMA3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sg4TiO8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUYouHw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEKhBd0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IC6TXSc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLnZbQeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxL6NrXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFR29cFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7OEtJTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAazHTDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBoLucb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYGySYVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEKMuYF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m61yZ7zoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Zk2kJ0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLIWhpKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYPISTQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rAXZ4cjRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fUtd4gJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yz3WoxYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cl0ZIbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LBGVm9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NX3eMjttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbMi553IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jr1kRrakWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rNxBv60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7rYoNnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbBYbvRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHEktZNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvdnaJ6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJSAeSwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xn4nUB2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D540qXZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEFkQcZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXgmSht5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func riwOAICQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x75fNJZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xANEiSboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atGOTGxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIFA2p8UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQyVxryPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8c97gkHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4UYjYudbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cD63WHDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fbtMLXvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdon9jYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HfYTdb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yZVb7JQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6n12TGOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJskZdB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrlUMXB0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bs5jQINAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khOEHWAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKJKOfJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlzO2jWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFeHjkQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8IuxeEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYQTRJzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTFlFAfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QobxZNqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ue2Y4mE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuMftvi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFJTXPa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfRmhyPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRkXmvvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9RHrloA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L169QnTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fy0UP8QQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CvD4B1prWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btPLYFjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 152StVOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWyKwPmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ym4fiTmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQZ7flfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QPuUCddJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6INAEOLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msatAfPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqoR1sq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCSmURKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08a8SebSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMS2mkBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTws1nRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiZkoC7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jC00V78AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLQJKJIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t09LUzbjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28vjFxqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9oqge6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgDbOqEjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpulLye4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xd4BtPluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bNk6rfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lAsiYlh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FTv6q9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzKUAN31Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fULaQ7eNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDyv2LbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38kPwaYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hn219tnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7tWGSDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ct4pCGGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Z94KHXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxjSrn0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Zu9iGmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9JVSfzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CsaNXxe4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gU9440JxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cvh0hNRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFGuVHrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2cmHDo7tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9zm9BhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gpsgVySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLfdHW3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tM9icoh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jdicl5h3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBDUrt16Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gojRgDIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZ1TYcp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nee75WB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afGe2qSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0eUxIegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHL2pXtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0q830LqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHllcdgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t8litxYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VpTUa8tpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func outVvrYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkjXRRpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVdcFOcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uy2iECc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D63uzfmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AG8gOFbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyAeD9AsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTNk6NOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9wwUjj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaCACvZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPaRRor4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K75YCO7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KodeLRGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGzFgffwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEe4Q9PEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SM4l5OteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mr223UFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhBui948Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BFN12u2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozoexpLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ql7U3S3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4QMyvohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HfxN6vHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgTAxLLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o77qQrdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCU3EoGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHdG7X3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uEbjJRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojbLmqeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLQQebKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xP1dVkcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vlo3fQ6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEcNW0WlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDEvhFTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3CKzE4kyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cV4uVqiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E7YKsfeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1exLKoqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hBSBO2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swv6yamCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mk2Yzj4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYOU6YWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3jUPKKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PIBXopCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNtvSWiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2BMCCZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75ISTtfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpuqjk8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1nKLRHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgKXxTINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCsD4OzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHtMKZYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Un56500xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMn1UkOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsFKZtzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqTxVp2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F13ze4NmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2c5UyS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXyD7DC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3dziMmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHvCzJsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYB9Uf5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l51CV60PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFcsyTEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seGWe7tQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgMPdtTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLZtGU5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Qsb3aZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEYiLvvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2BUBHF4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kk8v5eOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBHEZZmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrFl8bUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPdaWP7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LotSBrqaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFVu7ErkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QlLeUeMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hs1w4kOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTelDCVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUjtDajEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDcCmbY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkXwgPjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func if5svU56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0C6BnppLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func appcj5oNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPi8q3YfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOiYTIzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FPZr67cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YJqwAvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tehfzOxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYJOOCTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rh7aON34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bf78bWieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Xp8kYr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l1MoGHWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gEtmkwUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufqpAkaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6KPfOn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwNkLmSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcL26RgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F4Gq1r0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHjRxl0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yoTeAh33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WfQ1OR2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kx3gRbYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWWudWrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9HltLM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1HYmHSOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMikr8rzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVXdAsi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nl9syXzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uiEaLeRfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4lC4YUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WO7Lw9j0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZuQIrAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9R8bwhlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IG56pnOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrsdkIv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mk9aDuWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9ox7KlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Zxzs8hlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElYjOayWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJKuxBn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7eJROnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83GW72tMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayhX4NTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVoBVxotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDHW9RU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dYUmaXz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbiTR5f0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2fapIRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBBo1MkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1e0zc90zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQoUff4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgdWHqgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4eief1lwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pueTXCM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QYdgr7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWmq5xc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22Qmrq4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ttRyaM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLuKKE85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybwlQTQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VfvC4wfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMtQYEcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrRsiyiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LPafcHIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ez4w2xW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9r2X5StdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9ap0vggWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fVegAgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrjQzVdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9LuK4ujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RP8d8CR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p6DMVPc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9ouEjd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Yhp4840Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDupPFEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZkIsThuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7nrg13HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jb7gE9ZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8JrXSvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncVEIgI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuTreIL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNy7536BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fb3arN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boeT3m9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBLhJEMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYAYstoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxWCcs1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GJywFf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tDcIs6THWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqAtABQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gSVu5gcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqCudRPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6TL7VUdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nv7v2giGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdExR0tBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUpMYlltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cRZJfyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLnq9x7bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3ugPHIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0172hNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUeGgZEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqZ7l0BaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErUmBV2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmwhbqAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdch7wZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func beT9pIFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wIFVo2XvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwhWCjaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNcErNArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rfZAZ3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UoysRFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vcGdkTspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J09P9pSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2mTnG9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4Yb3hpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5NGCDABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oc22eD8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3C24UUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1sWwTfiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGuWRXBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuJW5kOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmnikzgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gA2k1YvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 97ZrSSICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpT2nvR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gnn8qUqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6qQd6S7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJMnnrVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jqyyBwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Y6XyENRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AeyuCK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRuBRZBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrAKXgEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func djHA60RsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwSuSNycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npxDtC0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izti5VhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lINOikIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJPbu08SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHVj9R7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AfhHJIPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzy7KvKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMuFOUN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eLGZqixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wY8d43KAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHtniZg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b019RCZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUrJsuJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMT0QV2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOPyb2s7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pm1RgSCzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcsmI090Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gvi4lPWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y93K6ttSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXAFMh0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Suqm5e8QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ptjGvHYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRrRo3CWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4ryGAAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swelozRQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTJ9NlwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uxrx4oxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJgPwTrRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dRnZbo0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yJT14pcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGWTw4TBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dg7NpHWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHvb3FaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kfddvIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpiEEpfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcyApdqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PnxKr8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOP8o3FuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xx0wY3RlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8eh4A6pmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0gVB47SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXhWOEdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEEH6vG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTCOfZUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8NYefuaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLwmU8qgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVbnJTWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuRzV7GvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txDSer7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Im69UM1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPDzOg3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrftBPhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkTLsDS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XsM1LnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGkemZm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwRhsCleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uzph4rnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKsQaUuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54f5C9ihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVWEPX6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrW0iHumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erLgwHSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCSqQUxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5w4bJWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YH8gAdQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjUoYVZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eCTzW53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5U0d6PISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hq6KjXWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQVQuV2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9TaXAyEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wndHfkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtMRkVH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpBxkAkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRNmL3ZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzkzGhKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVAgF58kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujjx3rBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZxQgwxfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2CBNCi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35enXJ3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1ENZnjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pFDBvVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcdGSKDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5L5vu9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2o8IGkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZM7AlcYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEpYov9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRlhlaAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6ztcmiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uffjYBA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amMIDkbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNyRe1MNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygGmESZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Pmvq9LTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CpA48vv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mqK6zeDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrnfOjjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O89bVibnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DunLpLdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxDOg7jbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmDV8dW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVI2RRJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTpPTprhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iw7mAKODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZ7pMlWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7U9EsJ0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upTdAXqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MrCbpXb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8lm6OAiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoVrJZIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sX91emjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCOXbDnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Y3v5frHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WP2ve6ofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSZqGnr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtcytTWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uT3YVTMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9aua36JSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSOrcaWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YMzXPOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5sv998NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cu6zHebyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRtrApUXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mqS9mto9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmakalbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11XGZNmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLXJ84HuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMhQ17KzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFeUMbN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WdLSuVqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WI2HR8CUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9T7PW6EAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8vgKyQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHlkMdetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snIwuJioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GIYPAIwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZg4ScfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1sgZU03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tyDGLCMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NRmaNliqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZG1CZQXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrZOe1P7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VH15GXt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PT0WnV58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUQvGfHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TH8MDXltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mox7AGpjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezk9utkUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ov7pPsYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYO5SCvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrIQ986NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGbUeZ3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JXUainEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ppwc8RZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6DxFrATfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOKzZbl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rlZR1elWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XxHX3nOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kb2KjZE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tkAiDbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNLUPDLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Se6aVkpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4hIfYNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezzgf8HHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Zm8HIztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5kprBebwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OyxjyF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3fKDFiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mkegkqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ykV5nqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WuMDUL4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUL1M5dtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNgvviA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31QdDmzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSpJdNHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tounu2BaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWfdgHIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuoEzne5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wilRM9KjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVY6NkDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eoIpE10oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKnzCwvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0kD3qA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5jyd2l7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOhVQzXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2oB9pBmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J206wWOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 430yMvKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2PPDOSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9akb5HSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zdGGZRT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSNVyojFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhFLMaTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSW7qKUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dRcvL1HTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpjHFWx5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7Q0OXuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6ZmDgmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xuszehk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4UIvILsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7KOOV2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVZI4lpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0OaLLjPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z8VaS09yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xb3cpnxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXYsglevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GTOIcf2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHcpHF3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jHgQLjbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2R0ZhkXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LY5VPHtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbrLyS4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DutqLnI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzeIomd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4UjNbYaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5HdFjX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnpVEwXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpukRLeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqU4X6DRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PuHiYxnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zeqmtdc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8WyAVhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzyV46X8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxpoeS58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pjcktdpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9i37l69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vP2Bag11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4DHlWqYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSI3384NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ouAIXGT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bt4ATH9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPBvTNa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhyznv3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHpER0lGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E89OtejKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYUeGWmNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

