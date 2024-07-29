package db

import (
	"fmt"

	"github.com/EmiliodDev/EmilioDev/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
    cfg := config.AppConfig.DB

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.Name,
    ) 
    
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
