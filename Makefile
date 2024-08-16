dev: 
	npx nodemon --signal SIGTERM -e "templ go css js" -x "templ generate && npm run build --prefix ui && go run base/main.go serve" -i "**/*_templ.go" -i "static/css" -i "static/js"

generate: 
	templ generate

ui:
	npm run build --prefix ui

build: generate
	go build

run: generate
	go run main.go serve