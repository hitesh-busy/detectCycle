package main

import (
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

type Mapping struct {
	EmpId  uint
	MangId uint
}

func init() {
	//database connection
	DB = pg.Connect(&pg.Options{

		User:     "postgres",
		Password: "1234",
		Database: "cycle",
		Addr:     "localhost:5432",
	})
	if DB == nil {
		log.Fatalln("Could not connect to the database ")
	}
	log.Println("Connection to DB successful")

	//Create tables
	err := createSchema()
	if err != nil {
		log.Fatal(err)
	}
}
func createSchema() error {
	err := DB.Model((*Mapping)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	log.Println("Tables succesffuly created ")
	return nil
}

func main() {
	fetchDetails()
}

func fetchDetails() {

	// Fetch data into a slice of structs
	var dataSlice []Mapping
	err := DB.Model((*Mapping)(nil)).Select(&dataSlice)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	// Build the graph from the models slice
	graph := make(map[uint][]uint)
	for _, item := range dataSlice {
		graph[item.EmpId] = append(graph[item.EmpId], item.MangId)
	}

	fmt.Println("Graph is ", graph)

	//make slices to check visited
	visited := make(map[uint]bool)
	pathVisited := make(map[uint]bool)

	//accepting empId and mangId to be updated in case there is no cycle

	var empId, mangId uint
	fmt.Println("Enter empId to update")
	fmt.Scanln(&empId)
	fmt.Println("Enter new MangId to update")
	fmt.Scanln(&mangId)

	graph[empId] = append(graph[empId], mangId)

	if checkCycle(empId, visited, pathVisited, graph) {
		fmt.Println("Cyc;e detected ")
		log.Fatalln("Could not updated emp as a cycle is forming")
	}

	fmt.Println("No cycle detected.")
	//if no cycle exists update the mang_id of the specific emp_id in the database
	
	_, err = DB.Model(&Mapping{}).Set("mang_id=?", mangId).Where("emp_id=?", empId).Update()
	if err != nil {
		log.Fatalln("Could not update due to internal error")
	}

}

/*
Idea is to carry two slices visited and path_visited. Using DFS we traverse each manager of the the concerned emp_id node
path_visited helps in igoring the 
 1     2
   \   /
    3 -- 6
   / \
  4   5

The arrows are pointing downwards to represent the direction. now  3 -> 6 and 3->5 seem like a cycle but its not. path_visited is used to avoid this
*/
func checkCycle(empId uint, visited map[uint]bool, path_visited map[uint]bool, mapping map[uint][]uint) bool {

	visited[empId] = true
	path_visited[empId] = true

	for _, manager := range mapping[empId] {
		if !visited[manager] {
			if checkCycle(manager, visited, path_visited, mapping) {
				fmt.Println("Detect cycle triggered for ", manager)
				return true
			}
		} else if path_visited[manager] {
			fmt.Println("Path visited triggere for ", manager)
			return true
		}
	}
	path_visited[empId] = false
	return false

}
