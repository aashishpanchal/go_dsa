# 🚀 Go DSA

A modern, clean, and well-tested implementation of **Data Structures and Algorithms** in Golang. Built with performance and readability in mind. ⚡

## 🛠️ Tech Stack

- **Language:** [Go (Golang)](https://go.dev/) 🐹
- **Testing:** [testify](https://github.com/stretchr/testify) & [gotestsum](https://github.com/gotest-tools/gotestsum) 🧪
- **Task Runner:** [Taskfile](https://taskfile.dev/) 📋

## 🚀 Quick Start

### 1. Prerequisites

Make sure you have [Go](https://go.dev/doc/install) and [Task](https://taskfile.dev/installation/) installed.

### 2. Setup

```bash
# Clone the repo
git clone https://github.com/aashish/go_dsa.git

# Tidy dependencies
task mod:tidy
```

### 3. Running Tests

We use `gotestsum` for beautiful test output!

```bash
# Run all tests
task code:test

# Run only sorting tests
task test:sort

# Run only search tests
task test:search
```

## 🤝 Contributing

Contributions are welcome! Feel free to open a PR for any new algorithm or improvement. 🚀

## 📜 License

This project is licensed under the [MIT License](LICENSE).

---

**Author:** [Aashish Panchal](mailto:aipanchal51@gmail.com) 👨‍💻
