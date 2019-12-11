// 用 `User` 结构体创建 `delete_users` 表
db.Table("deleted_users").CreateTable(&User{})

var deleted_users []User
db.Table("deleted_users").Find(&deleted_users)
//// SELECT * FROM deleted_users;

db.Table("deleted_users").Where("name = ?", "jinzhu").Delete()
//// DELETE FROM deleted_users WHERE name = 'jinzhu';

通过使用struct内部的table字段，您几乎可以做到：

type Product struct{
  ID int
  Name strig
  Quantity int

  // private field, ignored from gorm
  table string `gorm:"-"`
}

func (p Product) TableName() string {
  // double check here, make sure the table does exist!!
  if p.table != "" {
    return p.table
  }
  return "products" // default table name
}

// for the AutoMigrate
db.AutoMigrate(&Product{table: "jeans"}, &Product{table: "skirts"}, &Product{})

// to do the query
prod := Product{table: "jeans"}
db.Where("quantity > 0").First(&prod)

不幸的是，当您需要查询多个记录时，这不适用于db.Find() 。解决方法是在执行查询之前指定表

prods := []*Product{}
db.Table("jeans").Where("quantity > 0").Find(&prods)
