# Golang Assignment #week17

The challenge

- Create a RESTFUL API with two endpoints.
- One of them that fetches the data in the provided MongoDB collection and returns the results in the requested format.
- Second endpoint is to create (POST) and fetch(GET) data from an in-memory database.

REQUIREMENTS

- The code should be written in Golang without using a web framework (mux, routers, chi, fiber,...)
- MongoDB data fetch endpoint should just handle HTTP POST requests
- The up to date repo should be publicly available in GitHub.

DELIVERABLES

- public repo with clear instructions on configuration and running the application locally.

# To run the project

- have docker installed
- run `docker-compose up -d`
- open another shell and type `air` and enter
- Done. The project should start and work just fine
- To stop the project run `docker-compose down` and then `ctrl+c` in the Air terminal

# For Endpoint Testing

- Use this to test the mongo handler:
  {
  "startDate": "2017-01-26",
  "endDate": "2017-01-30",
  "minCount": 100,
  "maxCount": 5000
  }

- For the In Memory Post:
  {
  "key": "come-on-barbie",
  "value": "lets-go-barbie"
  }
