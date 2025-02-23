# GitHub Actions to Compile Go Programs

Automatically compile and test a simple Go program using GitHub Actions whenever changes are pushed or a pull request is created to the `main` branch. The action installs Go, compiles the Go code, and runs the resulting program, ensuring it functions as expected without manual intervention.

## Usage

### 1. Clone the Repository

To get started, clone this repository to your local machine:

```bash
git clone https://github.com/BaseMax/github-actions-compile-golang.git
```

### 2. GitHub Action Setup

The GitHub Action is configured to automatically compile the Go program every time a change is pushed to the main branch or a pull request is created.

Create a new file in the `.github/workflows` directory:

- File name: `c-compile.yml`
- Action Setup: The action is defined in the `c-compile.yml` file, which runs the following steps:
  - Checkout code using the actions/checkout@v3 action.
  - Install GCC using the apt-get package manager on Ubuntu.
  - Compile the Go program using go build -o my_program main.go.
  - Run the compiled program and ensure that it works.

### 3. Example Go Program

The program compiled in this repository is a simple Go program that calculates the factorial of a number.

**Example Go Program (main.go):**

```go
package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println("Hello, GitHub Actions!")

    num := 5
    fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}
```

### 4. Running the GitHub Action

Whenever you push to the main branch or create a pull request, the GitHub Action will automatically:

- Check out the code.
- Set up the GCC compiler.
- Compile the program.
- Run the program to verify that the compilation was successful.
- You can view the results of the action under the Actions tab of the GitHub repository.

### GitHub Action Workflow

The workflow for compiling the C program is defined in the .github/workflows/c-compile.yml file.

```yaml
name: Compile and Test Go Program

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v3
        with:
          go-version: '1.24.0'

      - name: Compile Go program
        run: go build -o my_program main.go

      - name: Run tests
        run: |
          if ./my_program; then
            echo "Test Passed"
          else
            echo "Test Failed"
            exit 1
          fi
```

### Workflow Explanation:

- **Checkout code:** The actions/checkout action checks out the code from the repository.
- **Set up Go compiler:** Installs Golang compiler to compile the Go program.
- **Compile Go program:** The `go build` command compiles the main.go file into an executable named my_program.
- **Run tests:** The program is run, and if the execution is successful, it outputs "Test Passed"; otherwise, it outputs "Test Failed" and exits with a failure code.

## Contributing

If you'd like to contribute to this project, feel free to fork the repository, create a pull request with your changes, and submit it for review.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

Max Base, 2025

