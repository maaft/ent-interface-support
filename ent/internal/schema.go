// Code generated by entc, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"backend/ent/schema","Package":"backend/ent","Schemas":[{"name":"Car","config":{"Table":""},"fields":[{"name":"id","type":{"Type":7,"Ident":"pulid.ID","PkgPath":"backend/ent/schema/pulid","PkgName":"","Nillable":false,"RType":{"Name":"ID","Ident":"pulid.ID","Kind":24,"PkgPath":"backend/ent/schema/pulid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"createdAt","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"updatedAt","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":1}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":2}},{"name":"wheelPressure","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"annotations":{"PULID":{"Prefix":"Car"}}},{"name":"Plane","config":{"Table":""},"fields":[{"name":"id","type":{"Type":7,"Ident":"pulid.ID","PkgPath":"backend/ent/schema/pulid","PkgName":"","Nillable":false,"RType":{"Name":"ID","Ident":"pulid.ID","Kind":24,"PkgPath":"backend/ent/schema/pulid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"Value","Ident":"driver.Value","Kind":20,"PkgPath":"database/sql/driver","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"createdAt","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"updatedAt","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":1}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":2}},{"name":"altitude","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"annotations":{"PULID":{"Prefix":"Plane"}}}],"Features":["privacy","entql","schema/snapshot"]}`