# workflow

## workflow-be (Backend)

**Building API with GO**

Requirements
 - **MUST** use go module
 - **MUST** use PostgresSQL
 - **MUST** use MUST use `go module go mod init github.com/<your github name>/workflow`
- **MUST** use go 1.23 or above
- api port **MUST** get from environment variable name `PORT` (should be able to config for api start from port `:2024` ) 
- database url **MUST** get from environment variable name `DATABASE_URL`

**Business Requirements**
- Item workflow system
- ให้สร้างระบบ REST API เพื่อจัดเก็บข้อมูลการเบิกงบ (item) ซึ่งความสามารถระบบมีดังนี้
  - ระบบสามารถจัดเก็บข้อมูลเรื่อง (title), ยอดขอเบิก (amount), จำนวน (quantity), สถานะ (status) และผู้สร้างคำร้อง (owner_id)
   - ระบบสามารถเพิ่มข้อมูลการเบิกงบใหม่ได้
   - ระบบสามารถดูข้อมูลการเบิกงบทั้งหมดได้
   - ระบบสามารถดูข้อมูลการเบิกงบทีละรายการได้
   - ระบบสามารถปรับเปลี่ยน/แก้ไขข้อมูลการเบิกงบได้
   - ระบบสามารถปรับเปลี่ยน/แก้ไขข้อมูลสถานะการเบิกงบได้
   - ระบบสามารถลบข้อมูลการเบิกงบได้
   - ระบบสามารถป้องกันการเข้าถึงข้อมูลด้วยการมี authentication
------
"postgres://postgres:123456@localhost:5432/workflow"
goose postgres "postgres://postgres:123456@localhost:5432/workflow" status    