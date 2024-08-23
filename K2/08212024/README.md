## Struct

đại diện cho đối tượng cấu trúc mà mình muốn đóng gói nó lại - hoàn toàn tương ứng với class của java

syntax: 

```
type Strudent struct {
    Name string
    Email string
    Age int
}
```

public: chữ hoa
private: chữ thường (private của cái gì? của cái pagkage đang đứng)

package oop

func main() {
    var student Student = student{
        Name: "Manh",
        Email: "nguyenmanh@gamil.com"
        Age: 100,
    }

    fmt.Println(student)
}

thư viện: https://github.com/davecgh/go-spew.git

Zero value???? 
- string ""
- int 0
- bool false
- pointer nil


tìm hiểu kỹ về type tý


cmd + shift + t: test function, test for file
argument là tham số truyền vào


2 loại Method trong Golang
method và reiever method

entity: object đại diện cho nghiệp vụ
repository: 



Một thông điệp commit tốt nên ngắn gọn, rõ ràng, và giải thích được lý do hoặc mục đích của thay đổi.
### **Format Cơ Bản:**
```bash
<type>(<scope>): <subject>

<body>

<footer>
```

### **Giải thích format:**
- **`<type>`**: Loại thay đổi, ví dụ như `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`, v.v.
- **`<scope>`**: Phạm vi thay đổi (không bắt buộc), ví dụ như `authentication`, `router`, `database`.
- **`<subject>`**: Một câu ngắn gọn mô tả sự thay đổi (không quá 50 ký tự).
- **`<body>`**: (Tùy chọn) Giải thích chi tiết về thay đổi, lý do thay đổi, và thông tin bổ sung nếu cần. Viết theo dạng câu và ngắt dòng sau 72 ký tự.
- **`<footer>`**: (Tùy chọn) Thông tin bổ sung như issue tracker ID, breaking changes, hoặc deprecated features.

### **Ví dụ Cụ Thể:**

#### **Commit Thêm Tính Năng Mới**
```bash
feat(authentication): add JWT-based authentication

Implemented JWT-based authentication to improve security.
Users can now log in with a token that expires after 24 hours.
This change replaces the previous session-based authentication system.

BREAKING CHANGE: Old session-based authentication is no longer supported.
```

#### **Commit Sửa Lỗi**
```bash
fix(router): resolve issue with route handling in POST requests

Fixed a bug where POST requests were not being routed correctly
due to incorrect middleware configuration. This fix ensures that
all routes handle POST requests as expected.

Closes #123
```

#### **Commit Cải Tiến Mã (Refactor)**
```bash
refactor(database): optimize query performance for user data retrieval

Refactored the SQL queries used in user data retrieval to improve
performance. The new queries reduce load time by approximately 30%
in large datasets.
```

### **Một Số Lưu Ý:**
- **Tránh commit chung chung** như `fix bugs` hoặc `update code`.
- **Sử dụng tiếng Anh** nếu bạn làm việc với một team quốc tế, hoặc nếu dự án của bạn có thể được tham khảo bởi người khác ngoài công ty.
- **Breaking Changes**: Nếu có thay đổi lớn mà sẽ phá vỡ hệ thống cũ, bạn nên ghi rõ trong phần `BREAKING CHANGE` ở footer.

Một thông điệp commit như trên sẽ giúp người khác dễ dàng hiểu được mục đích của thay đổi và lịch sử commit trở nên rõ ràng hơn.