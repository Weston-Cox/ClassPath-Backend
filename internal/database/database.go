//* internal/database/database.go:
//****************************************************************************************
//* Contains database connection logic and functions for interacting with the PostgreSQL database.
//****************************************************************************************

package database

import (
	"context"
	_ "log"
	"time"

	"github.com/Weston-Cox/ClassPath-Backend/internal/models"
	"github.com/jackc/pgx/v5"
)


func Connect(databaseURL string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetDegrees(conn *pgx.Conn) ([]models.Degree, error) {
	return queryRows(conn, "SELECT * FROM degrees", scanDegree)
}

func GetCourses(conn *pgx.Conn) ([]models.Course, error) {
	return queryRows(conn, "SELECT * FROM courses", scanCourse)
}

func GetDegreeElectives(conn *pgx.Conn) ([]models.Degree_Elective, error) {
	return queryRows(conn, "SELECT * FROM degree_electives", scanDegreeElective)
}

func GetCourseRequisites(conn *pgx.Conn) ([]models.Course_Requisite, error) {
	return queryRows(conn, "SELECT * FROM course_requisites", scanCourseRequisites)
}

func GetDegreeCourses(conn *pgx.Conn) ([]models.Degree_Course, error) {
	return queryRows(conn, "SELECT * FROM degree_courses", scanDegreeCourses)
}

func GetNodePositions(conn * pgx.Conn) ([]models.Node_Positions, error) {
	return queryRows(conn, "SELECT * FROM node_positions", scanNodePositions)
}



//*****************************************************************
//* HELPERS
//*****************************************************************

func queryRows[T any](conn *pgx.Conn, query string, scanFunc func(pgx.Rows) (T, error)) ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		result, err := scanFunc(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return results, nil
}


func scanDegree(rows pgx.Rows) (models.Degree, error) {
	var degree models.Degree
	err := rows.Scan(
		&degree.Degree_ID,
		&degree.Program,
		&degree.Title,
		&degree.Shorthand_Title,
		&degree.Rows,
		&degree.Columns,
	)
	return degree, err
}


func scanCourse(rows pgx.Rows) (models.Course, error) {
	var course models.Course
	err := rows.Scan(
		&course.Course_ID,
		&course.Title,
		&course.Description,
	)
	return course, err
}


func scanDegreeElective(rows pgx.Rows) (models.Degree_Elective, error) {
	var degree_elective models.Degree_Elective
	err := rows.Scan(
		&degree_elective.Degree_ID,
		&degree_elective.General_Electives,
		&degree_elective.Specific_Electives,
		&degree_elective.History_Elective,
		&degree_elective.Fine_Arts_Elective,
		&degree_elective.Social_Sciences_Elective,
		&degree_elective.Science_Elective,
	)
	return degree_elective, err
}


func scanCourseRequisites(rows pgx.Rows) (models.Course_Requisite, error) {
	var course_requisite models.Course_Requisite
	err := rows.Scan(
		&course_requisite.Source,
		&course_requisite.Target,
		&course_requisite.Corequisite,
	)
	return course_requisite, err
}


func scanDegreeCourses(rows pgx.Rows) (models.Degree_Course, error) {
	var degree_course models.Degree_Course
	err := rows.Scan(
		&degree_course.Degree_ID,
		&degree_course.Course_ID,
		&degree_course.Semester_Offered,
	)
	return degree_course, err
}


func scanNodePositions(rows pgx.Rows) (models.Node_Positions, error) {
	var node_position models.Node_Positions
	err := rows.Scan(
		&node_position.Course_ID,
		&node_position.Row,
		&node_position.Column,
		&node_position.Degree_ID,
	)
	return node_position, err
}