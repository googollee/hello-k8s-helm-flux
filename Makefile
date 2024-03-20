.PHONY = image

image:
	docker build -t hello .

chart:
	helm template ./charts/hello/
