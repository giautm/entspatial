## entspatial

An example repository for working with [MySQL spatial data types](https://dev.mysql.com/doc/refman/8.0/en/spatial-type-overview.html) in [ent](https://entgo.io).


```go
// Location holds the schema definition for the Location entity.
type Location struct {
	ent.Schema
}

// Fields of the Location.
func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Other("coords", &Point{}).
			SchemaType(Point{}.SchemaType()),
	}
}
```

Run following command to start PostGIS docker
```sh
docker run -d \
  -p 127.0.0.1:5433:5432/tcp \
  -e POSTGRES_DB=viec \
  -e POSTGRES_USER=viecco \
  -e POSTGRES_PASSWORD=example \
  postgis/postgis
```