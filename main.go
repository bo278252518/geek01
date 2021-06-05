package main

import (
	"flag"
	"fmt"

	"github.com/bo278252518/geek01/conf"
	"github.com/bo278252518/geek01/dao"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var confPath = flag.String("conf", "./config.yaml", "config file path")

func main() {
	flag.Parse()
	cfg, err := conf.NewConfig(*confPath)
	if err != nil {
		fmt.Printf("error is %+v\n", err)
		return
	}
	db, err := dao.NewMysql(cfg.Dsn)
	if err != nil {
		fmt.Printf("error is %+v\n", err)
		return
	}
	d, _ := dao.New(db)
	message, err := d.FindById(1)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("not found.")
			return
		}
		fmt.Printf("error is %+v\n", err)
		return
	}
	fmt.Println(message)
}
