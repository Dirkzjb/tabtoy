package i18n

import (
	"fmt"
)

type StringID int

const (
	ConvertValue_EnumTypeNil StringID = iota
	ConvertValue_StructTypeNil
	ConvertValue_EnumValueNotFound
	ConvertValue_UnknownFieldType
	StructParser_LexerError
	StructParser_ExpectField
	StructParser_UnexpectedSpliter
	StructParser_FieldNotFound
	StructParser_DuplicateFieldInCell
	Run_CacheFile
	Run_CollectTypeInfo
	Run_ExportSheetData
	Globals_CombineNameLost
	Globals_PackageNameDiff
	Globals_TableNameDuplicated
	Globals_OutputCombineData
	Globals_DuplicateTypeName
	File_TypeSheetKeepSingleton
	File_TypeSheetNotFound
	DataSheet_ValueConvertError
	DataSheet_ValueRepeated
	DataSheet_RowDataSplitedByEmptyLine
	DataSheet_MustFill
	DataHeader_TypeNotFound
	DataHeader_MetaParseFailed
	DataHeader_DuplicateFieldName
	DataHeader_RepeatedFieldTypeNotSameInMultiColumn
	DataHeader_RepeatedFieldMetaNotSameInMultiColumn
	DataHeader_UseReservedTypeName
	DataHeader_NotMatch
	DataHeader_FieldNotDefinedInMainTableInMultiTableMode
	DataHeader_NotMatchInMultiTableMode
	DataHeader_FieldPermParseFailed
	TypeSheet_PragmaParseFailed
	TypeSheet_TableNameIsEmpty
	TypeSheet_PackageIsEmpty
	TypeSheet_FieldTypeNotFound
	TypeSheet_EnumValueParseFailed
	TypeSheet_DescriptorKindNotSame
	TypeSheet_FieldMetaParseFailed
	TypeSheet_StructFieldCanNotBeStruct
	TypeSheet_FirstEnumValueShouldBeZero
	TypeSheet_UnexpectedTypeHeader
	TypeSheet_DuplicatedEnumValue
	TypeSheet_RowDataSplitedByEmptyLine
	TypeSheet_ObjectNameEmpty
	TypeSheet_FieldInvalid
	Printer_IgnoredByOutputTag
	Printer_OpenWriteOutputFileFailed
	Printer_IgnoredByTablePerm
	System_OpenReadXlsxFailed
)

var currLan map[StringID]string

var lanByStr = make(map[string]map[StringID]string)

func String(id StringID) string {

	if currLan == nil {
		return "!!i18n not set!!"
	}

	if str, ok := currLan[id]; ok {
		return str
	}

	return fmt.Sprintf("i18n:%v", id)
}

func SetLanguage(lan string) bool {
	if v, ok := lanByStr[lan]; ok {
		currLan = v
	} else {
		return false
	}

	return true
}

func registerLanguage(lan string, data map[StringID]string) {
	lanByStr[lan] = data
}
