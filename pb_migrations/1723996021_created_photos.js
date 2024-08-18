/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const collection = new Collection({
    "id": "an01031ytvklnd0",
    "created": "2024-08-18 15:47:01.062Z",
    "updated": "2024-08-18 15:47:01.062Z",
    "name": "photos",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "rvpv38fe",
        "name": "src",
        "type": "file",
        "required": true,
        "presentable": true,
        "unique": false,
        "options": {
          "mimeTypes": [
            "image/png",
            "image/jpeg"
          ],
          "thumbs": [
            "450x0",
            "900x0",
            "1440x0"
          ],
          "maxSelect": 1,
          "maxSize": 5242880,
          "protected": false
        }
      },
      {
        "system": false,
        "id": "fy5icszy",
        "name": "alt",
        "type": "text",
        "required": true,
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
        "id": "l9bdomcu",
        "name": "width",
        "type": "number",
        "required": true,
        "presentable": false,
        "unique": false,
        "options": {
          "min": 1,
          "max": null,
          "noDecimal": true
        }
      },
      {
        "system": false,
        "id": "3ltdiqfn",
        "name": "height",
        "type": "number",
        "required": true,
        "presentable": false,
        "unique": false,
        "options": {
          "min": 1,
          "max": null,
          "noDecimal": true
        }
      },
      {
        "system": false,
        "id": "m2ky7dgn",
        "name": "publication",
        "type": "relation",
        "required": true,
        "presentable": false,
        "unique": false,
        "options": {
          "collectionId": "js9kodn57zlyeua",
          "cascadeDelete": true,
          "minSelect": null,
          "maxSelect": 1,
          "displayFields": null
        }
      }
    ],
    "indexes": [
      "CREATE INDEX `idx_RoM9FY3` ON `photos` (`publication`)"
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("an01031ytvklnd0");

  return dao.deleteCollection(collection);
})
