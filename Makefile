install:
	@echo "Installing vagrantbox cli..."
	@cp ./bin/vagrantbox /usr/local/bin/vagrantbox
	@sleep 2
	@echo "Instalation is done."

.PHONY: install