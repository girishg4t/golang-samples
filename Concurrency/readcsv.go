package main

import (
	//"bufio"
	"context" // manage multiple requests
	"encoding/csv"
	"fmt" // Println() function
	"io"
	"log"
	"os" // os.Exit(1) on Error
	"sync"

	// get an object type
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var wg sync.WaitGroup
	// Open the file
	csvfile, err := os.Open("./input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	stockQuotes, err := r.Read()
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	//ctx, _ := context..WithTimeout(context.Background(), 1000*time.Second)
	ctx := context.TODO()
	db := client.Database("test")

	// Iterate through the records
	for {
		// Read each record from csv
		stockQuotes, err = r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if stockQuotes[1] != "EQ" {
			continue
		}
		layOut := "2-Jan-2006"

		date, _ := time.Parse(layOut, stockQuotes[10])
		open, _ := strconv.ParseFloat(stockQuotes[2], 64)
		high, _ := strconv.ParseFloat(stockQuotes[3], 64)
		low, _ := strconv.ParseFloat(stockQuotes[4], 64)
		close, _ := strconv.ParseFloat(stockQuotes[5], 64)
		volume, _ := strconv.ParseInt(stockQuotes[9], 20, 64)
		stockName := stockQuotes[0] + ".BO"
		todaysStockPrice := bson.D{
			{"date", date},
			{"open", open},
			{"high", high},
			{"low", low},
			{"close", close},
			{"volume", volume},
			{"adjClose", nil},
			{"symbol", stockName},
		}
		col := db.Collection(stockName)
		wg.Add(1)
		go insertIntoDb(col, ctx, todaysStockPrice, &wg)
	}
	wg.Wait()
}

func insertIntoDb(col *mongo.Collection, ctx context.Context, s primitive.D, wg *sync.WaitGroup) {
	_, insertErr := col.InsertOne(ctx, s)
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
	} else {
		fmt.Println("Done saving ")
	}
	wg.Done()
}

//Person to store in db
type todayStockBhav struct {
	Date     time.Time
	open     float64
	high     float64
	low      float64
	close    float64
	volume   int64
	adjClose *float64
	symbol   string
}
