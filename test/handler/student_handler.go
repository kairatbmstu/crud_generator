package handler

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/lib/pq"
)

type StudentHandler struct {
	StudentService service.StudentService
}
