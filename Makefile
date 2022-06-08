.PHONY: copy_graphs
copy_graphs:
	cp graphql/products.graphql packages/products/graph/products.graphql
	cp graphql/users.graphql packages/users/graph/users.graphql
	cp graphql/reviews.graphql packages/reviews/graph/reviews.graphql

.PHONY: generate
go_generate: 
	go generate ./...

.PHONY: run_all
run_all:
	(trap 'kill 0' SIGINT; go run packages/products/main.go & go run packages/reviews/main.go & go run packages/users/main.go)
