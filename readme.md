# Prefix Matcher

A Console application that matches each of the prefixes in the given list against a set of words in a file and returns the longest matching prefix word.

---

[Go](https://go.dev/)

---

**Contents**

1. [Setup](#setup)
1. [Docs](#api-docs)
1. [Quality](#quality)
1. [Future Scope](#future-scope)

---

### Setup ###

1. Install Golang and ensure Go project can be run in the system.
1. Clone the repo in local
1. User can input the list of prefix in `main.go:22` file.
1. Change the working directory to the project directory by `cd prefix_matcher` command.
1. Compile the project by  `go build main.go` command.
1. If compilation is success, run the file by `go run main.go` command.
1. Change the input list and verify the outputs.

---


### Implementation ###
1. During precomputation, the words in the main file are arranged and sorted to multiple sub files.
1. During run, the words are checked from the sub files. Approximated the line number by binary search and further linear searched to get the exact longest word.
1. Go routines are spanned for each prefix to be found and the spanning count is kept within limit.

---

### Quality ###

Unit Test cases are available in `prefix_matcher/usecase` directory.

Mocks are available in `prefix_matcher/mocks` directory.

---

### Future Scope ###

1. Few hardcoded values can be set to environment variables like the maximum_goroutine_count etc.
1. Could implement logger for logging and tracing.
1. The application could be converted to a web server, so the sub files will not be deleted until the server stops.
1. Test coverage can be improved further.

---