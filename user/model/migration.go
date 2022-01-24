package model

func migration()  {
	// 自动迁移
	DB.AutoMigrate(&User{})
}
