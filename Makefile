build:
	@echo "heading to ./remake-client..."; \
	cd ./remake-client && \
	@echo "Building client..."; \
	npm install; \
	npm run-script build;

