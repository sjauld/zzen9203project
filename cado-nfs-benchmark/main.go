package main

import (
	"encoding/csv"
	"io"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func processKey(row []string) []string {
	// Rows are bits,e,n
	bits, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		log.Printf("[ERROR] %v (%v)", err, row)
		return []string{}
	}
	// e will be 65537
	e, ok := big.NewInt(0).SetString(row[1], 10)
	if !ok {
		log.Printf("[ERROR] (%v)", row)
		return []string{}
	}
	n, ok := big.NewInt(0).SetString(row[2], 10)
	if !ok {
		log.Printf("[ERROR] %v", row)
		return []string{}
	}

	log.Printf("[INFO] cracking %d bit key exponent %d modulus %d", bits, e, n)

	start := time.Now()

	// Call out to cado-nfs
	cmd := exec.Command("../cado-nfs/cado-nfs.py", n.String())
	output, err := cmd.Output()
	if err != nil {
		log.Printf("[ERROR] %v (%v)", err, row)
		return []string{}
	}

	pq := strings.Split(string(output), " ")
	if len(pq) != 2 {
		log.Printf("[WARN] Unexpected output %v", pq)
		return []string{}
	}
	p, ok := big.NewInt(0).SetString(pq[0], 10)
	if !ok {
		log.Printf("[ERROR] (%v)", row)
		return []string{}
	}

	q, ok := big.NewInt(0).SetString(strings.TrimSpace(pq[1]), 10)
	if !ok {
		log.Printf("[ERROR] (%v)", row)
		return []string{}
	}

	log.Printf("[INFO] %s is %s * %s", n.String(), p.String(), q.String())

	d := findPrivateKey(p, q, e)
	t := time.Since(start)

	log.Printf("[INFO] got %d in %v", d, t)
	csvRow := []string{
		strconv.FormatInt(bits, 10),
		e.String(),
		n.String(),
		d.String(),
		strconv.FormatInt(t.Milliseconds(), 10),
	}
	return csvRow
}

func main() {
	input, err := os.Open("keys.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	output, err := os.Create("output.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	csvReader := csv.NewReader(input)
	csvWriter := csv.NewWriter(output)
	headers := []string{"bits", "e", "m", "d", "time (ms)"}
	err = csvWriter.Write(headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		row := processKey(rec)

		err = csvWriter.Write(row)
		if err != nil {
			log.Fatal(err)
		}
		// flush the buffer every time since each iteration takes ages
		csvWriter.Flush()
	}
}
