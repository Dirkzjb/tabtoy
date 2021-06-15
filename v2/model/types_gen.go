// Generated by github.com/Dirkzjb/tabtoy
// Version: 2.7.4
// DO NOT EDIT!!
package model

import (
	"encoding/json"
	"io/ioutil"
)

// Defined in table: Types
type ConditionType int32

const (

	// 无
	ConditionType_None ConditionType = 0

	// 值存在
	ConditionType_ValueExists ConditionType = 1

	// 值相等
	ConditionType_ValueEqual ConditionType = 2

	// 值范围
	ConditionType_ValueRange ConditionType = 3
)

// Defined in table: Builtin
type Builtin struct {
}

// Defined in table: Types
type TableTypes struct {

	//对象类型
	ObjectType string `MustFill:"true"`

	//字段名
	FieldName string `MustFill:"true"`

	//字段类型
	FieldType string `MustFill:"true"`

	//枚举值
	Value string

	//别名
	Alias string

	//默认值
	Default string

	//特性
	Meta string

	//注释
	Comment string
}

// Defined in table: Types
type TableVerify struct {

	//规则名
	RuleName string `MustFill:"true"`

	//条件
	Condition string `MustFill:"true"`

	//字段路径(类型名.字段名)
	FieldPath string `MustFill:"true"`

	//值
	Value string `MustFill:"true"`
}

// Defined in table: Types
type ValueRange struct {

	// 最小
	Min string

	// 最大
	Max string
}

// Builtin 访问接口
type BuiltinTable struct {

	// 表格原始数据
	Builtin

	// 索引函数表
	indexFuncByName map[string][]func(*BuiltinTable)

	// 清空函数表
	clearFuncByName map[string][]func(*BuiltinTable)
}

// 从json文件加载
func (self *BuiltinTable) Load(filename string) error {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	// 清除
	for _, list := range self.clearFuncByName {
		for _, v := range list {
			v(self)
		}
	}

	err = json.Unmarshal(data, &self.Builtin)
	if err != nil {
		return err
	}

	// 生成索引
	for _, list := range self.indexFuncByName {
		for _, v := range list {
			v(self)
		}
	}

	return nil
}

// 注册外部索引入口, 索引回调, 清空回调
func (self *BuiltinTable) RegisterIndexEntry(name string, indexCallback func(*BuiltinTable), clearCallback func(*BuiltinTable)) {

	indexList, _ := self.indexFuncByName[name]
	clearList, _ := self.clearFuncByName[name]

	if indexCallback != nil {
		indexList = append(indexList, indexCallback)
	}

	if clearCallback != nil {
		clearList = append(clearList, clearCallback)
	}

	self.indexFuncByName[name] = indexList
	self.clearFuncByName[name] = clearList
}

// 创建一个Builtin表读取实例
func NewBuiltinTable() *BuiltinTable {
	return &BuiltinTable{

		indexFuncByName: map[string][]func(*BuiltinTable){},

		clearFuncByName: map[string][]func(*BuiltinTable){},
	}
}
