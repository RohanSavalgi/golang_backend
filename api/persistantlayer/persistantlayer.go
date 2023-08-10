package persistantlayer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"application/datamodels"
	"application/db"
	"application/logger"

	// "application/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MongoDbInterface struct {
	mongoDb map[int]datamodels.Row
}

func InitiateMongoDbInterface() *MongoDbInterface {
	return &MongoDbInterface{ make(map[int]datamodels.Row) }
}

func (mdi MongoDbInterface) CreateRow(row datamodels.Row) error {
	if _, exists := mdi.mongoDb[row.ID]; exists {	
		return errors.New("The customer with this id already exists")
	}
	mdi.mongoDb[row.ID] = row
	return nil
}

func (mdi MongoDbInterface) GetAll() ([]datamodels.Row, error) {
	var allRows []datamodels.Row
	for _, row := range mdi.mongoDb {
		allRows = append(allRows, row)
	}
	return allRows, nil
}

func (mdi MongoDbInterface) UpdateRow(id int, updatedRow datamodels.Row) error {
	if _, exists := mdi.mongoDb[id]; exists {
		return errors.New("There is no row with this ID")
	}
	mdi.mongoDb[id] = updatedRow
	return nil
}

func (mdi MongoDbInterface) DeleteRow(id int) error {
	if _, exists := mdi.mongoDb[id]; exists {
		return errors.New("There is no row with this id")
	}
	delete(mdi.mongoDb, id)
	return nil
}

type GinInterfaceFunction struct {
	ginStore map[int]datamodels.Row
}

func InitiateGinInterface() *GinInterfaceFunction {
	object := GinInterfaceFunction{ make(map[int]datamodels.Row) }
	return &object
}

func (ginInterfaceFunctionObject GinInterfaceFunction) CreateRow(c *gin.Context) error {
	requestBody, requestError := ioutil.ReadAll(c.Request.Body)
	if requestError != nil {
		return errors.New("request could not be unpacked")
	}
	var creatorData datamodels.Row
	unpackedDataError := json.Unmarshal(requestBody, &creatorData)
	if unpackedDataError != nil {
		return errors.New("unmarshalling the data had a problem")
	}

	if _,exists := ginInterfaceFunctionObject.ginStore[creatorData.ID]; exists {
		fmt.Println("Returned from here")
		return errors.New("The User already exists with this ID")
	}

	ginInterfaceFunctionObject.ginStore[creatorData.ID] = creatorData
	// creationErr := db.Db.Create(&creatorData)
	// if creationErr != nil {
	// 	logger.ThrowErrorLog(creationErr)
	// }


	return nil
}

func (ginInterfaceFunctionObject GinInterfaceFunction) GetAll() ([]datamodels.Row, error) {
	// allUser := []datamodels.Row{}
	// for _, user := range ginInterfaceFunctionObject.ginStore {
	// 	allUser = append(allUser, user)
	// }
	// return allUser, nil

	allUser := []datamodels.Row{}
	var alterUser any
	db.Db.Table("myusers").Find(alterUser)
	fmt.Println(alterUser)
	return allUser, nil
}

func (ginInterfaceFunctionObject GinInterfaceFunction) UpdateRow(c *gin.Context) error {
	requestBody, requestConversionError := ioutil.ReadAll(c.Request.Body)
	if requestConversionError != nil {
		return errors.New("Request body error!")
	}

	var jsonData datamodels.Row
	unmarshallingError := json.Unmarshal(requestBody, &jsonData)
	if unmarshallingError != nil {
		return errors.New("Error in unmarshaling")
	}

	if _, exists := ginInterfaceFunctionObject.ginStore[jsonData.ID]; !exists {
		return errors.New("These is no user with this ID.")
	}
	
	ginInterfaceFunctionObject.ginStore[jsonData.ID] = jsonData
	return nil
}

func (ginInterfaceFunctionObject GinInterfaceFunction) DeleteRow(c *gin.Context) error {
	index, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(index)
	if _, exists := ginInterfaceFunctionObject.ginStore[index]; !exists {
		return errors.New("No user with this ID")
	}

	delete(ginInterfaceFunctionObject.ginStore, index)
	return nil
}

type PostgresInterface struct {
	postgresDb *gorm.DB
}

func PostgresInitilization() *PostgresInterface {
	pgInterfaceConnection := PostgresInterface{ postgresDb : db.Db }
	return &pgInterfaceConnection
}

func (pgi PostgresInterface) CreateRow(c *gin.Context) error {
	requestBody, requestError := ioutil.ReadAll(c.Request.Body)
	if requestError != nil {
		return errors.New("request could not be unpacked")
	}
	var creatorData datamodels.Row
	unpackedDataError := json.Unmarshal(requestBody, &creatorData)
	if unpackedDataError != nil {
		return errors.New("unmarshalling the data had a problem")
	}
	var user datamodels.Row
	pgi.postgresDb.Table("myusers").Find(&user, creatorData.ID)
	if user.ID > 0 {
		return errors.New("The User already exists with this ID")
	}

	creationErr := pgi.postgresDb.Table("myusers").Create(&creatorData)
	if creationErr.Error != nil {
		logger.ThrowErrorLog(creationErr)
		return errors.New("The user creation in postgres db was not successful.")
	}
	return nil
}

func (pgi PostgresInterface) GetAll() ([]datamodels.Row, error) {
	allUser := []datamodels.Row{}
	pgi.postgresDb.Table("myusers").Find(&allUser)
	fmt.Println(allUser)
	return allUser, nil
}

func (pgi PostgresInterface) UpdateRow(c *gin.Context) error {
	requestBody, requestConversionError := ioutil.ReadAll(c.Request.Body)
	if requestConversionError != nil {
		return errors.New("Request body error!")
	}

	var jsonData datamodels.Row
	unmarshallingError := json.Unmarshal(requestBody, &jsonData)
	if unmarshallingError != nil {
		return errors.New("Error in unmarshaling")
	}

	if exists := pgi.postgresDb.Table("myusers").Find(jsonData.ID); exists == nil {
		return errors.New("These is no user with this ID.")
	}
	
	updatingError := pgi.postgresDb.Table("myusers").Updates(jsonData)
	if updatingError.Error != nil {
		return errors.New("The updation at the database side was not successful.")
	}
	return nil
}

func (pgi PostgresInterface) DeleteRow(c *gin.Context) error {
	index, _ := strconv.Atoi(c.Param("id"))
	userForDeletion := datamodels.Row{ ID : index }
	if exists := pgi.postgresDb.Table("myusers").Find(userForDeletion); exists == nil {
		return errors.New("No user with this ID")
	}
	logger.ThrowDebugLog(userForDeletion)
	deletionLog := pgi.postgresDb.Table("myusers").Delete(&userForDeletion)
	if deletionLog != nil {
		return errors.New("There was some error while deleting the user in database")
	}
	return nil
}