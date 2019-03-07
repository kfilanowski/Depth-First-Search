// Package vertex - Defines a single vertex on a graph.
//
// Author - Kevin Filanowski
//
// Version - October 2018
package vertex

// Constant colors defined for the vertex.
const (
   // UNVISITED - A vertex has not been visited yet.
	UNVISITED string = "Black"
   // VISITED - A vertex has been visited.
	VISITED string = "Gray"
   // FINISHED - A vertex has been visited and is finished processing.
	FINISHED string = "White"
)

// Vertex - A struct defining a single vertex on a graph.
type Vertex struct {
   // Unique integer corresponding to a vertex's id number. vertices are
   // numbered from 0 to the number of vertices - 1
   ID int

   // Color of the vertex.
   // "Black" means unvisited, "Gray" is in the stack, and "White" is visited.
   Color string
}
