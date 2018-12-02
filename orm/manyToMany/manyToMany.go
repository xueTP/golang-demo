package manyToMany

import (
	"fmt"
	"golang-demo/orm"
)

//ForeignKey AssociationForeignKey
// jointable_foreignkey association_jointable_foreignkey

type User struct {
	ID        int32  `gorm:"primary_key"`
	Name      string `gorm:"column:userName;type: varchar(20);not null"`
	Age       int8   `gorm:"default:0"`
	Refer     int32
	Languages []Language `gorm:"many2many:user_languages;AssociationForeignKey:UserRefer;jointable_foreignkey:UserId;association_jointable_foreignkey:LangRefer;"` // 多对多
}

type Language struct {
	ID        int32  `gorm:"primary_key"`
	Language  string `gorm:"type:varchar(20);default value:''"`
	Level     int8   `gorm:"type:int(4)"`
	UserRefer int32
	Users     []User `gorm:"many2many:user_languages"`
}

type UserLanguage struct {
	UserID     int32 `gorm:"primary_key;Column:UserId;AUTO_INCREMENT:false;"`
	LanguageID int32 `gorm:"primary_key;Column:LangRefer;AUTO_INCREMENT:false;"`
	Other      string
}

func ManyToManyDemo() {
	orm.DB.DropTableIfExists(&User{})
	orm.DB.DropTableIfExists(&Language{})
	orm.DB.DropTableIfExists(&UserLanguage{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&UserLanguage{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&User{})
	orm.DB.Set("gorm:table_options", "ENGINE=Innodb").CreateTable(&Language{})
	//orm.DB.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(&UserLanguage{})

	//language
	language1 := Language{Language: "english", Level: 1, UserRefer: 1}
	language2 := Language{Language: "中文", Level: 1, UserRefer: 2}
	language3 := Language{Language: "lingua italiana", Level: 2, UserRefer: 3}
	language4 := Language{Language: "Deutsch", Level: 2, UserRefer: 4}
	orm.DB.Create(&language1)
	orm.DB.Create(&language2)
	orm.DB.Create(&language3)
	orm.DB.Create(&language4)

	// user
	user1 := User{Name: "sum", Languages: []Language{language1, language2}}
	user2 := User{Name: "bod", Languages: []Language{language3, language4}}
	orm.DB.Create(&user1)
	orm.DB.Create(&user2)

	// 多对多
	// userInfo := User{}
	// orm.DB.Model(&User{}).Where("userName = ?", "sum").Find(&userInfo)
	// fmt.Printf("User is : %+v", userInfo)
	// orm.DB.Model(&userInfo).Related(&userInfo.Languages, "Languages")
	// fmt.Printf("Language is : %+v", userInfo.Languages)
	// orm.DB.Model(&userInfo).Association("Languages").Append(&language3)
	// orm.DB.Model(&userInfo).Association("Languages").Append(&language4)
	// orm.DB.Model(&userInfo).Association("Languages").Find(&userInfo.Languages)
	// fmt.Printf("Language is : %+v", userInfo.Languages)

	// 关联模式
	orm.DB.LogMode(false)
	userAss := User{}
	orm.DB.Model(&User{}).Where("userName = ?", "sum").Find(&userAss)
	fmt.Printf("User is : %+v \n", userAss)
	orm.DB.Model(&userAss).Related(&userAss.Languages, "Languages")
	fmt.Printf("Language is : %+v \n", userAss.Languages)
	orm.DB.Model(&userAss).Association("Languages").Append(&language3)
	orm.DB.Model(&userAss).Association("Languages").Append(&language4)
	fmt.Printf("Language is : %+v \n", userAss.Languages)
	orm.DB.Model(&userAss).Association("Languages").Delete(language1)
	fmt.Printf("Language is : %+v \n", userAss.Languages)
	count := orm.DB.Model(&userAss).Association("Languages").Count()
	orm.DB.Model(&userAss).Association("Languages").Replace(language3)
	fmt.Printf("Language is : %+v \n count : %d \n", userAss.Languages, count)
	orm.DB.Model(&userAss).Association("Languages").Clear()
	fmt.Printf("Language is : %+v \n", userAss.Languages)
}
