package mongoconverter

import "go.mongodb.org/mongo-driver/bson"

func BsonDToMap(d bson.D) map[string]any {
	m := map[string]any{}
	for _, elem := range d {
		if bsonElem, ok := elem.Value.(bson.D); ok {
			m[elem.Key] = BsonDToMap(bsonElem)
		} else {
			m[elem.Key] = elem.Value
		}
	}
	return m
}
