# Kỹ năng học được
Restful API, kiến trúc code, logging, monitor, CI/CD

# Kiến thức nên biết trước khi tham gia
Nhập môn lập trình

Xin chào các anh em ham học hỏi, trong khoá học này mình sẽ hướng dẫn các anh em sử dụng Golang kèm theo các công cụ để từng bước giúp anh em tự tin có thể xây dựng được một hệ thống Backend phục vụ cả nhu cầu cá nhân và đi làm thực tế.

Khoá học không chỉ học code mà nó còn học cả tư duy hệ thống và để đảm bảo anh em 100% có thể ứng dụng được thì khoá học này sẽ không học theo các video quay sẵn mà mình sẽ tổ chức học qua google meet, anh em đăng ký xong mình sẽ đưa vào nhóm và bắt đầu quá trình trao đổi học tập.

Khoá học hỗ trợ chọn đời các vấn đề trong phạm vi khoá học nên anh em hoàn toàn yên tâm về các vấn đề chưa hiểu, mình sẽ có trách nhiệm giải thích cho để khi nào hiểu thì thôi, nội dung khoá học anh em tham khảo video giới thiệu và phần nội dung khoá học bên dưới.

Thân ái và hẹn gặp anh em trong lớp học nhé.


## Git workflow

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