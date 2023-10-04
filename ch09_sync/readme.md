# 同步

## mutex
鎖定共享資源，多個程序不可同時使用該物件。
都有多個線程進入時，其他協程要等釋放才可以再訪問。
```
func add(i){
    mutex.Lock()
    sum += i
    mutex.UnLock()
}
```

2222