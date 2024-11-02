package apdb

import (
	"context"
	"fmt"
	"Ga-backend/model"
	"net"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(mconn model.DBinfo) (*mongo.Database, error) {
    // Coba koneksi awal dengan DBString yang diberikan
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
    if err != nil {
        // Jika gagal, lakukan SRVLookup dan coba kembali
        mconn.DBString = SRVLookup(mconn.DBString)
        client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mconn.DBString))
        if err != nil {
            return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
        }
    }

    // Verifikasi koneksi ke MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := client.Ping(ctx, nil); err != nil {
        return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
    }

    // Pilih database dan kembalikan
    db := client.Database(mconn.DBName)
    return db, nil
}

func SRVLookup(srvuri string) (mongouri string) {
    atsplits := strings.Split(srvuri, "@")
    userpass := strings.Split(atsplits[0], "//")[1]
    mongouri = "mongodb://" + userpass + "@"

    slashsplits := strings.Split(atsplits[1], "/")
    domain := slashsplits[0]
    dbname := slashsplits[1]

    r := &net.Resolver{
        PreferGo: true,
        Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
            d := net.Dialer{
                Timeout: time.Millisecond * time.Duration(10000),
            }
            return d.DialContext(ctx, network, "8.8.8.8:53")
        },
    }

    // Lookup SRV records for MongoDB
    _, srvs, err := r.LookupSRV(context.Background(), "mongodb", "tcp", domain)
    if err != nil {
        panic(err)
    }

    var srvlist string
    for _, srv := range srvs {
        srvlist += strings.TrimSuffix(srv.Target, ".") + ":" + strconv.FormatUint(uint64(srv.Port), 10) + ","
    }

    // Lookup TXT records for MongoDB
    txtrecords, _ := r.LookupTXT(context.Background(), domain)
    var txtlist string
    for _, txt := range txtrecords {
        txtlist += txt
    }

    mongouri = mongouri + strings.TrimSuffix(srvlist, ",") + "/" + dbname + "?ssl=true&" + txtlist
    return
} 