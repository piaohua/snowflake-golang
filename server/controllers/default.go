package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	snowflake "github.com/piaohua/snowflake-golang"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	uri := c.Ctx.Request.URL.RequestURI()
	c.sendError(404, (uri + " Not Found"))
}

func (c *MainController) Post() {
	snowflake.DefaultNode()
	id := snowflake.Generate()
	c.sendID(id)
}

func (c *MainController) Count() {
	count := c.getUint32(":count")
	logs.Info("count: %d", count)
	var ids = make([]uint64, int(count))
	snowflake.DefaultNode()
	var i uint32
	for i = 0; i < count; i++ {
		id := snowflake.Generate()
		ids[i] = id
	}
	c.sendIDs(ids)
}

func (c *MainController) getUint32(key string) uint32 {
	value, err := c.GetUint32(key)
	if err != nil {
		c.sendErr(err)
	}
	return value
}

func (c *MainController) sendData(data interface{}) {
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = data
	c.ServeJSON()
	c.StopRun()
}

func (c *MainController) sendID(id uint64) {
	result := make(map[string]interface{}, 0)
	result["id"] = id
	c.sendData(result)
}

func (c *MainController) sendIDs(ids []uint64) {
	result := make(map[string]interface{}, 0)
	result["ids"] = ids
	c.sendData(result)
}

func (c *MainController) sendErr(err error) {
	c.sendError(200, err.Error())
}

func (c *MainController) sendError(code int, msg string) {
	logs.Error("error code: %d, msg: %s", code, msg)
	c.Ctx.Output.SetStatus(code)
	result := make(map[string]interface{}, 0)
	result["error"] = msg
	c.Data["json"] = result
	c.ServeJSON()
	c.StopRun()
}

// Prepare implemented Prepare() method for baseController.
func (c *MainController) Prepare() {
	token := beego.AppConfig.String("x-snowflake-access-token")
	if token != "snowflake" {
		//c.sendError(401, "Unauthorized")
	}
}

// init log
func init() {
	var beeLogger = logs.GetBeeLogger()
	beeLogger = logs.NewLogger(10000)
	filename := "snowflake.log"
	logConfig := `{"filename":"%s"}`
	logConfig = fmt.Sprintf(logConfig, filename)
	beeLogger.SetLogger(logs.AdapterFile, logConfig)
}
