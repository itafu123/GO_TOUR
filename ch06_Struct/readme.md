# 結構體 接口 工廠函數 繼承

## 結構體
```
type 名稱 struct{
    名稱 類型
}
```

---

## method
像是class裡的方法，提供你可以在某種型態上(struct, int, string)定義方法
```
func (名稱 任何型態) method名稱() 回傳型態{
}
```

---

## 接口
```
type 名稱 interface{
    名稱 類型
}
```

- 要每個函式都實現才可以

---

## 繼承 組合
- Go沒有繼承概念 只有「組合」
- 如果外部類型和內部類型有一樣的方法，則外部類型的會覆蓋內部類型