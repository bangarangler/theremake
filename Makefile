build:
	@echo "heading to ./remake-client..."; \
	cd ./remake-client && \
	@echo "Building client..."; \
	npm install; \
	npm run-script build;

email:
	@echo "headint to ./remake-backend"; \
	cd theremake/remake-backend && \
	@echo "Emailing because google is silly"; \
	sh ./curl/automatedEmail.sh;

