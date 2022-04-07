package v1

import (
	"gin/blog/models"
	setting "gin/blog/pkg"
	"gin/blog/util"
	"github.com/astaxie/beego/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	// c.Query可用于获取?name=test&state=1这类URL参数
	// c *gin.Context是Gin很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的JSON和呈现JSON响应
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := setting.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  setting.GetMsg(code),
		"data": data,
	})
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	var tag models.Tag
	code := setting.InvalidParams
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  setting.GetMsg(code),
			"data": make(map[string]string),
		})
		return
	}
	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Required(tag.CreatedBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(tag.CreatedBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	if !valid.HasErrors() {
		if !models.ExistTagByName(tag.Name) {
			code = setting.SUCCESS
			models.AddTag(tag.Name, tag.State, tag.CreatedBy)
		} else {
			code = setting.ErrorExistTag
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  setting.GetMsg(code),
		"data": make(map[string]string),
	})
}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := setting.InvalidParams
	if !valid.HasErrors() {
		code = setting.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			models.EditTag(id, data)
		} else {
			code = setting.ErrorNotExistTag
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  setting.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := setting.InvalidParams
	if !valid.HasErrors() {
		code = setting.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = setting.ErrorNotExistTag
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  setting.GetMsg(code),
		"data": make(map[string]string),
	})
}
