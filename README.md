# All the explaination about function defination are given below

## Used at Root ↓

### App
- app is a fiber instance which is of type -> *fiber.App

### Get, Post, Put and Delete functions
- this method takes two arguments first -> path which is of type string, second -> handler function which is type error (handler function should accept pointer of fiber.Ctx as parameter)

## Used in databse directory ↓

### Open
- this method takes two arguments first -> dialector, second -> options. It returns pointer to gorm.Db and error 