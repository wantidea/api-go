package mongodb

import (
	"api-go/lib/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"strconv"
	"time"
)

type Mongodb struct {
	DatabaseName   string      // 数据库名称
	CollectionName string      // 集合名称
	Model          interface{} // 模型
}

// Client 返回客户端连接
func (m *Mongodb) Client() *mongo.Client {
	return client
}

// Database 返回数据库连接
func (m *Mongodb) Database() *mongo.Database {
	if m.DatabaseName == "" {
		return m.Client().Database(config.MongodbConfig.Database)
	}
	return m.Client().Database(m.DatabaseName)
}

// Collection 返回集合连接
func (m *Mongodb) Collection() *mongo.Collection {
	return m.Database().Collection(m.CollectionName)
}

// SetId 转换 ObjectID
func (m *Mongodb) SetId(id string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return oid, err
}

// Field 获取结构体列键值 bson
func (m *Mongodb) Field() map[string]interface{} {
	data := make(map[string]interface{})

	t := reflect.TypeOf(m.Model)
	v := reflect.ValueOf(m.Model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct {
		return data
	}

	for i := 0; i < t.NumField(); i++ {
		bsonName := t.Field(i).Tag.Get("bson")
		if bsonName == "" {
			continue
		}

		bsonValue := v.Field(i).Interface()
		data[bsonName] = bsonValue
	}

	return data
}

// AddField 过滤添加列
func (m *Mongodb) AddField() map[string]interface{} {
	data := m.Field()
	delete(data, "_id")
	delete(data, "created_at")
	delete(data, "updated_at")
	delete(data, "deleted_at")
	delete(data, "state")
	return data
}

// UpdField 过滤修改列
func (m *Mongodb) UpdField() map[string]interface{} {
	data := m.Field()
	delete(data, "created_at")
	delete(data, "updated_at")
	delete(data, "deleted_at")
	delete(data, "state")
	return data
}

// Add 添加
func (m *Mongodb) Add() (*mongo.InsertOneResult, error) {
	data := m.AddField()
	data["_id"] = primitive.NewObjectID()
	data["created_at"] = time.Now().Unix()
	data["updated_at"] = 0
	data["deleted_at"] = 0
	data["state"] = 1
	return m.Collection().InsertOne(context.TODO(), data)
}

// Del 删除
func (m *Mongodb) Del() (*mongo.UpdateResult, error) {
	data := m.Field()
	oid, err := m.SetId(data["_id"].(string))
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id":   oid,
		"state": 1,
	}

	data = bson.M{"$set": bson.M{
		"deleted_at": time.Now().Unix(),
		"state":      0,
	}}

	return m.Collection().UpdateOne(context.TODO(), filter, data)
}

// DelFilter 条件删除
func (m *Mongodb) DelFilter(filter bson.M) (*mongo.UpdateResult, error) {
	filter["state"] = 1
	data := bson.M{"$set": bson.M{
		"deleted_at": time.Now().Unix(),
		"state":      0,
	}}

	return m.Collection().UpdateMany(context.TODO(), filter, data)
}

// Upd 修改
func (m *Mongodb) Upd() (*mongo.UpdateResult, error) {
	data := m.UpdField()
	oid, err := m.SetId(data["_id"].(string))
	if err != nil {
		return nil, err
	}
	delete(data, "_id")

	filter := bson.M{
		"_id":   oid,
		"state": 1,
	}

	data["updated_at"] = time.Now().Unix()

	return m.Collection().UpdateOne(context.TODO(), filter, bson.M{"$set": data})
}

// UpdFilter 条件修改
func (m *Mongodb) UpdFilter(filter bson.M) (*mongo.UpdateResult, error) {
	filter["state"] = 1
	data := m.AddField()

	data["updated_at"] = time.Now().Unix()

	return m.Collection().UpdateOne(context.TODO(), filter, bson.M{"$set": data})
}

// Item 查询
func (m *Mongodb) Item() (map[string]interface{}, error) {
	data := m.Field()
	oid, err := m.SetId(data["_id"].(string))
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id":   oid,
		"state": 1,
	}

	result := data

	err = m.Collection().FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	result = m.Format(result)
	return result, err
}

// ItemFilter 条件查询
func (m *Mongodb) ItemFilter(filter bson.M) (map[string]interface{}, error) {
	filter["state"] = 1
	result := m.Field()

	err := m.Collection().FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	result = m.Format(result)
	return result, err
}

// List 列表
func (m *Mongodb) List() ([]*map[string]interface{}, error) {
	filter := bson.M{
		"state": 1,
	}
	cur, err := m.Collection().Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var resultList []*map[string]interface{}
	for cur.Next(context.TODO()) {
		result := m.Field()
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		result = m.Format(result)
		resultList = append(resultList, &result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	_ = cur.Close(context.TODO())
	return resultList, nil
}

// ListFilter 条件列表
func (m *Mongodb) ListFilter(filter primitive.M, opts ...*options.FindOptions) ([]*map[string]interface{}, error) {
	filter["state"] = 1
	cur, err := m.Collection().Find(context.TODO(), filter, opts...)
	if err != nil {
		return nil, err
	}

	var resultList []*map[string]interface{}
	for cur.Next(context.TODO()) {
		result := m.Field()
		err = cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		result = m.Format(result)
		resultList = append(resultList, &result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	_ = cur.Close(context.TODO())
	return resultList, nil
}

// Count 统计条数
func (m *Mongodb) Count(filter primitive.M) (int64, error) {
	filter["state"] = 1
	return m.Collection().CountDocuments(context.TODO(), filter)
}

// 转换字段类型
func (m *Mongodb) Format(data map[string]interface{}) map[string]interface{} {
	filedList := make(map[string]string)
	t := reflect.TypeOf(m.Model).Elem()
	for i := 0; i < t.NumField(); i++ {
		bsonName := t.Field(i).Tag.Get("bson")
		if bsonName == "" {
			continue
		}

		var bsonType string
		if bsonName == "_id" {
			bsonType = "primitive.ObjectID"
		} else {
			bsonType = fmt.Sprintf("%v", t.Field(i).Type)

		}
		filedList[bsonName] = bsonType
	}

	for key, value := range data {
		valueType := filedList[key]
		if valueType == "" {
			continue
		}
		switch valueType {
		case "primitive.ObjectID":
			id := fmt.Sprintf("%v", value)
			data[key] = id[10 : len(id)-2]
		case "string":
			data[key] = value
		case "int":
			data[key], _ = strconv.Atoi(fmt.Sprintf("%v", value))
		case "int32":
			data[key], _ = strconv.ParseInt(fmt.Sprintf("%v", value), 10, 32)
		case "int64":
			data[key], _ = strconv.ParseInt(fmt.Sprintf("%v", value), 10, 32)
		default:
			data[key] = value
		}
	}
	return data
}
