build : bin/kuncen

.PHONY: bin/kuncen
bin/kuncen :
	go build -o bin/kuncen github.com/fajran/kuncen

