.PHONY: all clean dep midi2mark mark2midi

all: midi2mark mark2midi

clean:
	rm -f midi2mark mark2midi

dep:
	go get -u -d -v github.com/beevik/etree

midi2mark: dep
	go build ./cmd/midi2mark

mark2midi: dep
	go build ./cmd/mark2midi
