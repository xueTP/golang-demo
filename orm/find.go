package orm

import "fmt"

func FindDemo() {
	DB.DropTableIfExists(&Language{})
	DB.AutoMigrate(&Language{})

	language1 := Language{Language: "english", Level: 1}
	language2 := Language{Language: "中文", Level: 1}
	language3 := Language{Language: "lingua italiana", Level: 1}
	language4 := Language{Language: "Deutsch", Level: 1}
	DB.Create(&language1)
	DB.Create(&language2)
	DB.Create(&language3)
	DB.Create(&language4)

	res := Language{}
	DB.Model(&Language{}).Where("level = ?", 1).Find(&res)
	DB.Model(&Language{}).Where("level = ?", 1).Order("id desc").Find(&res)
	DB.Model(&Language{}).Where("level = ?", 1).Order("id").Find(&res)
	DB.Model(&Language{}).Where("level = ?", 1).Order("id desc").First(&res)
	DB.Model(&Language{}).Where("level = ?", 1).Order("id desc").Last(&res)
	fmt.Printf("%+v", res)
}
