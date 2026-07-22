package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
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
	n := len(data.Shipments)
	FilteredArr := []Shipment{}
	for i:=0;i<n;i++ {
		if(data.Shipments[i].Status=="Delayed" || data.Shipments[i].Status=="Customs Hold"){
			FilteredArr = append(FilteredArr,data.Shipments[i])
		}
	}
	sort.Slice(FilteredArr,func(i,j int) bool {
		return FilteredArr[i].EstimatedDelivery < FilteredArr[j].EstimatedDelivery
	})
	nn := len(FilteredArr)
	for i:=0;i<nn;i++ {
		fmt.Println(FilteredArr[i].ShipmentID,FilteredArr[i].Status,FilteredArr[i].EstimatedDelivery)
	}
}
func totalRevenue(data LogisticsData){
	//Calculate the total revenue (cost_breakdown.total) grouped by carrier
	TotalRevenue := make(map[string]float64)
	n := len(data.Shipments)
	for i:=0;i<n;i++ {
		Carrier := data.Shipments[i].Carrier
		Revenue := data.Shipments[i].CostBreakdown.Total
		TotalRevenue[Carrier] = TotalRevenue[Carrier]+Revenue
	}
	fmt.Println(TotalRevenue)
}
func maximumRevenue(data LogisticsData){
	TotalRevenue := make(map[string]float64)
	n := len(data.Shipments)
	for i:=0;i<n;i++ {
		Carrier := data.Shipments[i].Carrier
		Revenue := data.Shipments[i].CostBreakdown.Total
		TotalRevenue[Carrier] = TotalRevenue[Carrier]+Revenue
	}
	var Maxi float64
	var Carrier string
	for key,value := range TotalRevenue {
		if(value>Maxi){
			Maxi = value
			Carrier = key
		}
	}
	fmt.Println(Carrier,Maxi)
}
func longestTransit(data LogisticsData){
	var maxDuration time.Duration
	var longestShipment Shipment
	for _,shipment := range data.Shipments {
		if len(shipment.TrackingEvents) < 2 {
			continue
		}
		firstTime , err := time.Parse(time.RFC3339,shipment.TrackingEvents[0].Timestamp)
		if err!=nil {
			continue
		}
		lastTime,err := time.Parse(time.RFC3339,shipment.TrackingEvents[len(shipment.TrackingEvents)-1].Timestamp)
		if err!=nil {
			continue
		} 
		duration := lastTime.Sub(firstTime)
		if duration > maxDuration {
			maxDuration = duration
			longestShipment = shipment
		}
	}
	fmt.Println("Shipment ID :", longestShipment.ShipmentID)
	fmt.Println("Transit Time:", maxDuration.Hours(), "hours")
}
func riskReport(data LogisticsData) {
	//Identify all is_international shipments where weight_kg > 20 and cost_breakdown.insurance == 0.
	//These represent heavy international shipments with no insurance — flag them as a "risk report."
	for _,shipment := range data.Shipments {
		if shipment.WeightKg > 20 && shipment.CostBreakdown.Insurance==0 {
			fmt.Println("Risk Report : ",shipment.ShipmentID)
		}
	}
}
func groupByCountry(data LogisticsData) {
	summary := make(map[string]map[string]int)
	for _, shipment := range data.Shipments {
		country := shipment.Destination.Country
		status := shipment.Status
		_, exists := summary[country]
		if !exists {
			summary[country] = make(map[string]int)
		}

		summary[country][status]++
	}

	fmt.Println(summary)
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
	totalRevenue(logisticsData)
	maximumRevenue(logisticsData)
	longestTransit(logisticsData)
	riskReport(logisticsData)
	groupByCountry(logisticsData)
}