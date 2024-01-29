package geoip

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func GetCityRecord(reqIP string) *geoip2.City {
	db, err := geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ip := net.ParseIP(reqIP)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	return record

}
