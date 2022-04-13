.PHONY: copy_graphs
copy_graphs:
	cp graphql/users.graphql packages/users/graph/users.graphql

.PHONY: generate
go_generate: 
	go generate ./...
