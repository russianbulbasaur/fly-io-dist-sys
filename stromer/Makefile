clean:
	rm -f stromer

test: test1 test2

test1: ./stromer
	../../maelstrom/maelstrom test -w echo --bin stromer --node-count 1 --time-limit 10

test2: ./stromer
	../../maelstrom/maelstrom  test -w unique-ids --bin stromer \
    --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition


test3a: ./stromer
	../../maelstrom/maelstrom test -w broadcast --bin stromer --node-count 1 --time-limit 20 --rate 10


test3b: ./stromer
	../../maelstrom/maelstrom test -w broadcast --bin stromer --node-count 5 --time-limit 20 --rate 10

test3c: ./stromer
	../../maelstrom/maelstrom test -w broadcast --bin stromer \
	--node-count 5 --time-limit 20 --rate 10 --nemesis partition

test3d: ./stromer
	../../maelstrom/maelstrom test -w broadcast --bin stromer \
	--node-count 25 --time-limit 20 --rate 100 --latency 100

test3e: ./stromer
	../../maelstrom/maelstrom test -w broadcast --bin stromer \
	--node-count 25 --time-limit 20 --rate 100 --latency 100


test4: ./stromer
	../../maelstrom/maelstrom test -w g-counter --bin stromer \
 	--node-count 3 --rate 100 --time-limit 20 --nemesis partition


serve:
	../../maelstrom/maelstrom serve

./stromer:
	go build .