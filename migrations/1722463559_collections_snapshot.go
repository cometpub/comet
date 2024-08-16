package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `[
			{
				"id": "_pb_users_auth_",
				"name": "users",
				"type": "auth",
				"system": false,
				"schema": [
				{
					"system": false,
					"id": "users_name",
					"name": "name",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": ""
					}
				},
				{
					"system": false,
					"id": "users_avatar",
					"name": "avatar",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"mimeTypes": [
						"image/jpeg",
						"image/png",
						"image/svg+xml",
						"image/gif",
						"image/webp"
					],
					"thumbs": [],
					"maxSelect": 1,
					"maxSize": 5242880,
					"protected": false
					}
				}
				],
				"indexes": [],
				"listRule": "id = @request.auth.id",
				"viewRule": "id = @request.auth.id",
				"createRule": "",
				"updateRule": "id = @request.auth.id",
				"deleteRule": "id = @request.auth.id",
				"options": {
				"allowEmailAuth": true,
				"allowOAuth2Auth": true,
				"allowUsernameAuth": true,
				"exceptEmailDomains": null,
				"manageRule": null,
				"minPasswordLength": 8,
				"onlyEmailDomains": null,
				"onlyVerified": false,
				"requireEmail": false
				}
			},
			{
				"id": "rbhhb9mhlj5f6e6",
				"name": "categories",
				"type": "base",
				"system": false,
				"schema": [
					{
						"system": false,
						"id": "m0be6sjf",
						"name": "slug",
						"type": "text",
						"required": true,
						"presentable": false,
						"unique": false,
						"options": {
						"min": null,
						"max": null,
						"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
						}
					}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_9yDmoQz` + "`" + ` ON ` + "`" + `categories` + "`" + ` (` + "`" + `slug` + "`" + `)"
				],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {}
			},
			{
				"id": "tv8jz3pyoo2ontw",
				"name": "entries",
				"type": "base",
				"system": false,
				"schema": [
				{
					"system": false,
					"id": "gtaafroo",
					"name": "name",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": ""
					}
				},
				{
					"system": false,
					"id": "ku2uwvfw",
					"name": "slug",
					"type": "text",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
					}
				},
				{
					"system": false,
					"id": "2yqdbca8",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"maxSelect": 1,
					"values": ["article", "note", "photo"]
					}
				},
				{
					"system": false,
					"id": "rlugtnue",
					"name": "summary",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": ""
					}
				},
				{
					"system": false,
					"id": "jcsmrc3h",
					"name": "content",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": ""
					}
				},
				{
					"system": false,
					"id": "vx13qjjj",
					"name": "published",
					"type": "date",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": "",
					"max": ""
					}
				},
				{
					"system": false,
					"id": "oqqdlnlx",
					"name": "authors",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"collectionId": "_pb_users_auth_",
					"cascadeDelete": false,
					"minSelect": null,
					"maxSelect": null,
					"displayFields": null
					}
				},
				{
					"system": false,
					"id": "t6kuqqmf",
					"name": "categories",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"collectionId": "rbhhb9mhlj5f6e6",
					"cascadeDelete": false,
					"minSelect": null,
					"maxSelect": null,
					"displayFields": null
					}
				},
				{
					"system": false,
					"id": "k1v1brdm",
					"name": "photos",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"mimeTypes": [
						"image/png",
						"image/jpeg",
						"image/gif",
						"image/webp",
						"image/avif"
					],
					"thumbs": [],
					"maxSelect": 99,
					"maxSize": 5242880,
					"protected": false
					}
				},
				{
					"system": false,
					"id": "vughvabz",
					"name": "bookmarkOf",
					"type": "url",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"exceptDomains": null,
					"onlyDomains": null
					}
				},
				{
					"system": false,
					"id": "hf8evx3s",
					"name": "audioUrl",
					"type": "url",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"exceptDomains": null,
					"onlyDomains": null
					}
				},
				{
					"system": false,
					"id": "odprcccn",
					"name": "videoUrl",
					"type": "url",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"exceptDomains": null,
					"onlyDomains": null
					}
				},
				{
					"system": false,
					"id": "t0crklkc",
					"name": "publication",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"collectionId": "js9kodn57zlyeua",
					"cascadeDelete": false,
					"minSelect": null,
					"maxSelect": 1,
					"displayFields": null
					}
				}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_EteF5vw` + "`" + ` ON ` + "`" + `entries` + "`" + ` (\n  ` + "`" + `publication` + "`" + `,\n  ` + "`" + `type` + "`" + `,\n  ` + "`" + `slug` + "`" + `\n)"
				],
				"listRule": "published != \"\"",
				"viewRule": "published != \"\"",
				"createRule": "@request.auth.id = authors.id",
				"updateRule": "@request.auth.id = authors.id",
				"deleteRule": "@request.auth.id = authors.id",
				"options": {}
			},
			{
				"id": "js9kodn57zlyeua",
				"name": "publications",
				"type": "base",
				"system": false,
				"schema": [
				{
					"system": false,
					"id": "xewkcyis",
					"name": "slug",
					"type": "text",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
					}
				},
				{
					"system": false,
					"id": "2mz4aowz",
					"name": "title",
					"type": "text",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": ""
					}
				},
				{
					"system": false,
					"id": "usxiyvvi",
					"name": "subtitle",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": ""
					}
				},
				{
					"system": false,
					"id": "iifss1m5",
					"name": "domain",
					"type": "url",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"exceptDomains": [],
					"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "hcjach4z",
					"name": "authors",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"collectionId": "_pb_users_auth_",
					"cascadeDelete": false,
					"minSelect": 1,
					"maxSelect": null,
					"displayFields": null
					}
				},
				{
					"system": false,
					"id": "sdqfmyba",
					"name": "icon",
					"type": "file",
					"required": true,
					"presentable": true,
					"unique": false,
					"options": {
					"mimeTypes": [
						"image/png",
						"image/jpeg",
						"image/gif",
						"image/svg+xml"
					],
					"thumbs": [],
					"maxSelect": 1,
					"maxSize": 5242880,
					"protected": false
					}
				},
				{
					"system": false,
					"id": "6czc0kvu",
					"name": "logo",
					"type": "file",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"mimeTypes": [
						"image/png",
						"image/jpeg",
						"image/gif",
						"image/svg+xml"
					],
					"thumbs": [],
					"maxSelect": 1,
					"maxSize": 5242880,
					"protected": false
					}
				}
				],
				"indexes": [
					"CREATE UNIQUE INDEX ` + "`" + `idx_0LJZc47` + "`" + ` ON ` + "`" + `publications` + "`" + ` (` + "`" + `slug` + "`" + `)",
					"CREATE UNIQUE INDEX ` + "`" + `idx_VN07dil` + "`" + ` ON ` + "`" + `publications` + "`" + ` (` + "`" + `domain` + "`" + `)"
				],
				"listRule": "@request.auth.id = authors.id",
				"viewRule": "@request.auth.id = authors.id",
				"createRule": "@request.auth.id = authors.id",
				"updateRule": "@request.auth.id = authors.id",
				"deleteRule": "@request.auth.id = authors.id",
				"options": {}
			},
			{
				"id": "885ryta7umsl6ue",
				"name": "publication_categories",
				"type": "view",
				"system": false,
				"schema": [
				{
					"system": false,
					"id": "3wtcqnkr",
					"name": "publication",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"collectionId": "js9kodn57zlyeua",
					"cascadeDelete": false,
					"minSelect": null,
					"maxSelect": 1,
					"displayFields": null
					}
				},
				{
					"system": false,
					"id": "pnvgf1ri",
					"name": "domain",
					"type": "url",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"exceptDomains": [],
					"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "x6vo3gau",
					"name": "categories",
					"type": "json",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"maxSize": 1
					}
				},
				{
					"system": false,
					"id": "hufm96bs",
					"name": "count",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"noDecimal": false
					}
				}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
				"query": "SELECT \n    (ROW_NUMBER() OVER()) as id,\n    p.id as publication,\n    p.domain as domain,\n    json_group_array(DISTINCT c.slug) as categories,\n    count(DISTINCT c.slug) as count\nFROM entries e\nLEFT JOIN publications p ON e.publication = p.id\nLEFT JOIN json_each(e.categories) as ec\nLEFT JOIN categories c ON ec.value = c.id\nWHERE e.categories <> \"[]\"\nGROUP BY p.id"
				}
			},
			{
				"id": "sqqksbztgydkbhw",
				"name": "publication_category_entries_count",
				"type": "view",
				"system": false,
				"schema": [
				{
					"system": false,
					"id": "vioeagfm",
					"name": "publication",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"collectionId": "js9kodn57zlyeua",
					"cascadeDelete": false,
					"minSelect": null,
					"maxSelect": 1,
					"displayFields": null
					}
				},
				{
					"system": false,
					"id": "5ffojgua",
					"name": "domain",
					"type": "url",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"exceptDomains": [],
					"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "kpnkcukb",
					"name": "slug",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"pattern": "^[a-z0-9]+(?:-[a-z0-9]+)*$"
					}
				},
				{
					"system": false,
					"id": "djq2iaq2",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"maxSelect": 1,
					"values": ["article", "note", "photo"]
					}
				},
				{
					"system": false,
					"id": "iziu9qp5",
					"name": "items",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"noDecimal": false
					}
				}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
				"query": "SELECT \n    (ROW_NUMBER() OVER()) as id,\n    p.id as publication,\n    p.domain as domain,\n    c.slug as slug,\n    e.type as type,\n    count(DISTINCT e.id) as items\nFROM entries e\nLEFT JOIN publications p ON e.publication = p.id\nLEFT JOIN json_each(e.categories) as ec\nLEFT JOIN categories c ON ec.value = c.id\nWHERE e.categories <> \"[]\"\nGROUP BY e.publication, e.type, c.slug"
				}
			},
			{
				"id": "4sqq8whsdzdlxpw",
				"name": "publication_entries_count",
				"type": "view",
				"system": false,
				"schema": [
				{
					"system": false,
					"id": "uezw2fdn",
					"name": "publication",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"collectionId": "js9kodn57zlyeua",
					"cascadeDelete": false,
					"minSelect": null,
					"maxSelect": 1,
					"displayFields": null
					}
				},
				{
					"system": false,
					"id": "nzijk7il",
					"name": "domain",
					"type": "url",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"exceptDomains": [],
					"onlyDomains": []
					}
				},
				{
					"system": false,
					"id": "pxxr2ucw",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
					"maxSelect": 1,
					"values": ["article", "note", "photo"]
					}
				},
				{
					"system": false,
					"id": "fs0bd4ji",
					"name": "items",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
					"min": null,
					"max": null,
					"noDecimal": false
					}
				}
				],
				"indexes": [],
				"listRule": null,
				"viewRule": null,
				"createRule": null,
				"updateRule": null,
				"deleteRule": null,
				"options": {
				"query": "SELECT \n    (ROW_NUMBER() OVER()) as id,\n    p.id as publication,\n    p.domain as domain,\n    e.type as type,\n    count(DISTINCT e.id) as items\nFROM entries e\nLEFT JOIN publications p ON e.publication = p.id\nGROUP BY e.publication, e.type"
				}
			}
		]`

		collections := []*models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return daos.New(db).ImportCollections(collections, true, nil)
	}, func(db dbx.Builder) error {
		return nil
	})
}
