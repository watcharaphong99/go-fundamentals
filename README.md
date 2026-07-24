# Go Fundamental & DSA

โปรเจกต์สำหรับฝึกเขียน **Go พื้นฐาน** และ **Data Structures & Algorithms (DSA)** ด้วย Go  
แต่ละโฟลเดอร์เป็นโจทย์/ตัวอย่างแยกกัน — อ่านโค้ด แก้ไข แล้วรันดูผลได้ทันที

## สิ่งที่ต้องมี

- [Go](https://go.dev/dl/) เวอร์ชัน **1.25+**
- ตรวจสอบ: `go version`

## โครงสร้างโปรเจกต์

```
fundamental/
├── helloWorld/              # เริ่มต้น — Hello World
├── go-variable/               # ตัวแปร (var, type)
├── go-package/                # โครงสร้าง package / module
├── go-concurrency/
│   ├── go-routines/           # goroutine, channel, sync.WaitGroup
│   └── goroutines-workerpool/ # Worker Pool pattern
└── DSA/                       # โครงสร้างข้อมูล & อัลกอริทึม
    └── main/
```

## วิธีรัน

แต่ละโฟลเดอร์รันแยกกัน ไปที่โฟลเดอร์ที่ต้องการแล้วใช้คำสั่ง:

```bash
# โฟลเดอร์ที่มี go.mod
cd go-package && go run main/main.go
cd DSA && go run main/main.go
cd go-concurrency/go-routines/main && go run main.go

# โฟลเดอร์ที่ยังไม่มี go.mod
cd helloWorld && go run main.go
cd go-variable/main && go run main.go
cd go-concurrency/goroutines-workerpool/main && go run main.go
```

## แผนการเรียน (แนะนำ)

| ลำดับ | โฟลเดอร์ | หัวข้อ |
|:-----:|----------|--------|
| 1 | `helloWorld` | โครงสร้างโปรแกรม Go, `package main`, `fmt` |
| 2 | `go-variable` | การประกาศตัวแปร `var`, type พื้นฐาน |
| 3 | `go-package` | module, package, การจัดโครงสร้างโปรเจกต์ |
| 4 | `go-concurrency/go-routines` | goroutine, channel, `sync.WaitGroup`, mutex |
| 5 | `go-concurrency/goroutines-workerpool` | Worker Pool, backpressure, buffered channel |
| 6 | `DSA` | map, slice, sort, การนับความถี่ |

## รายละเอียดแต่ละส่วน

### helloWorld

จุดเริ่มต้น — รันโปรแกรม Go แรก แสดงผลด้วย `fmt.Println`

### go-variable

ฝึกประกาศตัวแปรแบบ explicit (`var a int = 10`) และการใช้ type

### go-package

ฝึกจัดโครงสร้าง module (`go.mod`) และ entry point ใน `main/`

### go-concurrency

**go-routines** — ตัวอย่าง pipeline งานผ่าน channel:

- ส่ง job ผ่าน `jobCh` ให้ worker หลายตัวประมวลผล
- รวบรวมผลด้วย `resultCh`
- ใช้ `sync.WaitGroup` รอ worker ครบก่อนปิด channel

**goroutines-workerpool** — ขยายเป็น Worker Pool จริง:

- producer ส่งงาน 1,000 ชิ้น
- worker 3 ตัวประมวลผลพร้อมกัน
- buffered channel เป็น backpressure

### DSA

ฝึกโครงสร้างข้อมูลและอัลกอริทึมใน Go:

- **Map** — สร้าง, อ่าน, ลบ, ตรวจ `ok`, นับความถี่
- **Slice** — `append`, slicing (มีตัวอย่างใน comment)
- **Sort** — เรียง key ของ map ด้วย `sort.Strings`
- **Algorithm** — นับตัวอักษรซ้ำ (มีฟังก์ชัน `countDuplicates` ใน comment)

> ในไฟล์ `DSA/main/main.go` มีโค้ดตัวอย่างหลายแบบใน comment — ลอง uncomment ทีละบล็อกเพื่อทดลอง

## แนวทางฝึก

1. **อ่าน** — เปิดไฟล์ `main.go` อ่านทีละบรรทัด
2. **แก้** — เปลี่ยนค่า ลอง edge case หรือ uncomment ตัวอย่างใน comment
3. **รัน** — `go run` แล้วดู output
4. **ขยาย** — เพิ่มฟังก์ชัน/โจทย์ใหม่ในโฟลเดอร์ที่เกี่ยวข้อง

## หัวข้อ DSA ที่แนะนำให้เพิ่มต่อ

- Array / Slice — two pointers, sliding window
- Stack / Queue — ด้วย slice หรือ channel
- Linked List — struct + pointer
- Binary Search, Sorting algorithms
- Tree, Graph — traversal (BFS, DFS)
- Hash Map — โจทย์ LeetCode ระดับ Easy

## อ้างอิง

- [A Tour of Go](https://go.dev/tour/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
