// Package graph - Defines a graph that contains edges and vertices.
// A couple algorithms are available for this algorithm, such as Depth First
// Search, Transitive Closure, and Cycle Search.
//
// Author - Kevin Filanowski
//
// Version - October 2018
package graph

import (
   "fmt"
   "sort"
   "strconv"
   "KF_Project2_DFS/vertex"
   "errors"
   "math"
   "strings"
)

///*** BEGIN SECTION FOR TYPES ***///

// Vertex - A shorthand way to typing vertex.Vertex.
type Vertex vertex.Vertex

// vertexStack - Go does not have a native Stack data structure, so I created
// a simple one of my own. Creates a stack that can only hold type Vertex.
type vertexStack struct {
	stack []Vertex
}

// InvalidVertex - A custom error thrown when there is an invalid vertex in
// the input file.
type InvalidVertex struct {
   Reason string
}

// Graph      - A struct defining a graph. A graph contains three fields:
// VertexList - A list of verticies stored in a slice.
// AdjList    - A 2-D list of vertex pointers that store the vertices adjacent
//              to a specified vertex.
// AdjMatrix  - A 2-D matrix of boolean values that evaluate to true if there
//              is an edge between two vertices.
type Graph struct {
   // List of vertices in the graph where the index
   // corresponds to the vertex number.
   VertexList []Vertex

   // List at each vertex corresponds to the vertices adjacent to vertex.
   AdjList [][]*Vertex

   // Value of true represents an edge between two vertices.
   AdjMatrix [][]bool
}

///*** END SECTION FOR TYPES ***///
///*** BEGIN SECTION FOR STACKS ***///

// add - This function adds a vertex onto a vertexStack.
func (s *vertexStack) add(element Vertex) {
   s.stack = append(s.stack, element)
}

// pop - This function pops off an element from the stack.
// Unlike traditional pop methods, it does not return the element popped.
func (s *vertexStack) pop() {
   if (len(s.stack) > 0) {
      s.stack = s.stack[:len(s.stack)-1]
   }
}

// peek - This function returns the top element of the stack, but does not
// pop it off the stack.
// If there are no more elements in the stack, a zero-value Vertex is returned.
// A zero-values Vertex will have an ID of: 0 and a color of: "".
func (s *vertexStack) peek() Vertex {
   var vertex Vertex
   if (len(s.stack) > 0) {
	   vertex = s.stack[len(s.stack)-1]
   }
   return vertex
}

///*** END SECTION FOR STACKS ***///
///*** BEGIN SECTION FOR VERTECES ***///

// ContainsVertexID - Checks the vertex slice to determine if it contains a
// vertex with the specified ID.
// Returns true if a vertex with the specified ID is in vertexList.
func (g Graph) ContainsVertexID(id int) bool {
    for _, e := range g.VertexList {
        if e.ID == id {
            return true
        }
    }
    return false
}

// Error - The toString equivalent method of the InvalidVertex error.
// Prints the default error message if the error is printed.
func (e InvalidVertex) Error() string {
	return fmt.Sprintf("%s", "Invalid vertex in input file. Vertices must be" +
      " expressed as integers.")
}

// ResetColors - This method resets all the colors of the vertices in
// vertexList, as well as AdjList, to UNVISITED. This is necessary when
// wanting to traverse the graph more than once, using colors.
func (g *Graph) ResetColors() {
   for i := 0; i < len(g.VertexList); i++ {
      g.VertexList[i].Color = vertex.UNVISITED
   }
}

///*** END SECTION FOR VERTECES ***///
///*** BEGIN SECTION FOR GRAPH ***///

// Run - Entry point to run operations specified in the handout.
func (g *Graph) Run(result []string) error {
   var err error // Declaring a spot for errors to be returned.

   err = g.SetUpVertices(result)        // Set up vertices.
   if err != nil { return err }         //
   err = g.SetUpAdjacencyList(result)   // Set up adjacency list.
   if err != nil { return err }         //
   g.SetUpAdjacencyMatrix()             // Set up adjacency matrix.
                                        //
   src, dest, err := g.FindSourceDest() // Find Source and Destination Vertex.
   if err != nil { return err }         //
   disc, path := g.DepthFirstSearch(src, dest) // Perform a depth first search.
   newEdges := g.TransitiveClosure()    // Perform Transitive Closure algorithm.
   cycle := g.CycleSearch()             // Searches for a cycle.
   g.PrintGraphStats(disc, path, newEdges, cycle) // Print Results.

   return err
}

// FindSourceDest - Finds the source and destination vertices.
// On invalid input, asks the user to enter new vertices.
// Throws an error if the source and destination vertices
// are not found in the graph.
// Returns a potential error.
func (g Graph) FindSourceDest() (int, int, error) {
   // Message to be printed.
   msg := "Please enter valid source and destination vertices >> "
   var src int
   var dest int

   // BUG - Entering plaintext causes the print message to be repeated once
   // for every character. Program will not crash or fail, but has peculiar
   // output. This occurs because Scanf stops once it recieves two inputs, then
   // the for loop continues, and it takes the next two characters as input
   // again, and we get a cycle of invalid input and printed messages.
   fmt.Print(msg)
   for _, e := fmt.Scanf("%d %d", &src, &dest); e != nil; {
      fmt.Print(msg)
      _, e = fmt.Scanf("%d %d", &src, &dest)
   }

   // Check if these are in the graph. Return an error if they are not.
   if !g.ContainsVertexID(src) || !g.ContainsVertexID(dest) {
      return src, dest, errors.New("vertices specified " +
                                   "do not exist in the graph")
   }
   return src, dest, nil
}

// DepthFirstSearch - Performs a depth-first search from source to destination
// or until the search is exhausted.
func (g Graph) DepthFirstSearch(src int, dest int) ([]Vertex, []Vertex) {
   VL := g.VertexList // Assigning this to reduce keystrokes.

   var lowestID int // Used for comparison between vertex ID's
   var discoveredVertex []Vertex // Used to track what vertex were discovered.
   var path []Vertex // Path from source to destination.

   stack := vertexStack{} // Go does not have built in stacks, so I made one.
   VL[src].Color = vertex.VISITED // Change the color of the source
   stack.add(VL[src])             // Add the source vertex.
   discoveredVertex = append(discoveredVertex, VL[src])
   path = append(path, VL[src])
   if src == dest {
      return discoveredVertex, path
   }

   for len(stack.stack) > 0 {
      lowestID = math.MaxInt32
      for _, e := range g.AdjList[stack.peek().ID] { // For each loop.
         if e.ID < lowestID &&
               strings.Compare(e.Color, vertex.UNVISITED) == 0 {
            lowestID = e.ID // Pick the lowest vertex ID that is unvisited.
         }
      }
      if lowestID == math.MaxInt32 { // If no new vertices found, then pop.
         VL[stack.peek().ID].Color = vertex.FINISHED
         path = path[:len(path)-1]
         stack.pop()
      } else { // Otherwise, add it and repeat the loop to go deeper.
         VL[lowestID].Color = vertex.VISITED
         stack.add(VL[lowestID])
         discoveredVertex = append(discoveredVertex, VL[lowestID])
         path = append(path, VL[lowestID])
         if VL[lowestID].ID == VL[dest].ID { // If found, return immediately.
            return discoveredVertex, path
         }
      }
   }
   return discoveredVertex, []Vertex{} // If no path was found.
}

// CycleSearch - Performs a search for cycles.
// Returns true if there exists a cycle in the graph.
func (g Graph) CycleSearch() bool {
   VL := g.VertexList // Assigning this to reduce keystrokes.
   var lowestID int   // Used for comparison between vertex ID's.
   stack := vertexStack{} // Go does not have built in stacks, so I made one.
   for _, v := range VL {             // Run for each vertex.
      g.ResetColors()                 // Reset the colors.
      VL[v.ID].Color = vertex.VISITED // Change the color of the vertex.
      stack.add(VL[v.ID])             // Add the source vertex.
      for len(stack.stack) > 0 {
         lowestID = math.MaxInt32
         for _, e := range g.AdjList[stack.peek().ID] { // For each loop.
            if strings.Compare(e.Color, vertex.VISITED) == 0 {
               return true; // Cycle found.
            }
            if e.ID < lowestID &&
                  strings.Compare(e.Color, vertex.UNVISITED) == 0 {
               lowestID = e.ID // Pick the lowest vertex ID that is unvisited.
            }
         }
         if lowestID == math.MaxInt32 { // If no new vertices found, then pop.
            VL[stack.peek().ID].Color = vertex.FINISHED
            stack.pop()
         } else { // Otherwise, add it and repeat the loop to go deeper.
            VL[lowestID].Color = vertex.VISITED
            stack.add(VL[lowestID])
         }
      }
   }
   return false // Entire tree was traversed without a cycle.
}

// SetUpAdjacencyList - Sets up the adjacency list to correspond to the graph.
// Returns a potential error.
func (g *Graph) SetUpAdjacencyList(result []string) error {
   // Allocate memory for the adjacency list.
   g.AdjList = make([][]*Vertex, len(g.VertexList), len(g.VertexList))
   for i := 0; i < len(g.VertexList); i++ {
      g.AdjList[i] = make([]*Vertex, 0, len(g.VertexList))
   }

   // Add to adjacency list.
   for i := 0; i < len(result); {
      src, err := strconv.Atoi(result[i])
      if err != nil { return InvalidVertex{} }
      i++
      dest, err := strconv.Atoi(result[i])
      if err != nil { return InvalidVertex{} }
      g.AdjList[src] = append(g.AdjList[src], &g.VertexList[dest])
      i++
   }
   return nil
}

// SetUpAdjacencyMatrix - Sets up the adjacency matrix to correspond
// to the graph.
func (g *Graph) SetUpAdjacencyMatrix() {
   // Allocate memory for the adjacency matrix.
   g.AdjMatrix = make([][]bool, len(g.VertexList), len(g.VertexList))
   for i := 0; i < len(g.VertexList); i++ {
      g.AdjMatrix[i] = make([]bool, len(g.VertexList), len(g.VertexList))
   }

   // Adjust values in adjacency matrix based on the adjacency list.
   for row := 0; row < len(g.AdjList); row++ {
      for col := 0; col < len(g.AdjList[row]); col++ {
         g.AdjMatrix[row][g.AdjList[row][col].ID] = true
      }
   }
}

// SetUpVertices - Adds the vertices in the graph.
// Returns a potential error.
func (g *Graph) SetUpVertices(result []string) error {
   // Add vertices to the list.
   for _, element := range result {
      // Convert scanned character to integer value.
      val, err := strconv.Atoi(element)
      if err != nil {
         return InvalidVertex{}
      }

      // Check if the vertex is already in the list, and add it if it is not.
      if !g.ContainsVertexID(val) {
         g.VertexList = append(g.VertexList, Vertex{val, vertex.UNVISITED})
      }
   }

   // Sort vertices by ID number.
   sort.Slice(g.VertexList, func(i, j int) bool {
		return g.VertexList[i].ID < g.VertexList[j].ID
	})
   return nil
}

// TransitiveClosure - Performs the transitive closure algorithm.
// Returns a slice of new edges formed.
func (g *Graph) TransitiveClosure() []int {
   // adj is defined as just a way to shorten key-strokes.
   adj := g.AdjMatrix

   // The individual values in the selected row and column.
   var rowVal bool
   var colVal bool

   // Keeps track of new edges that were added to the graph.
   // Allocating memory to slice.
   newEdges := make([]int, 0, len(adj) * len(adj))

   // Algorithm
   for diag := 0; diag < len(adj); diag++ {
      for row := 0; row < len(adj); row++ {
         rowVal = adj[row][diag]
         for col := 0; col < len(adj); col++ {
            colVal = adj[diag][col]
            if rowVal && colVal {
               if !g.AdjMatrix[row][col] {
                  g.AdjMatrix[row][col] = true
                  newEdges = append(newEdges, row, col)
               }
            }
         }
      }
   }
   return newEdges
}

// PrintGraphStats - Prints the information as specified in the handout.
func (g Graph) PrintGraphStats(discover []Vertex, path []Vertex,
                                    newEdges []int, cycle bool) {
   if (len(path) > 0) {
      // Prints the discovered vertices of the Depth First Search.
      fmt.Printf("\n[DFS Discovered Vertices: %d, %d] ",
                                 discover[0].ID, discover[len(discover)-1].ID)
      for i := 0; i < len(discover); i++ {
         if i == len(discover) - 1 {
            fmt.Printf("Vertex %d", discover[i].ID)
         } else {
            fmt.Printf("Vertex %d, ", discover[i].ID)
         }
      }
      // Prints the path from source to destination of the Depth First Search.
      fmt.Printf("\n\n[DFS Path: %d, %d] ", path[0].ID, path[len(path)-1].ID)
      for i := 0; i < len(path); i++ {
         if i == len(path) - 1 {
            fmt.Printf("Vertex %d\n", path[i].ID)
         } else {
            fmt.Printf("Vertex %d -> ", path[i].ID)
         }
      }
   } else {
      fmt.Println("\n[DFS Not Found]")
   }
   // Prints the new edges added by the Transitive Closure algorithm.
   for i := 0; i < len(newEdges)-1; i += 2 {
      if i != 0 {
         fmt.Printf("%17d \t %d\n", newEdges[i], newEdges[i+1])
      } else {
         fmt.Printf("\n[TC: New Edges] %d \t %d\n", newEdges[i], newEdges[i+1])
      }
   }
   // Prints the result of CycleSearch.
   if cycle {
      fmt.Println("\n[Cycle]: Cycle detected")
   } else {
      fmt.Println("\n[Cycle]: No Cycle detected")
   }
}
///*** END SECTION FOR GRAPH ***///
