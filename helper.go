package helper

import (
	"crypto/x509"
	"encoding/json"
	"errors"
	"github.com/getsentry/raven-go"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func FloatToString(float float64) string {
	return strconv.FormatFloat(float, 'f', 6, 64)
}

func StringToFloat(string string, bitSize int) float64 {
	float, err := strconv.ParseFloat(string, bitSize)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return float
}

func TimestampStringToDate(timestamp string) time.Time {
	timestampInt, err := strconv.ParseInt(timestamp, 10, 64)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return time.Unix(0, timestampInt*int64(time.Millisecond)).UTC()
}

func Debug(data ...interface{}) {
	for _, v := range data {
		log.Printf("[%v] %+v\n", rand.Intn(1000), v)
	}

	log.Print("\n")
}

func Dump(data interface{}) {
	log.Printf("%+v\n", data)
}

func DD(data interface{}) {
	Dump(data)
	os.Exit(1)
}

func ToJSON(data interface{}) string {
	JSON, err := json.Marshal(data)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return string(JSON)
}

func Sleep(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}

func CurrentMinute() int64 {
	return time.Now().UTC().Unix() / 60
}

func ThrowError(message string) {
	err := errors.New(message)
	raven.CaptureError(err, nil)
	log.Fatal(err.Error())
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func StringToInt64(value string) int64 {
	res, err := strconv.ParseInt(value, 10, 64)

	if err != nil {
		raven.CaptureError(err, nil)
		log.Fatal(err.Error())
	}

	return res
}

func CertPool(certs []string) *x509.CertPool {
	roots := x509.NewCertPool()

	for _, cert := range certs {
		cert, err := ioutil.ReadFile(cert)

		if err != nil {
			raven.CaptureError(err, nil)
			log.Fatal(err.Error())
		}

		roots.AppendCertsFromPEM(cert)
	}

	return roots
}

func LoadFile(file string) string {
	callPath := filepath.Dir(os.Args[0])

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return callPath + "/" + file
	}

	return file
}
