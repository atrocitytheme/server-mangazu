//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Reference struct {
	GalleryUUID  string `sql:"primary_key"`
	MetaInternal bool
	MetaPath     *string
	MetaMatch    *int32
	Urls         *string
	ExhGid       *int32
	ExhToken     *string
	AnilistID    *int32
}
