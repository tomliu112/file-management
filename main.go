package main

import (
	"FileManagementDemo/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var DB *gorm.DB

func main() {


	db, err := gorm.Open(sqlite.Open("./files/mydatabase.db"), &gorm.Config{})

	//创建pv表
	db.AutoMigrate(&Pv{})
	//创建文件表
	db.AutoMigrate(&File{})
	db = db.Debug()
	DB=db

	if err != nil {
		// 处理错误
	}
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1/file"

	 fileManage := router.Group("/api/v1/file")
	 {
		fileManage.GET("/getList", getFileList)
		fileManage.POST("/upload", uploadFile)
		fileManage.GET("/download", downloadFile)
		fileManage.DELETE("/delete/:path", deleteFile)
		fileManage.GET("/getPv/:period", getPv)
	 }

	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

}



type Pv struct {
	ID   uint   `gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `gorm:"not null"`
	URI string `gorm:"not null"`
}

type File struct {
	gorm.Model
	Name string `gorm:"not null"`
}

//删除文件
// ShowAccount godoc
// @Summary      删除文件
// @Description  删除文件
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        path   path  string  true  "file_path"
// @Router       /delete/{path} [delete]
func deleteFile(context *gin.Context) {
	var file File
	path := context.Param("path")
	e := os.Remove("./files/"+path)
	if e != nil {
		context.JSON(200, gin.H{
			"msg": e.Error(),
		})
	}else{
		DB.Transaction(func(tx *gorm.DB) error {
			//存入文件表, 逻辑删除，设置deleted_at
			if err := tx.Where("name = ?", path).Delete(&file).Error; err != nil {
				fmt.Println(err.Error())
				return err
			}
			//存入pv表
			if err := tx.Create(&Pv{URI: "delete"}).Error; err != nil {
				fmt.Println(err.Error())
				return err
			}

			// 返回 nil 提交事务
			return nil
		})

		context.JSON(200, gin.H{
			"msg": "删除成功",
		})
	}


	
}

//下载文件
// ShowAccount godoc
// @Summary      下载文件
// @Description  下载文件
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        file_path   query  string  true  "file_path"
// @Router       /download [get]
func downloadFile(c *gin.Context) {
	filePath := c.Query("file_path")
	file, err := os.Open("./files/"+filePath)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Open file error: %s", err.Error()))
		return
	}
	defer file.Close()

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(filePath)))
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.Header("Pragma", "no-cache")

	_, err = io.Copy(c.Writer, file)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Copy file error: %s", err.Error()))
	}

	//存入pv表
	 DB.Create(&Pv{URI: "download"})

}


// @BasePath /api/v1
// 上传文件
// uploadFile godoc
// @Summary uploadFile
// @Schemes
// @Description 上传文件
// @Tags example
// @Accept mpfd
// @Produce json
// @Param   file    formData  file  true  "file to upload"  file
// @Success 200 {string} Helloworld
// @Router /upload [post]
func uploadFile(context *gin.Context) {
	f, _ := context.FormFile("file")
	context.SaveUploadedFile(f, "./files/"+f.Filename)
	context.JSON(200, gin.H{
		"msg": f.Filename+" 上传成功",
	})


	DB.Transaction(func(tx *gorm.DB) error {
		//存入文件表
		if err := tx.Create(&File{Name: f.Filename}).Error; err != nil {
			return err
		}
		//存入pv表
		if err := tx.Create(&Pv{URI: "upload",CreatedAt: time.Now()}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})

}

// @BasePath /api/v1
// 获取文件列表
// getFileList godoc
// @Summary getFileList
// @Schemes
// @Description 获取文件列表
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /getList [get]
func getFileList(context *gin.Context) {
	result := []map[string]interface{}{}
	DB.Table("files").Select("name" ).Where("deleted_at is  ?", nil).Find(&result)
	var pvs []Pv
	DB.Find(&pvs)
	context.JSON(200, gin.H{
		"filelists": result,
	})
}


// ShowAccount godoc
// @Summary      获取pv
// @Description  根据时间按天，月，年，获取pv
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        period   path  string  true  "day/month/year"
// @Router       /getPv/{period} [get]
func getPv(context *gin.Context) {
	period := context.Param("period")
	end:= time.Now()
	var start time.Time


	switch period {
	case "day":
		start=end.AddDate(0, 0, -1);
	case "month":
		start=end.AddDate(0, -1, 0);
	default:
		start=end.AddDate(-1, 0, 0);
	}
	var pvCount int64
	DB.Model(&Pv{}).Where("created_at >= ? AND created_at <= ?", start, end).Count(&pvCount)
	context.JSON(200, gin.H{
		"period": period ,
		"pvCount": pvCount,
	})

}
