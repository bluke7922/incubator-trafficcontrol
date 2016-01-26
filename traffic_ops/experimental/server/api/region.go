// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Region struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
	Links       RegionLinks `json:"_links" db:-`
}

type RegionLinks struct {
	Self         string       `db:"self" json:"_self"`
	DivisionLink DivisionLink `json:"division" db:-`
}

type RegionLink struct {
	ID  int64  `db:"region" json:"id"`
	Ref string `db:"region_id_ref" json:"_ref"`
}

// @Title getRegionById
// @Description retrieves the region information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Region
// @Resource /api/2.0
// @Router /api/2.0/region/{id} [get]
func getRegionById(id int, db *sqlx.DB) (interface{}, error) {
	ret := []Region{}
	arg := Region{}
	arg.Id = int64(id)
	queryStr := "select *, concat('" + API_PATH + "region/', id) as self "
	queryStr += ", concat('" + API_PATH + "division/', division) as division_id_ref"
	queryStr += " from region where id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getRegions
// @Description retrieves the region
// @Accept  application/json
// @Success 200 {array}    Region
// @Resource /api/2.0
// @Router /api/2.0/region [get]
func getRegions(db *sqlx.DB) (interface{}, error) {
	ret := []Region{}
	queryStr := "select *, concat('" + API_PATH + "region/', id) as self "
	queryStr += ", concat('" + API_PATH + "division/', division) as division_id_ref"
	queryStr += " from region"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postRegion
// @Description enter a new region
// @Accept  application/json
// @Param                 Body body     Region   true "Region object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/region [post]
func postRegion(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Region
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO region("
	sqlString += "name"
	sqlString += ",division"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:division"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putRegion
// @Description modify an existing regionentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Region   true "Region object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/region/{id}  [put]
func putRegion(id int, payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Region
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE region SET "
	sqlString += "name = :name"
	sqlString += ",division = :division"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delRegionById
// @Description deletes region information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Region
// @Resource /api/2.0
// @Router /api/2.0/region/{id} [delete]
func delRegion(id int, db *sqlx.DB) (interface{}, error) {
	arg := Region{}
	arg.Id = int64(id)
	result, err := db.NamedExec("DELETE FROM region WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
