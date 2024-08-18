/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tv8jz3pyoo2ontw")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ekevtaxp",
    "name": "featuredImage",
    "type": "file",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "mimeTypes": [
        "image/png",
        "image/jpeg"
      ],
      "thumbs": [
        "1200x675"
      ],
      "maxSelect": 1,
      "maxSize": 5242880,
      "protected": false
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tv8jz3pyoo2ontw")

  // remove
  collection.schema.removeField("ekevtaxp")

  return dao.saveCollection(collection)
})
