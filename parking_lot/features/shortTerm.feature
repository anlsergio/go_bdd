Feature: Calculate price for Short-Term Parking
  As an app developer at Southeast Airports Inc.
  I want to be able to send a specific duration time in minutes to an API endpoint and receive the corresponding parking cost for Short-Term Parking
  So that I have a simple way to calculate the corresponding parking cost for an elapsed time spent at the parking lot, which can be used by the frontend apps

Scenario: Calculate Short-Term Parking for half an hour
  When a "GET" request is sent to the endpoint "/short-term/30"
  Then the HTTP response code should be 200
  And the response content should be "$2.00"
