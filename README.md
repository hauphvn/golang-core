Jack' Blog 
=
Welcome to Go Core
-
This is a mini project that research about Golang.

### Packages

Một chương trình Go thì gồm nhiều package  
Khi chạy một project Go thì package main sẽ chạy đầu tiên.  
Theo quy ước, khi ta import package math/rand thì chương trình sẽ import package rand.  
Chúng ta có thể import nhiều package bằng cách:  
```go
import (
	"package 1"
	"package 2"
)
```

### Exported names

Trong Go, một biến sẽ trở thành global nếu được viết hoa chữ đầu tiên.  
Ví dụ: 
```go
package animal
Tiger : "Con cọp"
```
Khi ta import package animal thì ta có thể sử dụng thoải mái biến Tiger.  

### Functions
Cách viết hàm trong Go:  
```go
func hamA(number1 int, number 2 int) int{
	return number1 + number2
}
```
Như ta thấy, hàm A phía trên nhận vào 2 tham số và trả về một kết quả kiểu số nguyên.  
**Chú ý:** Trong Go, cuối dòng code ta không cần dấu **;**  

Cách khác để viết hàm trong Go:  
```go
func hamB(number1, number2 int) int{
	return number1 + number2
}
```
**Chú ý:** Ta có thể khai báo kiểu dữ liệu ngắn gọn như trên hàm B nếu tham số có cùng kiểu dữ liệu.  

### Multiple results 
Điểm đặt biệt trong Go là hàm có thể trả về nhiều kết quả.  
Ví dụ:  
```go
func swap() (string, string) {
	return "meo meo", "gau gau"
}

func main() {
	a, b := swap()
	fmt.Println(a, b)
}
```

Khi chạy chương trình sẽ trả về kết quả : meo meo gau gau  
Tương ứng, a: = meo meo, b = gau gau  

### Named return values

