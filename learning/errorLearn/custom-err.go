package errorLearn

import (
	"fmt"
)

// CustomError is a custom error type
type ConnectionDBError struct {
	Msg string
	Status int16
}

func (e *ConnectionDBError) Error() string {
	return fmt.Sprintf("Error: %s", e.Msg)
}

type TableDBError struct {
	Msg string
}

func (e *TableDBError) Error() string {
	return fmt.Sprintf("Error: %s", e.Msg)
}

func TryDBConnection(ip string, table string) error {

	if ip == "" {
		return &ConnectionDBError{Msg: "Connection to database failed", Status: 404}
	}

	if table == "" {
		return &TableDBError{Msg: "Table not found"}
	}

	return nil

}

func ErrorCustom() {

	err := TryDBConnection("", "table")
	if err != nil {
		fmt.Println(err)
		// jika kita ingin melakukan pengecekan error apa yang terjadi
		// kita bisa menggunakan switch case atau if
		switch errData := err.(type) {
		case *ConnectionDBError:
			fmt.Println("Connection error", errData.Status)
		case *TableDBError:
			fmt.Println("Table error")
		}
	} else {
		fmt.Println("Connection success")
	}

}
