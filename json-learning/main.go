package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Metadata struct {
	GeneratedOn   string `json:"generated_on"`
	RecordCount   int    `json:"record_count"`
	SchemaVersion string `json:"schema_version"`
}

type Location struct {
	Country       string `json:"country"`
	City          string `json:"city"`
	WarehouseCode string `json:"warehouse_code"`
	AddressLine   string `json:"address_line"`
	PostalCode    string `json:"postal_code"`
}

type Item struct {
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
}

type Cost struct {
	BaseRate      float64 `json:"base_rate"`
	FuelSurcharge float64 `json:"fuel_surcharge"`
	Insurance     float64 `json:"insurance"`
	Tax           float64 `json:"tax"`
	Total         float64 `json:"total"`
	Currency      string  `json:"currency"`
}

type Event struct {
	Timestamp    string `json:"timestamp"`
	Location     string `json:"location"`
	StatusUpdate string `json:"status_update"`
}

type Shipment struct {
	ShipmentID        string   `json:"shipment_id"`
	OrderID           string   `json:"order_id"`
	Carrier           string   `json:"carrier"`
	Status            string   `json:"status"`
	IsInternational   bool     `json:"is_international"`
	Origin            Location `json:"origin"`
	Destination       Location `json:"destination"`
	Items             []Item   `json:"items"`
	CostBreakdown     Cost     `json:"cost_breakdown"`
	WeightKg          float64  `json:"weight_kg"`
	CreatedAt         string   `json:"created_at"`
	EstimatedDelivery string   `json:"estimated_delivery"`
	TrackingEvents    []Event  `json:"tracking_events"`
}

type LogisticsData struct {
	Metadata  Metadata   `json:"metadata"`
	Shipments []Shipment `json:"shipments"`
}
func shipments(data LogisticsData) {
	//Filtering + Sorting: Find all shipments where status is "Delayed" or "Customs Hold", 
	// and sort them by estimated_delivery in ascending order.
	n := len(data.Shipments)
	StatusArr := []string{}
	for i:=0;i<n;i++ {
		if(data.Shipments[i].Status=="Delayed" || data.Shipments[i].Status=="Customs Hold"){
			StatusArr = append(StatusArr,data.Shipments[i].ShipmentID)
		}
	}
	fmt.Println(StatusArr)
}
func main() {
	file, err := os.ReadFile("shipping_data.json")
	if err != nil {
		panic(err)
	}
	var logisticsData LogisticsData
	err = json.Unmarshal(file, &logisticsData)
	if err != nil {
		panic(err)
	}
	shipments(logisticsData)
}