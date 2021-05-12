// Code generated by rpc_gen. DO NOT EDIT.

package cmltestdb

import (
	"reflect"

	"git.pinquest.cn/qlb/brick/dbproxy"
)

var ModelUserRole = &dbproxy.ModelObjectType{
	Name: "cmltest.ModelUserRole",
	PrimaryKey: "id",
	PrimaryKeyType: "uint64",
	FieldList: &dbproxy.ObjectFieldList{
		List: []*dbproxy.ObjectField{
			{
				FieldName: "id",
				Type:      "uint64",
				RefType:   "",
				FormName:  "",
				FormType:  "",
				Validate:  "",
				DbDef:     "",
				ScopeCond: "",
				LinkTo:    "",
				Editable:  false,
				IsArray:   false,
			},
			{
				FieldName: "created_at",
				Type:      "uint32",
				RefType:   "",
				FormName:  "",
				FormType:  "",
				Validate:  "",
				DbDef:     "",
				ScopeCond: "",
				LinkTo:    "",
				Editable:  false,
				IsArray:   false,
			},
			{
				FieldName: "updated_at",
				Type:      "uint32",
				RefType:   "",
				FormName:  "",
				FormType:  "",
				Validate:  "",
				DbDef:     "",
				ScopeCond: "",
				LinkTo:    "",
				Editable:  false,
				IsArray:   false,
			},
			{
				FieldName: "deleted_at",
				Type:      "uint32",
				RefType:   "",
				FormName:  "",
				FormType:  "",
				Validate:  "",
				DbDef:     "",
				ScopeCond: "",
				LinkTo:    "",
				Editable:  false,
				IsArray:   false,
			},
			{
				FieldName: "corp_id",
				Type:      "uint32",
				RefType:   "",
				FormName:  "",
				FormType:  "",
				Validate:  "",
				DbDef:     "",
				ScopeCond: "",
				LinkTo:    "",
				Editable:  true,
				IsArray:   false,
			},
			{
				FieldName: "app_id",
				Type:      "uint32",
				RefType:   "",
				FormName:  "",
				FormType:  "",
				Validate:  "",
				DbDef:     "",
				ScopeCond: "",
				LinkTo:    "",
				Editable:  true,
				IsArray:   false,
			},
			{
				FieldName: "uid",
				Type:      "uint64",
				RefType:   "",
				FormName:  "",
				FormType:  "",
				Validate:  "",
				DbDef:     "",
				ScopeCond: "",
				LinkTo:    "",
				Editable:  true,
				IsArray:   false,
			},
			{
				FieldName: "role_id",
				Type:      "uint64",
				RefType:   "",
				FormName:  "",
				FormType:  "",
				Validate:  "",
				DbDef:     "",
				ScopeCond: "",
				LinkTo:    "",
				Editable:  true,
				IsArray:   false,
			},
		},
	},
}

func init() {
	ObjectTypeList = append(ObjectTypeList, ModelUserRole)
}

var IdModelUserRole = &dbproxy.StructField{
	StructFieldName: "Id",
	DbFieldName:     "id",
}

var CreatedAtModelUserRole = &dbproxy.StructField{
	StructFieldName: "CreatedAt",
	DbFieldName:     "created_at",
}

var UpdatedAtModelUserRole = &dbproxy.StructField{
	StructFieldName: "UpdatedAt",
	DbFieldName:     "updated_at",
}

var DeletedAtModelUserRole = &dbproxy.StructField{
	StructFieldName: "DeletedAt",
	DbFieldName:     "deleted_at",
}

var CorpIdModelUserRole = &dbproxy.StructField{
	StructFieldName: "CorpId",
	DbFieldName:     "corp_id",
}

var AppIdModelUserRole = &dbproxy.StructField{
	StructFieldName: "AppId",
	DbFieldName:     "app_id",
}

var UidModelUserRole = &dbproxy.StructField{
	StructFieldName: "Uid",
	DbFieldName:     "uid",
}

var RoleIdModelUserRole = &dbproxy.StructField{
	StructFieldName: "RoleId",
	DbFieldName:     "role_id",
}

type ModelUserRoleReflect struct {
	Id dbproxy.StructFieldWithObj
	CreatedAt dbproxy.StructFieldWithObj
	UpdatedAt dbproxy.StructFieldWithObj
	DeletedAt dbproxy.StructFieldWithObj
	CorpId dbproxy.StructFieldWithObj
	AppId dbproxy.StructFieldWithObj
	Uid dbproxy.StructFieldWithObj
	RoleId dbproxy.StructFieldWithObj
}

func NewModelUserRoleReflect(i interface{}) *ModelUserRoleReflect {
	obj := reflect.ValueOf(i)
	if obj.Kind() == reflect.Ptr {
	    obj = obj.Elem()
	}
	r := &ModelUserRoleReflect{
		Id: dbproxy.StructFieldWithObj{Field: IdModelUserRole, Obj: obj},
		CreatedAt: dbproxy.StructFieldWithObj{Field: CreatedAtModelUserRole, Obj: obj},
		UpdatedAt: dbproxy.StructFieldWithObj{Field: UpdatedAtModelUserRole, Obj: obj},
		DeletedAt: dbproxy.StructFieldWithObj{Field: DeletedAtModelUserRole, Obj: obj},
		CorpId: dbproxy.StructFieldWithObj{Field: CorpIdModelUserRole, Obj: obj},
		AppId: dbproxy.StructFieldWithObj{Field: AppIdModelUserRole, Obj: obj},
		Uid: dbproxy.StructFieldWithObj{Field: UidModelUserRole, Obj: obj},
		RoleId: dbproxy.StructFieldWithObj{Field: RoleIdModelUserRole, Obj: obj},
	}

	return r
}
