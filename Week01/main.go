package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := service()
	if errors.Cause(err) == sql.ErrNoRows {
		// %+v 输出完整的调用栈信息
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Println(err)
}

func dao() error {
	return errors.Wrap(sql.ErrNoRows, "dao error")
}

func service() error {
	return errors.WithMessage(dao(), "service error")
}
