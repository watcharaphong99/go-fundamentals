broker = server ของแต่ละตัว เช่น 3 
producer = ตัว produce message เข้าไปใน queue
consumer = ตัวรับ message จาก queue
topic = หัวข้อที่ producer จะส่ง message เข้าไป 
patition = การแบ่ง topic ออกเป็น ส่วน เช่น topic 1 patition ที่ 0 topic 1 pation ที่ 1 topic 1 patition ที่ 2 จำนวน patition แบ่งได้เรื่อยๆ
สำคัญเราสามารถมี consumer เยอะสุดเท่ากับ patition ที่มีอยู่
ถ้าอยากมีหลาย consumer ก็ต้องมี หลาย patition 
==================
detail 
producer = เวลา producer produce message ถ้าเกิด message นั้นไม่ได้ระบุ id ของ message มาด้วย  producer จะ randome ของ load balance ของแต่ patition  ถ้าเรา message นั้น ระบุ ID มา producer จะ ทำการ Hash เพื่อหาว่า Id นั้นอยู่ที่ไหน messageId เดียวกัน จะวิ่งเข้า patition เดียวกัน  message ที่อยู่ใน patition จะเรียงลำดับกัน ใน pation ของตัวเอง เวลาต่อกับ consumer consumer จะดึง message ที่น้อยกว่าออกไปก่อนเสมอ 
แต่ละ patition จะมีเลข offset เป็นของตัวเอง 

consumer = จะมีคำศัพท์อยู่ คือ consumergroup ที่กลุ่มเดียวกันแต่อาจเป็นคนละ consumer ที่อ่าน topic เดียวกัน และได้ message ไม่ซ้ำกัน ถ้าหาก consumer ตัวนึงตายไปก็อาจมี consumer นึงที่อ่านข้อมูลมากกว่า 1 patition ข้อมูลจากทุก pation จะต้องถูกอ่านจาก consumer เสมอ จะไม่มี pation ไหนอยู่ว่างๆไม่ถูกอ่าน 
ตัวอย่าง มี 3 consumer แต่ 1 consumer ตายไป process ตัวนึงจะทำงานชื่อว่า Rebalance  การ หยุดและเริ่มต้นใหม่ว่าใครจะได้ pation ไหนบ้าง เมื่อจบการ Rebalance จะมี consumer ที่เป็นเจ้าของ pation เสมอ ถ้าหาก มี consumer แต่ไม่มี pation ก็อยู่ว่างๆปเลยเพราะไม่มีเจ้าของ

