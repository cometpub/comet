/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tv8jz3pyoo2ontw")

  // remove
  collection.schema.removeField("k1v1brdm")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "6ikzem5i",
    "name": "photos",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "an01031ytvklnd0",
      "cascadeDelete": false,
      "minSelect": 1,
      "maxSelect": 10,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tv8jz3pyoo2ontw")

  // add
  collection.schema.addField(new SchemaField({
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
      "thumbs": [
        "450x0",
        "900x0"
      ],
      "maxSelect": 10,
      "maxSize": 5242880,
      "protected": false
    }
  }))

  // remove
  collection.schema.removeField("6ikzem5i")

  return dao.saveCollection(collection)
})
