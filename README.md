# Author
Kevin Filanowski

# Version
10/25/18 - 1.0.0

Practicing the Go programming language.

# Table of Contents
* [Description](#description)
* [Contents](#contents)
* [Set Up](#set-up)
* [Usage](#usage)
* [Input File](#input-file)
* [Documentation](#documentation)

# Description
This program performs several algorithms tailored toward graphs, such as 
Depth First Search, Cycle search, and Transitive Closure.
It takes a file as input, where the file contains information about the graph.
Given that information, the program can attempt to locate a path from a specified
source and destination vertex. The program will print a formatted message including
the discovered vertices, the path from source to destination (if found), the new 
edges added by the transitive closure algorithm, and if a cycle was found in the graph.

Depth First Search - An algorithm that searches the 'depth' of a graph
before searching 'breadth'. 
Transitive Closure - For every possible path from every vertex to
another vertex, make an edge between them.
Cycle Search - Detects a cycle in the graph.

# Contents
* graph
  * A folder containing the graph source file.
* graph/Graph.go
  * This file models a graph and performs various algorithms on it.

* vertex
  * A folder containing the vertex source file.
* vertex/Vertex.go
  * This file models a simple vertex contained in a graph.

* main
  * A folder containing the driver source file.
* main/Driver.go
  * The main driver of the program. Reads in command line arguments, which accepts a file and forms a graph on it. 
* main/assign.txt
  * An example format of how a graph is represented.

# Set Up

To run the program, we need to have GO installed and setup on the computer.
The setup needs to include a GOPATH, and in the GOATH, we should have 
three folders: src, pkg, and bin. Ensure the directory 'KF_Project2_DFS'
is in the src folder. 


# Usage
To run the program, Use the terminal to navigate to 'KF_Project2_DFS/main' 
directory and run the following command in the terminal:

`go run Driver.go <input_file>`

The program will ask for user input, requiring the user to specify
a source and a destination.

NOTE://
Not required, but we can create an executable file by going into the 
'KF_Project2_DFS/main' directory and typing: `go build -o dfs`
Then we can just run:
`dfs <input_file>`
or `./dfs <input_file>`
We can also place this file into the $GOPATH/bin directory to run it from
anywhere on the system, assuming the path environment variables
are properly set up. 


# Input File
The graph is represented by an adjacency list. In the input file,
we have the format:

source destination
source destination
source destination

etc.

Where source and destination are a vertex, separated by whitespace.
Source has a path to destination, but not necessarily the other 
way around. The vertices are represented by integer numbers.
The order in which the vertices are assigned do not matter, neither does
the direction, so as long as there exists a vertex 0 and there are no
skipped numbers, the input file is valid. We cannot have a vertex pointing
to nowhere.

For an example on an input file, see: KF_Project2_DFS/main/assign.txt

# Documentation

Go sets up a server with the documentation. In order to view the
documentation, we need to ensure that the program is in the
correct directory. The program should be in $GOPATH/src/.
Then run the following command in terminal:
`godoc -http=:6060`

Then we can put the process in the background by typing 
Control + Z and entering the command: "bg" 
without the quotations into the terminal.

To view the documentation, you have two options:

1) Using the terminal. 
To view documentation in the terminal, type:
`godoc KF_Project2_DFS/<file>`
where <file> is the file we want to view the documentation for.
For example: `godoc KF_Project2_DFS/graph`

2) Using a web browser. 
Open a web browser of your choosing and type into the address bar:
`http://localhost:6060/pkg/KF_Project2_DFS/`
From there you can select which package to view the documentation for.
