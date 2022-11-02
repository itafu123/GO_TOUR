# 錯誤

## Error
可預期，且不影響程式運行

## Panic
會導致程式中斷的錯誤

## recover
類似try-catch這類的捕捉語法。
因為 defer 在程式碼最後必執行的特性，recover 必須和 defer 配合使用。
```
    defer func() {
        if err := recover(); err != nil {
           fmt.Println("couldn't convert number:", err)
        }
    }()
```

## Defer
此函數保證文件關閉後，一定會被執行，不管有無發生錯誤