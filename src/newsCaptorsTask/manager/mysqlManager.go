package manager

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"newsCaptorsTask/model"
	"newsCaptorsTask/config"
	"newsCaptorsTask/log"
)


func init() {
	log := log.MyLogger("mysq")

	log.Debug("-----------------mysql init called---------------")
	conf := config.AppConf()

	port := conf.GetValue("mysql","port")
	host := conf.GetValue("mysql","host")
	database := conf.GetValue("mysql","database")
	user := conf.GetValue("mysql","user")
	password := conf.GetValue("mysql","password")

	orm.RegisterDriver("mysql",orm.DRMySQL)
	url := user + ":" + password + "@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local"

	err := orm.RegisterDataBase("default","mysql",url)
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(model.CaptorInfo),new(model.Hobby))
	orm.RunCommand()

	//captor := new(model.CaptorInfo)
	////captor.Id = 2
	////captor.Url = "http://www.baidu.com"
	////captor.Createat = time.Now()
	////captor.Updateat = time.Now()
	////if err := captor.Insert();err != nil {
	////	panic(err)
	////}
	//var captors []*model.CaptorInfo
	//t,err := captor.Query().All(&captors)
	//fmt.Println(t,err)
	//
	//for _,v := range captors {
	//	fmt.Println((*v).ToJsonString())
	//}
}