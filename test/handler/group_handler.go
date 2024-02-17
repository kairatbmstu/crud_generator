package handler

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/lib/pq"
)

type GroupHandler struct {
	StudentService service.StudentService
}
