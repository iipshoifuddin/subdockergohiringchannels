package controllers

import (
	"fmt"
	"gohiringchannels/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

type Engineer struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Skill       string    `json:"skill"`
	Address     string    `json:"address"`
	DateOfBirth string    `json:"dateofbirth"`
	PhoneNumber string    `json:"phonenumber"`
	Showcase    string    `json:"showcase"`
	Email       string    `json:"email"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Enterprice struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	PhoneNumber string    `json:"phonenumber"`
	Email       string    `json:"email"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type CreateEngineerInput struct {
	UUID     string `form:"uuid"`
	Name     string `form:"name" binding:"required,general"`
	Email    string `form:"email" binding:"required,email"`
	Username string `form:"username" binding:"required,usernm"`
	Password string `form:"password" binding:"required,passwd"`
}

type UpdateEngineerInput struct {
	Name        string `form:"name" binding:"required,general"`
	Description string `form:"description" binding:"required,general"`
	Skill       string `form:"skill" binding:"required,general"`
	Address     string `form:"address" binding:"required,general"`
	DateOfBirth string `form:"dateofbirth" binding:"required,general"`
	PhoneNumber string `form:"phonenumber" binding:"required,general"`
}

type Image struct {
	Path        string `validate:"required"`
	Filename    string `validate:"required"`
	Ext         string `validate:"required"`
	ContentType string `validate:"required"`
	Bytes       int32  `validate:"required,gt=0"`
}

type Sizer interface {
	Size() int64
}

type UploadImageEngineer struct {
	Showcase string `form:"showcase" binding:"required,imagefile"`
}

var listViewEngineers = []string{
	"id",
	"uuid",
	"name",
	"description",
	"skill",
	"address",
	"date_of_birth",
	"phone_number",
	"showcase",
	"email",
}

var listViewEnterprices = []string{
	"id",
	"uuid",
	"name",
	"description",
	"address",
	"phone_number",
	"logo",
	"email",
	"updated_at",
}

type Booking struct {
	// CheckIn  time.Time `form:"check_in"`
	// CheckOut time.Time `form:"check_out"`

	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func FineEngineers(c *gin.Context) {
	var enginners []Engineer
	db := c.MustGet("db").(*gorm.DB)

	s := c.Query("search")
	if s == "" {
		s = "%%"
	}
	// s := "%yul%"
	fmt.Println(s)
	setOrder := c.Query("order")
	if setOrder == "" {
		setOrder = "name"
	}

	setOrderBy := c.Query("orderBy")
	if setOrder == "" {
		setOrderBy = "ASC"
	}

	order := fmt.Sprintf("%s %s", setOrder, setOrderBy)

	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	db.Offset(offset).Limit(pageSize)
	db.Select(listViewEngineers).Offset(offset).Limit(pageSize).Where("name LIKE ? OR skill LIKE ?", "%"+s+"%", "%"+s+"%").Order(order).Find(&enginners)

	c.JSON(http.StatusOK, gin.H{"data": enginners})
}

func FineEngineer(c *gin.Context) {
	var enginners []Engineer
	db := c.MustGet("db").(*gorm.DB)

	db.Select(listViewEngineers).Where("id = ?", c.Param("id")).Find(&enginners)
	c.JSON(http.StatusOK, gin.H{"data": enginners})
}

func EngineersLookEnterprices(c *gin.Context) {
	var enterprices []Enterprice
	db := c.MustGet("db").(*gorm.DB)

	s := c.Query("search")
	if s == "" {
		s = "%%"
	}
	// s := "%yul%"
	fmt.Println(s)
	setOrder := c.Query("order")
	if setOrder == "" {
		setOrder = "name"
	}

	setOrderBy := c.Query("orderBy")
	if setOrder == "" {
		setOrderBy = "ASC"
	}

	order := fmt.Sprintf("%s %s", setOrder, setOrderBy)

	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	db.Offset(offset).Limit(pageSize)
	db.Select(listViewEnterprices).Offset(offset).Limit(pageSize).Where("name LIKE ? ", "%"+s+"%").Order(order).Find(&enterprices)

	c.JSON(http.StatusOK, gin.H{"data": enterprices})
}

func EngineerLookEnterpricesDetail(c *gin.Context) {
	var enterprice []Enterprice
	db := c.MustGet("db").(*gorm.DB)

	db.Select(listViewEnterprices).Where("id = ?", c.Param("id")).Find(&enterprice)
	c.JSON(http.StatusOK, gin.H{"data": enterprice})
}

func CreateEngineer(c *gin.Context) {
	var input CreateEngineerInput
	hashUuid := uuid.NewV4().String()
	if err := c.ShouldBindWith(&input, binding.Form); err == nil {
		engineerCreate := models.Engineer{
			UUID:     hashUuid,
			Name:     input.Name,
			Email:    input.Email,
			Username: input.Username,
			Password: input.Password,
		}
		db := c.MustGet("db").(*gorm.DB)
		if result := db.Where("email = ? AND username = ?", input.Email, input.Username).First(&engineerCreate); result.Error != nil {
			db.Create(&engineerCreate)
			c.JSON(http.StatusOK, gin.H{"data": "data succsesfully created"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Username is Exist !"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func UpdateEngginer(c *gin.Context) {
	var input UpdateEngineerInput
	if err := c.ShouldBindWith(&input, binding.Form); err == nil {
		var enginners models.Engineer
		engineerUpdate := models.Engineer{
			Name:        input.Name,
			Description: input.Description,
			Skill:       input.Skill,
			Address:     input.Address,
			DateOfBirth: input.DateOfBirth,
			PhoneNumber: input.PhoneNumber,
		}
		db := c.MustGet("db").(*gorm.DB)
		if result := db.Where("id = ?", c.Param("id")).First(&enginners); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found in database !"})
		} else {
			db.Model(&enginners).Updates(engineerUpdate)
			c.JSON(http.StatusOK, gin.H{"data": "data succsesfully Updated !"})

		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func UploadShowcaseEngginer(c *gin.Context) {
	const SR_File_Max_Bytes = 1 * 1024 * 1024
	imgeName := uuid.NewV4().String()

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	if header.Size > SR_File_Max_Bytes {
		c.JSON(http.StatusBadRequest, gin.H{"error": " File size to big max 1024 kb !"})
		return
	}

	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	fileType := http.DetectContentType(fileHeader)
	if fileType == "image/jpeg" || fileType == "image/png" || fileType == "image/jpg" {
		// Set Folder untuk menyimpan filenya
		nameoffile := imgeName + header.Filename
		path := os.Getenv("PATH_UPLOAD_ENG") + nameoffile
		if err := c.SaveUploadedFile(header, path); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}
		var enginners models.Engineer
		engineerUpdate := models.Engineer{
			Showcase: os.Getenv("PUBLIC_UPLOAD_DB") + nameoffile,
		}
		db := c.MustGet("db").(*gorm.DB)
		if result := db.Where("id = ?", c.Param("id")).First(&enginners); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Not Found in database !"})
		} else {
			db.Model(&enginners).Updates(engineerUpdate)
			c.JSON(http.StatusOK, gin.H{"data": "File succsesfully Uploaded !"})

		}
		// Response
		// c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("File : %s successfully Uploaded !", header.Filename)})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": "file type is error !"})
	}
}

func DeleteAccount(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var engineer models.Engineer
	if err := db.Where("id = ?", c.Param("id")).First(&engineer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&engineer)

	c.JSON(http.StatusOK, gin.H{"data": "Account Has been deleted !"})
}
