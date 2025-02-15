//* internal/handlers/tree_handler.go:
//****************************************************************************************
//* Contains HTTP handlers for the API endpoints.These handlers will receive requests,
//* call the appropriate services, and return responses.
//****************************************************************************************

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Weston-Cox/ClassPath-Backend/internal/database"
	"github.com/Weston-Cox/ClassPath-Backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type Http_Handler struct {
	DB *pgx.Conn
}

// type Models struct {
// 	Degree []models.Degree
// 	Course []models.Course
// 	Degree_Course []models.Degree_Course
// 	Degree_Elective []models.Degree_Elective
// 	Course_Requisite []models.Degree_Course
// }

func NewHttpHandler(db *pgx.Conn) *Http_Handler {
	return &Http_Handler{DB: db}
}


func (http_handler *Http_Handler) SetupRoutes() {
	http.HandleFunc("/", http_handler.enableCors(http_handler.RootHandler))
	http.HandleFunc("/degrees", http_handler.enableCors(http_handler.DegreesHandler))
	http.HandleFunc("/degree-electives", http_handler.enableCors(http_handler.DegreeElectivesHandler))
	http.HandleFunc("/courses", http_handler.enableCors(http_handler.CoursesHandler))
	http.HandleFunc("/degree-courses", http_handler.enableCors(http_handler.DegreeCoursesHandler))
	http.HandleFunc("/course-requisites", http_handler.enableCors(http_handler.CourseRequisitesHandler))
	http.HandleFunc("/node-positions", http_handler.enableCors(http_handler.NodePositionsHandler))
}


func (http_handler *Http_Handler) RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Connected to the server")
} 


func (http_handler *Http_Handler) DegreesHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestHandler(w, *http_handler, r, getDegrees)
}


func (http_handler *Http_Handler) DegreeElectivesHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestHandler(w, *http_handler, r, getDegreeElectives)
}


func (http_handler *Http_Handler) DegreeCoursesHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestHandler(w, *http_handler, r, getDegreeCourses)
}


func (http_handler *Http_Handler) CoursesHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestHandler(w, *http_handler, r, getCourses)
}


func (http_handler *Http_Handler) CourseRequisitesHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestHandler(w, *http_handler, r, getCourseRequisites)
}


func (http_handler *Http_Handler) NodePositionsHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestHandler(w, *http_handler, r, getNodePositions)
}


//***********************************************************************************************
//* HELPERS
//***********************************************************************************************

func (http_handler *Http_Handler) enableCors(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		// allowedOrigin := "http://localhost:5173/TreeScreen"
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next(w, r)
    }
}


func httpRequestHandler[T any](w http.ResponseWriter, http_handler Http_Handler, r *http.Request, queryDbFunc func(Http_Handler) ([]T, error)) {
	if r.Method == http.MethodGet {
		data, err := queryDbFunc(http_handler)
		if err != nil {
			fmt.Printf("Erros: %v\n", err)
			http.Error(w, "Failed to get data", http.StatusInternalServerError)
			return
		}

		send(w, data)

	} else if r.Method == http.MethodPost {
		fmt.Println("Now allowed to add data yet.")
	}
}


func getDegrees(http_handler Http_Handler) ([]models.Degree, error) {
	return database.GetDegrees(http_handler.DB)
}


func getDegreeElectives(http_handler Http_Handler) ([]models.Degree_Elective, error) {
	return database.GetDegreeElectives(http_handler.DB)
}


func getDegreeCourses(http_handler Http_Handler) ([]models.Degree_Course, error) {
	return database.GetDegreeCourses(http_handler.DB)
}


func getCourses(http_handler Http_Handler) ([]models.Course, error) {
	return database.GetCourses(http_handler.DB)
}


func getCourseRequisites(http_handler Http_Handler) ([]models.Course_Requisite, error) {
	return database.GetCourseRequisites(http_handler.DB)
}


func getNodePositions(http_handler Http_Handler) ([]models.Node_Positions, error) {
	return database.GetNodePositions(http_handler.DB)
}


func send[T any](w http.ResponseWriter, data T) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}