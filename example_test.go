package entspatial

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/a8m/entspatial/ent"
	"github.com/a8m/entspatial/ent/schema"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func Example_Point() {
	client := Open("postgresql://viecco:example@localhost:5433/viec")
	defer client.Close()
	ctx := context.Background()

	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		panic(err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	tlv := client.Location.
		Create().
		SetName("TLV").
		SetCoords(schema.Point{32.109333, 34.855499}).
		SaveX(ctx)
	fmt.Println(tlv.Name, tlv.Coords)
	office := client.Location.
		Create().
		SetName("FB").
		SetCoords(schema.Point{32.072184, 34.78471}).
		SetParent(tlv).
		SaveX(ctx)
	fmt.Println(office.Name, office.Coords)

	// Output:
	// BEGIN;
	// ALTER TABLE "locations" ALTER COLUMN "coords" TYPE Geometry(POINT,4326), ALTER COLUMN "coords" SET NOT NULL;
	// COMMIT;
	// TLV [32.109333 34.855499]
	// FB [32.072184 34.78471]
}
