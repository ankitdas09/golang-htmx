run:
	@npx tailwindcss --config tailwind.config.js -i input.css -o public/css/output.css
	@templ generate
	@go run ./cmd/app