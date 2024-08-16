# Why Go?
## Tại sao lại là Golang (tại sao không phải JavaCript, PHP, .NET)

- **Python** thua xa hiệu năng, nhưng độ tiện dụng thì Py hơn
- **Java** = Nhà giàu, mạnh mới chơi được, máy ảo tốn ram
- **NodeJS** = 1m2 vài nghìn tk Node, nhiều trường phái các kiểu, hiệu năng không được như mong đợi
- **PHP** = app nội bộ ok, Lavarrel nghiệp vụ ok
Chinh phục project hiệu năng cao = Go sẽ là lựa chọn Ok

## Lượng Star của Go trên Github 122k cũng từng là ngôn ngữ Hot nhất. 
Thực tế Go mang lại

### App phần lớn chạy trên Cloud thì Go giống như sinh ra cho Go. Terraform, Docker, Kubernets đều được viết bằng Golang
### Điểm mạnh của Go:
- Khai báo không dùng, ko cho combine
- Cú pháp đơn giản, ít keywords
- Đồ chơi Built-in mạnh mẽ (Built-in: những gì ngôn ngữ hỗ trợ, chỉ cần lấy ra dùng, những core của ngôn ngữ rất OK)
- Khởi động chương trình rất nhanh, dễ đóng gói => rất phù hợp với microservices (Mua cloud, 1 core 2GB RAM Go OK nhưng Java thì khó)
- (Về mặt ngôn ngữ) Performance ấn tượng (có thể tiệm cận C++) + Quản lý bộ nhớ tự động (Tổng thể hệ thống ngôn ngữ chỉ là 1 mảnh ghép trong bức tranh lớn.
vd: code logic hoành trang, db cùi, hạ tầng lởm, đường truyền lởm, load ?? không có. thì cũng tạch).
vd: ngôn ngữ Ngon. Logic như ... không nuốt nổi!!!
- Cộng đồng đông không kém quân nguyên (Ở VietNam thì không đông lắm.)
- Đặc biệt hộ trỡ Concurrency rất ấn tượng (đa luồng. Go hỗ trợ rất mạnh, và thể hiện ở dạng rất đơn giản. Đơn giản hóa được vấn đề khó nhằn.
Với những bài toán hay sử dụng thì học Go có thể chơi được đa luồng.)
- Có thể Build ra các platform khác nhay: Windows, Mac, Linux.

``` GOOS=windows GOARCH=amd64 go build ```

## Syntax: 
 Go routine: line thread? có thể hiểu như là thread, thread này do go runtime tạo ra trên cái user space trong máy.
 Mỗi 1 request chạy trên 1 cái go routine.
 ```go
 go func(){
    fmt.Println("Hello!")
 }
 ```
**Go là 1 dạng compiled language** 

| Compiled language             | Interpreted                   |
| :-----------------------------| :-----------------------------|
| C, C++, Go, Fortran, Pascal   | Python, PHP, Ruby, JavaScript |

- Lý thuyết thôi.... Virtual Machine cũng có thể nhanh thôi!

- Frontend thay đổi nhanh. Backend thay đổi chậm hơn, thậm chí là không cần thay đổi.

- Lý do chọn Go cũng vì ít cạnh tranh. cũng mong là có việc làm lắm!!! Nếu ổn thì có thể phát triển nhanh được

Setting Goroot Goland tự Setup đường dẫn 

```
go build -o app . (build file thực thi)
.app
```