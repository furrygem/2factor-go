.PHONY: all
write_hello:
	echo "hello" > hello
cat_hello:
	@cat hello
clean:
	rm hello
DEFAULT_GOAL: all