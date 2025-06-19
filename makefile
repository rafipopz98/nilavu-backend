# Default ENV is local
ENV ?= local

# Run the Go app with nodemon and correct env file
run:
	@echo "Loading .env.$(ENV)"
	ENV_FILE=".env.$(ENV)" nodemon \
		--watch './**/*.go' \
		--ext go \
		--signal SIGTERM \
		--exec 'go run main.go'
