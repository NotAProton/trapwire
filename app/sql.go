package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func insertLog(time string, ip string, ua string, referrer string, wire string, dest string) {
	stmtIns, err := db.Prepare("INSERT INTO logs (time, ip, ua, referrer, wire, dest) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	// Execute the statement with the provided values
	_, err = stmtIns.Exec(time, ip, ua, referrer, wire, dest)
	if err != nil {
		panic(err.Error())
	}
}

func insertWire(wire string, dest string) {
	_, found := getDest(wire)
	if found {
		stmtMod, err := db.Prepare("UPDATE wires SET dest = ? WHERE wire = ?")
		if err != nil {
			panic(err.Error())
		}
		defer stmtMod.Close()
		_, err = stmtMod.Exec(dest, wire)
		if err != nil {
			panic(err.Error())
		}
		return
	}
	stmtIns, err := db.Prepare("INSERT INTO wires (wire, dest) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(wire, dest)
	if err != nil {
		panic(err.Error())
	}
}

func getDest(wire string) (string, bool) {
	var dest string
	q := db.QueryRow("SELECT dest FROM wires WHERE wire = ?", wire)
	err := q.Scan(&dest)
	if err != nil {
		return "", false
	}
	return dest, true
}

func getLogsData() []map[string]string {
	rows, err := db.Query("SELECT time, ip, ua, referrer, wire, dest FROM logs")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	var logs []map[string]string
	for rows.Next() {
		var time, ip, ua, referrer, wire, dest string
		err := rows.Scan(&time, &ip, &ua, &referrer, &wire, &dest)
		if err != nil {
			panic(err.Error())
		}
		logs = append(logs, map[string]string{
			"time":     time,
			"ip":       ip,
			"ua":       ua,
			"referrer": referrer,
			"wire":     wire,
			"dest":     dest,
		})
	}
	return logs
}
